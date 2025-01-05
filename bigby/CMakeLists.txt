cmake_minimum_required(VERSION 3.31.3)
project(bigby)

# Library
find_package(PostgreSQL REQUIRED)
find_package(Lua REQUIRED)

set(CMAKE_CXX_STANDARD 26)
set(CMAKE_CXX_STANDARD_REQUIRED True)

add_subdirectory(src)

add_executable(bigby main.cpp)

include_directories(${LUA_INCLUDE_DIR})

# Include the directorys
# target_include_directories(bigby PRIVATE ${PostgreSQL_INCLUDE_DIRS})

# Linkar a biblioteca Lua com seu aplicativo
# target_link_libraries(bigby PRIVATE lua)

# Link library
# target_link_libraries(bigby PRIVATE ${PostgreSQL_LIBRARIES})
