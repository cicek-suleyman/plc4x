add_library(freertos-plc4x-abs-library)

target_include_directories(freertos-plc4x-abs-library PUBLIC
        Inc
)

target_sources(freertos-plc4x-abs-library PRIVATE
        Src/inet.c
)

target_link_libraries(freertos-plc4x-abs-library PUBLIC freertos_plus_tcp)