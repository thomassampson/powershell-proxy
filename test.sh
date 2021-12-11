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
████████ ███████ ███████ ████████ ███████ 
   ██    ██      ██         ██    ██      
   ██    █████   ███████    ██    ███████ 
   ██    ██           ██    ██         ██ 
   ██    ███████ ███████    ██    ███████ 
"
if [ -z "$1" ]; then
MAJOR=0
MINOR=0
BUILD=$(date +%s)
VERSION="$MAJOR.$MINOR.$BUILD"
else
VERSION="$1"
fi

START=$(date +%s)

echo "[TESTS] 🔵 Running Tests"
echo ""
if go test -v .; then
  echo ""
  echo "[TESTS] 🟢 Tests All Passed"
else
echo ""
echo "[TEST FAILED] 🔴 Tests Failed"
echo "[FAILED] ❌ Test Powershell Proxy | Version: '${VERSION}' | Build Time: '$BUILD_SECONDS sec'"
exit 1
fi
END=$(date +%s)
BUILD_SECONDS=$(echo "$END - $START" | bc)
echo "[SUCCESS] ✅ Test Powershell Proxy | Version: '${VERSION}' | Build Time: '$BUILD_SECONDS sec'"
