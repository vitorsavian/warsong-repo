// #include <iostream>

// int main(int argc, char *argv[]) {
//   std::cout << "Hello Bruv" << std::endl;
//   return 0;
// }

#include <iostream>
#include <lua.hpp>  // Incluir o cabeçalho Lua
#include "./src/cli/cli.hpp"    // Incluir o cabeçalho do CLI

int main() {
    // Inicializar Lua
    lua_State* L = luaL_newstate();
    luaL_openlibs(L);  // Carregar bibliotecas padrão do Lua

    // Executar um script Lua
    runLuaScript(L, "cli/script1.lua");

    // Fechar o estado Lua
    lua_close(L);
    return 0;
}
