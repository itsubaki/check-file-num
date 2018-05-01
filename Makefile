.PHONY: build
build:
	rm -rf build
	mkdir build

	GOOS=linux GOARGH=amd64 go build
	zip build/check-file-num_linux_amd64.zip check-file-num

	GOOS=linux GOARGH=386 go build
	zip build/check-file-num_linux_386.zip check-file-num

	GOOS=darwin GOARGH=amd64 go build
	zip build/check-file-num_darwin_amd64.zip check-file-num
