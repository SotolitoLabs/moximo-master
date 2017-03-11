# Targets (see each target for more information):
#   all: Build code.

# Default is arm
#

ifeq ($(GOARCH), )
	GOARCH=arm
endif

BINARY=_output/build/${GOARCH}/moximo-master

all:
	go build -o ${BINARY} *.go

clean:
	rm -rf _output

install:
	cp $BINARY /usr/bin/moximo-master
