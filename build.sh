#!/bin/bash
MAJOR=0
MINOR=0
BUILD=$(date +%s)
VERSION="$MAJOR.$MINOR.$BUILD"
BUILD_DIR="./build"
WINDOWS="${BUILD_DIR}/win/powershell-proxy_${VERSION}"
LINUX="${BUILD_DIR}/powershell-proxy_${VERSION}"
echo "[BUILD START] ✋ Building Powershell Proxy - Version: $VERSION"
echo "[BUILD] --> Cleaning Build Directory ${BUILD_DIR}"
rm -rf ${BUILD_DIR}
echo "[BUILD] --> Compiling Windows Binary"
env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o $WINDOWS .
echo "[BUILD] --> Windows Binary Compiled to $WINDOWS"
echo "[BUILD] --> Compiling Linux Binary"
env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o $LINUX .
echo "[BUILD] --> Linux Binary Compiled to $LINUX"
echo "[BUILD SUCCESS] Powershell Proxy"


