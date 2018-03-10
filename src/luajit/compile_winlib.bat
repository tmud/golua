@echo off
SET LUAJIT="LuaJIT-2.1.0-beta3"
SET MinGW64=C:\MinGW64

cd %LUAJIT%
cd src
del *.o
del *.a
cd ..
SET PATH=%PATH%;%MinGW64%\mingw32\bin
call mingw32-make BUILDMODE=static
cd ..
echo Done
pause