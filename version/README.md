# SunnyVersion 工具

## 使用方法

###只要在 import里面添加这一句就可以
```go
_ "github.com/sunnyregion/util/version"
```

### Makefile

```go
BUILD_NAME      := appRun          # _$(shell date "+%Y%m%d%H" )
BUILD_VERSION   := v0.1.0
BUILD_TIME      := $(shell date "+%F %T")
SOURCE          := main.go
TARGET_DIR      := ./
GO_VERSION      := $(shell git rev-parse HEAD )
EMOJI           := _|￣|○ -----🎉🎉🎉👍💁👌⚽🎍😍🎉🎉🎉------○|￣|_

all:
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	go build -ldflags  \
	"                  \
	-X 'github.com/sunnyregion/util/version.VERSION=${BUILD_VERSION}' \
	-X 'github.com/sunnyregion/util/version.BUILD_TIME=${BUILD_TIME}' \
	-X 'github.com/sunnyregion/util/version.BUILD_NAME=${BUILD_NAME}'  \
	-X 'github.com/sunnyregion/util/version.GO_VERSION=${GO_VERSION}' \
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
