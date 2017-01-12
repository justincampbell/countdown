NAME := countdown

GOX_OSARCH := darwin/amd64 linux/amd64
VERSION := 1.0.0

LDFLAGS := "-X main.Version=$(VERSION)" 

default: test build

test:
	go test ./...

build:
	go build -ldflags $(LDFLAGS) -o release/countdown

install: build
	cp release/countdown /usr/local/bin/countdown

release:
	go get github.com/mitchellh/gox
	gox \
	  -ldflags $(LDFLAGS) \
	  -osarch="$(GOX_OSARCH)" \
	  -output="release/$(NAME)_{{.OS}}_{{.Arch}}_$(VERSION)" \
	  ./...
	cd release/; for f in *; do mv -v $$f $(NAME); tar -zcf $$f.tar.gz $(NAME); rm $(NAME); done

clean:
	rm -rf release/

.PHONY: build clean default release test
