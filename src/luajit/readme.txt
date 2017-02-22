1) unpack luajit archive (LuaJIT-2.1.0-beta2 folder).
2) build luajit static lib
- on the Windows (use mingw64, cygwin not compatible with go): mingw32-make BUILDMODE=static in LuaJIT-2.1.0-beta2 folder.
- on the mac/linux: make BUILDMODE=static in LuaJIT-2.1.0-beta2 folder.
3) build go application (run go build) from project root dir. Don't forget set GOPATH variable.
- For Windows use go 1.8+! (1.7 has problems with linking confilct ntdll vs msvcrt).
- On other platforms works 1.7.
