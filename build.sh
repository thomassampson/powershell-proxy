#!/bin/bash
echo "


██████╗ ██╗    ██╗███████╗██╗  ██╗    ██████╗ ██████╗  ██████╗ ██╗  ██╗██╗   ██╗     █████╗ ██████╗ ██╗
██╔══██╗██║    ██║██╔════╝██║  ██║    ██╔══██╗██╔══██╗██╔═══██╗╚██╗██╔╝╚██╗ ██╔╝    ██╔══██╗██╔══██╗██║
██████╔╝██║ █╗ ██║███████╗███████║    ██████╔╝██████╔╝██║   ██║ ╚███╔╝  ╚████╔╝     ███████║██████╔╝██║
██╔═══╝ ██║███╗██║╚════██║██╔══██║    ██╔═══╝ ██╔══██╗██║   ██║ ██╔██╗   ╚██╔╝      ██╔══██║██╔═══╝ ██║
██║     ╚███╔███╔╝███████║██║  ██║    ██║     ██║  ██║╚██████╔╝██╔╝ ██╗   ██║       ██║  ██║██║     ██║
╚═╝      ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝    ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝       ╚═╝  ╚═╝╚═╝     ╚═╝                                                                                                                                                                                                                                                                                              
"
START=$(date +%s)
MAJOR=0
MINOR=0
BUILD=$(date +%s)
VERSION="$MAJOR.$MINOR.$BUILD"
BUILD_DIR="./build"
TMP_DIR="./tmp"
WINDOWS="${BUILD_DIR}/win/powershell-proxy_${VERSION}"
LINUX="${BUILD_DIR}/linux/powershell-proxy_${VERSION}"
echo "[BUILD START] 🔥 Building Powershell Proxy - Version: $VERSION"
echo "[BUILD] 🔵 Cleaning Build & Temp Directories - ${BUILD_DIR} & ${TMP_DIR}"
rm -rf ${BUILD_DIR}
rm -rf ${TMP_DIR}
echo "[BUILD] 🟢 Build Directory Cleaned"
echo "[BUILD] 🔵 Compiling Windows Binary"
env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$VERSION" -o "$WINDOWS.exe" .
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
echo "[BUILD SUCCESS] ✅ Built Powershell Proxy | Version: '${VERSION}' | Build Time: '$BUILD_SECONDS sec'"


