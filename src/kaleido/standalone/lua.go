package main

import (
	"fmt"
	"kaleido/lua"
	"kaleido/lmathlib"
	"kaleido/linit"
	"kaleido/lstate"
	"kaleido/lualib"
	"kaleido/utils"
	. "kaleido/proxy"
)

func main() {
	L := &LuaState.Lua_State{}
	fmt.Println("kaleido start...")
	fmt.Println(Lua.LUA_VERSION)
	fmt.Println(LuaMathLib.PI)
	fmt.Println(LuaLib.LUA_FILEHANDLE)
	Tools.Printf("%s\n", "hello")
	d := NewDateTimeProxyInit()
	d.SetNow()
	LuaInit.LuaL_openlibs(L)
}
