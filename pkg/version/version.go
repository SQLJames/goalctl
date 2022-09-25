package version

import (
	"fmt"
	"runtime"

	"github.com/sqljames/goalctl/pkg/info"
)

// These variables are usually defined via LDFLags at compile-time
var (
	// Version is the default version of the package/tool
	Version = Info{
		applicationName: applicationName,
		release:         release,
		commitHash:      commitHash,
		buildDate:       buildDate,
		buildTag:        buildTag,
		buildOS:         runtime.GOOS,
		buildARCH:       runtime.GOARCH,
		buildTarget:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		buildVersion:    buildVersion,
		buildBranch:     buildBranch,
		goVersion:       runtime.Version(),
	}
	// Name of application
	applicationName = info.GetApplicationName()
	commitHash      = "unknown"
	buildDate       = "unknown"
	buildVersion    = "0.0.0"
	buildTag        = "unknown"
	release         = "dev"
	buildBranch     = "unknown"
)

// Info contains all the versioning information for a package/tool
type Info struct {
	applicationName string
	release         string
	commitHash      string
	buildDate       string
	buildOS         string
	buildARCH       string
	buildTarget     string
	buildTag        string
	buildVersion    string
	buildBranch     string
	goVersion       string
}

func (vi *Info) String() string {
	return fmt.Sprintf("\tApplication:\t%s\n\tVersion:\t%s-v%s \n\tGitBranch:\t%s \n\tGitBuildTag:\t%s \n\tGitCommitHash:\t%s \n\tBuildTarget:\t%s \n\tBuildDate:\t%s", vi.applicationName, vi.release, vi.buildVersion, vi.buildBranch, vi.buildTag, vi.commitHash, vi.buildTarget, vi.buildDate)
}

func (vi *Info) CliPrinter() string {
	return fmt.Sprintf("Application:\t%s\nVersion:\t%s-v%s \nGitBranch:\t%s\nGitBuildTag:\t%s \nGitCommitHash:\t%s \nBuildTarget:\t%s \nBuildDate:\t%s", vi.applicationName, vi.release, vi.buildVersion, vi.buildBranch, vi.buildTag, vi.commitHash, vi.buildTarget, vi.buildDate)
}

func GetGOARCH() string {
	return Version.buildARCH
}
func GetGOOS() string {
	return Version.buildOS
}

func GetVersion() string {
	return Version.buildVersion
}
