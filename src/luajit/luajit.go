
package luajit

/*
#cgo CFLAGS: -ILuaJit-2.1.0-beta2/src
#cgo LDFLAGS: -L${SRCDIR}/LuaJit-2.1.0-beta2/src -lluajit
#include <stdlib.h>
#include <stdio.h>
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// LuaStatePtr is a type to respresent `struct lua_State`
type LuaStatePtr *C.struct_lua_State

// State stores lua state
type State struct {
	state LuaStatePtr
}

// NewState creates new Lua state
func NewState() *State {
	l := C.luaL_newstate()
	return &State{l}
}

// OpenLibs loads lua libraries
func (L *State) OpenLibs() {
	C.luaL_openlibs(L.state)
}

// LoadString loads code into lua stack
func (L *State) LoadString(str string) error {
	csrc := C.CString(str)
	defer C.free(unsafe.Pointer(csrc))

	if C.luaL_loadstring(L.state, csrc) != 0 {
		return fmt.Errorf(L.errorString())
	}

	return nil
}

// Run executes code in stack
func (L *State) Run() error {
	if C.lua_pcall(L.state, 0, 0, 0) != 0 {
		return fmt.Errorf(L.errorString())
	}

	return nil
}

// Close destroys lua state
func (L *State) Close() {
	C.lua_close(L.state)
}

func (L *State) errorString() string {
	return C.GoString(C.lua_tolstring(L.state, -1, nil))
}
