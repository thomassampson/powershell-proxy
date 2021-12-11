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
echo "[START] 🔥 Building Powershell Proxy - Version: $VERSION"
echo "[CLEANUP] 🔵 Cleaning Build & Temp Directories - ${BUILD_DIR} & ${TMP_DIR}"
rm -rf ${BUILD_DIR}
rm -rf ${TMP_DIR}
echo "[CLEANUP] 🟢 Build Directory Cleaned"
echo "[TESTS] 🔵 Running Tests"
echo ""
if go test -v .; then
  echo ""
  echo "[TESTS] 🟢 Tests All Passed"
else
echo ""
echo "[TEST FAILED] 🔴 Tests Failed"
echo "[FAILED] ❌ Built Powershell Proxy | Version: '${VERSION}' | Build Time: '$BUILD_SECONDS sec'"
exit 1
fi
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
echo "[SUCCESS] ✅ Built Powershell Proxy | Version: '${VERSION}' | Build Time: '$BUILD_SECONDS sec'"


