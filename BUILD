cc_library(
    name = "libtensorflow",
    srcs = glob(["tensorflow/**/*.h"]) + [
        "@libtensorflow-cpu-linux-x86_64-2.9.0//:lib/libtensorflow.so",
        "@libtensorflow-cpu-linux-x86_64-2.9.0//:lib/libtensorflow.so.2",
        "@libtensorflow-cpu-linux-x86_64-2.9.0//:lib/libtensorflow.so.2.9.0",
    ],
    visibility = ["//visibility:public"],
)
