#include <iostream>
#include <lua.hpp>

void cmd(lua_State *L, const std::string& command, const std::string &args)
{
    lua_getglobal(L, "Main");
    lua_pushstring(L, command.c_str());

    size_t start = 0;
    size_t end = args.find(' ');

    while (end != std::string::npos) {
        lua_pushstring(L, args.substr(start, end - start).c_str());
        start = end + 1;
        end = args.find(' ', start);
    }

    lua_pushstring(L, args.substr(start).c_str());

    if (lua_pcall(L, 2, 0, 0) != LUA_OK) {
        std::cerr << "Error calling Lua function: " << lua_tostring(L, -1) << std::endl;
    }
}
