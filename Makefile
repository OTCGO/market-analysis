.PHONY: build deps run clean

MAIN = "main.go"
OUTPUT = "./bin/main"

build:
	@go build -o ${OUTPUT} ${MAIN}  

deps:
	@dep ensure 

pull:
	git reset --hard
	git pull origin master

run: pull build
	${OUTPUT}

clean:
	rm -rf ${OUTPUT}