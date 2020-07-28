# SunnyVersion 工具

## 使用方法

###只要在 import里面添加这一句就可以
```go
_ "github.com/sunnyregion/util/version"
```

### Makefile

```go
BUILD_NAME      := appRun          # _$(shell date "+%Y%m%d%H" )
BUILD_VERSION   := V1.0.1.9-$(shell date "+%Y%m%d")
BUILD_TIME      := $(shell date "+%F %T")
SOURCE          := main.go
TARGET_DIR      := ./
GIT_ID           := $(shell git rev-parse HEAD )
GO_VERSION      := $(shell go version)
GIT_BRANCH	:= $(shell git symbolic-ref --short -q HEAD)
EMOJI           := _|￣|○ -----🎉🎉🎉👍💁👌huizhi api⚽🎍😍🎉🎉🎉------○|￣|_


all:
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags  \
		" \
		-X 'github.com/sunnyregion/util/version.VERSION=${BUILD_VERSION}' \
		-X 'github.com/sunnyregion/util/version.BuildTime=${BUILD_TIME}' \
		-X 'github.com/sunnyregion/util/version.BuildName=${BUILD_NAME}'  \
		-X 'github.com/sunnyregion/util/version.GitID=${GIT_ID}' \
		-X 'github.com/sunnyregion/util/version.GitBranch=${GIT_BRANCH}' \
		-X 'github.com/sunnyregion/util/version.GoVersion=${GO_VERSION}' \
		-X 'github.com/sunnyregion/util/version.EMOJI=${EMOJI}' \
		" \

    -o ${BUILD_NAME} ${SOURCE}

clean:
	rm ${BUILD_NAME} -f

install:
	mkdir -p ${TARGET_DIR}
	cp ${BUILD_NAME} ${TARGET_DIR} -f

.PHONY : all clean install ${BUILD_NAME}


```

## 编译

make

## install

make install
