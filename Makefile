# Targets (see each target for more information):
#   all: Build code.

BINARY=moximo-master

all:
	go build -o ${BINARY} *.go

