#!/bin/bash
echo "


██████╗ ██╗    ██╗███████╗██╗  ██╗    ██████╗ ██████╗  ██████╗ ██╗  ██╗██╗   ██╗     █████╗ ██████╗ ██╗
██╔══██╗██║    ██║██╔════╝██║  ██║    ██╔══██╗██╔══██╗██╔═══██╗╚██╗██╔╝╚██╗ ██╔╝    ██╔══██╗██╔══██╗██║
██████╔╝██║ █╗ ██║███████╗███████║    ██████╔╝██████╔╝██║   ██║ ╚███╔╝  ╚████╔╝     ███████║██████╔╝██║
██╔═══╝ ██║███╗██║╚════██║██╔══██║    ██╔═══╝ ██╔══██╗██║   ██║ ██╔██╗   ╚██╔╝      ██╔══██║██╔═══╝ ██║
██║     ╚███╔███╔╝███████║██║  ██║    ██║     ██║  ██║╚██████╔╝██╔╝ ██╗   ██║       ██║  ██║██║     ██║
╚═╝      ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝    ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝       ╚═╝  ╚═╝╚═╝     ╚═╝                                                                                                                                                                                                                                                                                              
"
echo "
██████  ██    ██ ██ ██      ██████  
██   ██ ██    ██ ██ ██      ██   ██ 
██████  ██    ██ ██ ██      ██   ██ 
██   ██ ██    ██ ██ ██      ██   ██ 
██████   ██████  ██ ███████ ██████  
"

if [ -z "$1" ]; then
VERSION="$(date +%Y).$(date +%m).$(date +%d).$(date +%s)"
else
VERSION="$1"
fi

START=$(date +%s)

BUILD_DIR="./build"
TMP_DIR="./tmp"
WINDOWS="${BUILD_DIR}/win/powershell-proxy_win_amd64.exe"
LINUX="${BUILD_DIR}/linux/powershell-proxy_linux_amd64"
echo "[START] 🔥 Building Powershell Proxy - Version: $VERSION"
echo "[CLEANUP] 🔵 Cleaning Build & Temp Directories - ${BUILD_DIR} & ${TMP_DIR}"
rm -rf ${BUILD_DIR}
rm -rf ${TMP_DIR}
echo "[CLEANUP] 🟢 Build Directory Cleaned"
echo "[BUILD] 🔵 Compiling Windows Binary"
env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o $WINDOWS .
echo "[BUILD] 🟢 Windows Binary Compiled to $WINDOWS"
echo "[BUILD] 🔵 Compiling Linux Binary"
env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o $LINUX .
echo "[BUILD] 🟢 Linux Binary Compiled to $LINUX"
echo "[BUILD] ⬇️  Binaries Successfully Created"
echo ""
ls -R ${BUILD_DIR}
echo ""
END=$(date +%s)
BUILD_SECONDS=$(echo "$END - $START" | bc)
echo "[SUCCESS] ✅ Built Powershell Proxy | Version: '${VERSION}' | Build Time: '$BUILD_SECONDS sec'"
