.PHONY: build deps run clean scp

MAIN = "main.go"
OUTPUT = "./bin/main"
SERVER = "root@future.otcgo.cn:/root/gowork/src/market-analysis/bin"

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

scp: build
	scp ${OUTPUT}  ${SERVER}	