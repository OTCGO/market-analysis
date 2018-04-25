.PHONY: build deps run clean

build:
	@go build main.go 

deps:
	@dep ensure 

run:
	./main
clean:
	rm ./main