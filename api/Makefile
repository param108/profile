.PHONY: build clean version
version:
	git branch --show-current > server/instance/version.txt
	git for-each-ref --points-at `git branch --show-current` | head -1 | cut -d' ' -f 1 >> server/instance/version.txt

build:	version
	- rm -rf build
	mkdir -p build
	go build -o build ./...

test:	version
	go test -v ./...

clean:
	rm server/instance/version.txt
	rm -rf build
