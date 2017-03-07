# Targets (see each target for more information):
#   all: Build code.

# Default is x86_64
#

ifeq ($(GOARCH), )
	GOARCH=amd64
endif

BINARY=_output/build/${GOARCH}/moximo-master

all:
	go build -o ${BINARY} *.go

clean:
	rm -rf _output
