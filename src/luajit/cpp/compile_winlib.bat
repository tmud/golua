@echo off
SET LUAJIT="LuaJIT-2.1.0-beta3"
del /S /F /Q %LUAJIT%
RMDIR %LUAJIT /S /Q
del /S /F /Q luajitlib
RMDIR luajit /S /Q
7za.exe x %LUAJIT%.zip
move %LUAJIT% luajit
copy Makefile luajit\src
cd luajit
call mingw32-make BUILDMODE=static
cd ..
echo Done
pause
