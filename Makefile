BIN=sp33ch

.PHONY: dep dep-update dep-download build

all: dep-download build

dep:
	go mod init
	go mod tidy

dep-update:
	go mod verify
	go mod tidy

dep-download:
	go mod download

build:
	go build -o ${BIN} ./cmd/server/rest