package sqlite

import (
	"bytes"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"strings"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

//go:embed sqlc/schema/*.sql
var schema embed.FS

// EmbedFS interface for supporting embed.FS as injected filesystem and provide possibility to mock.
type EmbedFS interface {
	fs.ReadFileFS
}

// EmbedFileMigrationSource implements MigrationSource and provide migrations from a native embed filesystem.
type EmbedFileMigrationSource struct {
	Filesystem EmbedFS
}

var (
	ErrEmbedWalkFailed    = errors.New(`failed to walk recursive over embed source directory`)
	ErrEmbedReadDirFailed = errors.New(`directory read failed`)
)

// FindMigrations is part of MigrationSource implementation.
func (f *EmbedFileMigrationSource) FindMigrations() ([]*migrate.Migration, error) {
	items := make([]*migrate.Migration, 0)
	err := fs.WalkDir(f.Filesystem, `.`, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())

			return fmt.Errorf(`%w: %v`, ErrEmbedReadDirFailed, err)
		}
		// from now we always return nil cause if we send err, WalkDir stop processing.
		if entry.IsDir() {
			return nil
		}

		if !strings.HasSuffix(entry.Name(), `.sql`) {
			return nil
		}

		content, err := f.Filesystem.ReadFile(path)
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())

			return err
		}

		migration, err := migrate.ParseMigration(entry.Name(), bytes.NewReader(content))
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())

			return err
		}

		items = append(items, migration)

		return nil
	})

	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return items, fmt.Errorf(`%w: %v`, ErrEmbedWalkFailed, err)
	}

	return items, nil
}

func runMigrations(sqlDB *sql.DB) (int, error) {
	MigrationSource := EmbedFileMigrationSource{Filesystem: schema}

	migration, err := MigrationSource.FindMigrations()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return 0, err
	}

	numberOfMigrations, err := migrate.Exec(sqlDB, "sqlite3", &migrate.MemoryMigrationSource{Migrations: migration}, migrate.Up)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return 0, err
	}

	jlogr.Logger.ILog.Trace("Applied migrations!", "numberOfMigrations", numberOfMigrations)

	return numberOfMigrations, nil
}
