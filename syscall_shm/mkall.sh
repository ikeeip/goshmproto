#!/bin/sh

GOOSARCH="${GOOS}_${GOARCH}"

mktypes="GOARCH=$GOARCH go tool cgo -godefs"

if [ -n "$mktypes" ]; then echo "$mktypes types_$GOOS.go > ztypes_$GOOSARCH.go"; fi

# GOARCH=amd64 GOOS=darwin go tool cgo -godefs types_darwin.go > ztypes_darwin_amd64.go
# GOARCH=amd64 GOOS=linux go tool cgo -godefs types_linux.go > ztypes_linux_amd64.go