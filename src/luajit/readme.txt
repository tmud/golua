1) unpack luajit archive (LuaJIT-2.1.0-beta3 folder).
2) build luajit static lib
- on the Windows (use mingw64!, cygwin not compatible with go): mingw32-w64-make BUILDMODE=static in LuaJIT-2.1.0-beta3 folder.
  (try compile_winlib.bat)
  http://tdm-gcc.tdragon.net/download
- on the mac/linux: make BUILDMODE=static in LuaJIT-2.1.0-beta3 folder.
3) build go application (run go build) from project root dir. Don't forget set GOPATH variable.
- For Windows use go 1.8+! (1.7 has problems with linking confilct ntdll vs msvcrt).
- On other platforms works 1.7.
