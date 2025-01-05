cmake_minimum_required(VERSION 3.31.3)
project(mordenkeinen)

# Library
find_package(PostgreSQL REQUIRED)

add_executable(mordenkeinen main.cpp)

set(CMAKE_CXX_STANDARD 26)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Include the directorys
target_include_directories(mordenkeinen PRIVATE ${PostgreSQL_INCLUDE_DIRS})

# Link library
target_link_libraries(mordenkeinen PRIVATE ${PostgreSQL_LIBRARIES})
