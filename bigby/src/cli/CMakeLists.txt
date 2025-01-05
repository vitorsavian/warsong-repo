# Adicionar os arquivos de código C++ para o CLI
add_library(cli STATIC cli.cpp)

# Linkar a biblioteca Lua com o CLI
target_link_libraries(cli PRIVATE ${LUA_LIBRARIES})

# Incluir o diretório onde estão os arquivos de cabeçalhos
target_include_directories(cli PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

# Copiar os arquivos .lua para o diretório de build
file(GLOB LUA_SCRIPTS "*.lua")
configure_file(${LUA_SCRIPTS} ${CMAKE_BINARY_DIR}/cli/ COPYONLY)
