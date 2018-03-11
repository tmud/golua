1) unpack luajit archive (LuaJIT-2.1.0-beta3 folder).
2) build luajit static lib
- on the Windows install mingw64 with x86_64 arch!, cygwin not compatible with go, run compile_winlib.bat
- on the mac/linux: run compile_linuxmac.sh
3) build go application (run go build) from project root dir. Don't forget set GOPATH variable.
- For Windows use go 1.8+! (1.7 has problems with linking confilct ntdll vs msvcrt).
- On other platforms works 1.7.
