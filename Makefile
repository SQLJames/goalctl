local:
	go run ./magefiles/mage.go build

release:
	go run ./magefiles/mage.go release

clean:
	go run ./magefiles/mage.go clean

benchmark:
	go test ./... -bench=. -benchtime=10s -benchmem

sqlc:
	sqlc generate --file ./pkg/storage/sqlite/sqlc/sqlc.yaml

.PHONY: build release benchmark clean