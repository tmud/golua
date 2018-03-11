#!/bin/bash

LUAJIT="LuaJIT-2.1.0-beta3"

rm -rf $LUAJIT
rm -rf luajit
echo $LUAJIT
unzip $LUAJIT
mv $LUAJIT luajit
cp Makefile luajit/src
cd luajit
make BUILDMODE=static
cd ..