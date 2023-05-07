PROJECT_NAME = $(shell basename $${PWD})
PLATFORM_MACOS = mac
PLATFORM_LINUX = linux
PLATFORM_WINDOWS = exe
PLATFORM_ARM = armlinux
BIN_PATH = ${PWD}/bin


flags="-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git describe --long --dirty --abbrev=14` -X 'main.goversion=$(go version)'"

all: compile

build:
	@go build

clean:
	@rm -rf ${PWD}/${PROJECT_NAME}
	@rm -rf ${BIN_PATH}/*

compile:
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "${flags}" -x -o ${BIN_PATH}/${PROJECT_NAME}_${PLATFORM_ARM}
	GOOS=linux GOARCH=amd64 go build -ldflags "${flags}" -x -o ${BIN_PATH}/${PROJECT_NAME}_${PLATFORM_LINUX}
	GOOS=windows GOARCH=amd64 go build -ldflags "${flags}" -x -o ${BIN_PATH}/${PROJECT_NAME}.${PLATFORM_WINDOWS}
	GOOS=darwin GOARCH=amd64 go build -ldflags "${flags}" -x -o ${BIN_PATH}/${PROJECT_NAME}_${PLATFORM_MACOS}

