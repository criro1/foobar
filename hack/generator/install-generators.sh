#!/usr/bin/env bash


ROOT_PROJECT=$PWD
TOOLS_MAIN_DIR=".tools"

VER="3.19.3"
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     OS=linux;;
    Darwin*)    OS=osx;;
    *)          OS="xxx"
esac

echo ${ROOT_PROJECT} ${OS} $VER
if [ ! -d ./${TOOLS_MAIN_DIR} ]; then
mkdir ${TOOLS_MAIN_DIR}
fi
if [ ! -d ./${TOOLS_MAIN_DIR}/bin ]; then
mkdir ${TOOLS_MAIN_DIR}/bin
fi
if [ ! -d ./${TOOLS_MAIN_DIR}/dist ]; then
mkdir ${TOOLS_MAIN_DIR}/dist
fi
if [ ! -d ./${TOOLS_MAIN_DIR}/protoc ]; then
mkdir ${TOOLS_MAIN_DIR}/protoc
fi


echo https://github.com/protocolbuffers/protobuf/releases/download/v$VER/protoc-$VER-$OS-x86_64.zip

if [ ! -f ./${TOOLS_MAIN_DIR}/dist/protoc.zip ]; then
curl https://github.com/protocolbuffers/protobuf/releases/download/v$VER/protoc-$VER-$OS-x86_64.zip -o ./${TOOLS_MAIN_DIR}/dist/protoc.zip -L
fi

unzip -q -o ./${TOOLS_MAIN_DIR}/dist/protoc.zip -d ./${TOOLS_MAIN_DIR}/protoc
mv ./${TOOLS_MAIN_DIR}/protoc/bin/protoc ./${TOOLS_MAIN_DIR}/protoc/bin/protoc-b2bchain
echo "installed "`./${TOOLS_MAIN_DIR}/protoc/bin/protoc-b2bchain --version`

for genpkg in `go list -f '{{ join .Imports "\n" }}'  tool_import.go`
do
    echo "building $genpkg..."
    go build -mod=vendor -o ./${TOOLS_MAIN_DIR}/bin/`basename $genpkg`-b2bchain -trimpath ${genpkg}
    echo "installed $genpkg"
done

go get github.com/lygo/graphql-go/cmd/gql-gen-resolver
mv ${GOPATH}/bin/gql-gen-resolver ${TOOLS_MAIN_DIR}/bin

go get github.com/golang/mock/mockgen@v1.4.3
mv ${GOPATH}/bin/mockgen ${TOOLS_MAIN_DIR}/bin
