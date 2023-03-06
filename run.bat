@echo off

set "MONGODB_URI=mongodb://localhost:27017"
set TF2_LOGPATH=C:\\Program Files (x86)\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\console.log
set MONGODB_NAME=TF2
set ENABLE_AUTOBALANCE_COMMENT=1

.\build\main-windows-amd64.exe
