add_subdirectory(cli)

# Adicionar o código principal (main.cpp) fora do diretório src
# add_executable(bigby ../main.cpp)

# Linkar a biblioteca Lua com o código principal
#target_link_libraries(bigby PRIVATE ${LUA_LIBRARIES} cli)

# Incluir diretórios de cabeçalhos para o código C++ dentro de src
#target_include_directories(bigby PRIVATE ${CMAKE_SOURCE_DIR}/src)
