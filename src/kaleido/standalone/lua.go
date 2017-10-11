package main

import (
	"fmt"
	"kaleido/lua"
	"kaleido/lmathlib"
	"kaleido/linit"
	"kaleido/lstate"
)

func main() {
	L := &LuaState.Lua_State{}
	fmt.Println("kaleido start...")
	fmt.Println(Lua.LUA_VERSION)
	fmt.Println(LuaMathLib.PI)
	LuaInit.LuaL_openlibs(L)
}
