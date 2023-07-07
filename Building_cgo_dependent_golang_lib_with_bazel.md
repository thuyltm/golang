### 1.  Golang external dependencies

#### <span style="color:blue"> Method 1.</span> Get checksum sha256 and version from go.mod and update macro 

```
go get github.com/zergon321/reisen
go install github.com/zergon321/reisen
```

Next, use gazelle to update deps.bzl
```shell
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

Lastly, update repositories.bzl
```shell
go_repository(
    name = "github_com_zergon321_reisen",
    importpath = "github.com/zergon321/reisen",    
    sum = "h1:rflBVtSoc4xXR3tXN+b2Hvk/Fb3qaJhgj+CJ/YWOYmI=",
    version = "v0.1.8",
)
```

#### <span style="color:blue"> Method 2.</span> Manual add dependency package in WORKSPACE.bazel and repositories.bzl 
WORKSPACE
```
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

gazelle_dependencies(go_repository_default_config = "//:WORKSPACE.bazel")

load("//:repositories.bzl", "go_repositories")

go_repository(
    name = "github_com_zergon321_reisen",
    importpath = "github.com/zergon321/reisen", # from go.sum data
    sum = "h1:rflBVtSoc4xXR3tXN+b2Hvk/Fb3qaJhgj+CJ/YWOYmI=", # from go.sum data
    version = "v0.1.8" # from go.sum data
)

# gazelle:repository_macro repositories.bzl%go_repositories
go_repositories()
```
Or repositories.bzl
```shell
go_repository(
    name = "github_com_zergon321_reisen",
    importpath = "github.com/zergon321/reisen",
    sum = "h1:rflBVtSoc4xXR3tXN+b2Hvk/Fb3qaJhgj+CJ/YWOYmI=",
    version = "v0.1.8",
)
```

#### <span style="color:blue"> Method 3.</span> If dependency package depend on C lib, there more new steps 
<div style="color:orange"> · Step 1. Install C dependencies to system </div> for example in Ubuntu apt-get install
<div style="color:orange"> · Step 2. Let Bazel knew it </div> 

<div style="color:greenyellow">· Method 1. Download source code and compiling CMake Projects From Within a Bazel Workspace </div>
Reference Link: https://rotemtam.com/2020/10/30/bazel-building-cgo-bindings/

WORKSPACE
```shell
http_archive(
    name = "rules_foreign_cc",
    # TODO: Get the latest sha256 value from a bazel debug message or the latest
    #       release on the releases page: https://github.com/bazelbuild/rules_foreign_cc/releases
    #
    # sha256 = "...",
    strip_prefix = "rules_foreign_cc-6ecc134b114f6e086537f5f0148d166467042226",
    url = "https://github.com/bazelbuild/rules_foreign_cc/archive/6ecc134b114f6e086537f5f0148d166467042226.tar.gz",
)

load("@rules_foreign_cc//foreign_cc:repositories.bzl", "rules_foreign_cc_dependencies")

rules_foreign_cc_dependencies()

http_archive(
    name = "librdkafka",
    build_file_content = """load("@rules_foreign_cc//foreign_cc:defs.bzl", "cmake")

filegroup(
    name = "sources",
    srcs = glob(["**"]),
)

cmake(
    name = "librdkafka",
    cache_entries = {
        "RDKAFKA_BUILD_STATIC": "ON",
        "WITH_ZSTD": "OFF",
        "WITH_SSL": "OFF",
        "WITH_SASL": "OFF",
        "ENABLE_LZ4_EXT": "OFF",
        "WITH_LIBDL": "OFF",
    },
    lib_source = ":sources",
    out_static_libs = [
        "librdkafka++.a",
        "librdkafka.a",
    ],
    visibility = ["//visibility:public"],
)
""",
    sha256 = "7be1fc37ab10ebdc037d5c5a9b35b48931edafffae054b488faaff99e60e0108",
    strip_prefix = "librdkafka-2.1.1",
    urls = ["https://github.com/confluentinc/librdkafka/archive/refs/tags/v2.1.1.tar.gz"],
)
```

<div style="color:greenyellow"> · Method 2. Complex BUILD file. allow the code to `#include opus.h` directly </div>
Reference Link: https://blog.modest-destiny.com/posts/building-golang-cgo-with-bazel/

WORKSPACE
```shell
new_local_repository(
    name = "libopus",
    build_file = "@//thirdparty/opus:BUILD.opus",
    path = "/usr",
)
```

BUILD.bazel
```shell
load("@rules_cc//cc:defs.bzl", "cc_library", "cc_import")

# The headers can be added in hdrs param in cc_import() or cc_library().
# However, if cc_library() wrapper is used, it is suggested to add in the cc_library().
# The header files are included in the package manager. For debian related systems, we can use
# $ dpkg -L libopus-dev
# $ dpkg -L libopusfile-dev
# to list the installed files.
cc_library(
    name = "libopus",
    hdrs = [
        "//:include/opus/opus.h",
        "//:include/opus/opus_defines.h",
        "//:include/opus/opus_multistream.h",
        "//:include/opus/opus_projection.h",
        "//:include/opus/opus_types.h",
        "//:include/opus/opusfile.h",
    ],
    strip_include_prefix = "include/opus", # allow the code to #include <opus.h> directly
    deps = [               # the sequence of deps is important because it impacts the sequence of linking
        "//:libopusfile_private",
        "//:libopusurl_private",
        "//:libopus_private",
    ],
    visibility = ["//visibility:public"],
)

# Each cc_import() shall contain only one library.
# If multiple libraries are needed, we must use the cc_library() as a wrapper
# to depend on the required multiple libraries.
cc_import(
    name = "libopus_private",
    static_library = "//:lib/x86_64-linux-gnu/libopus.a",
    shared_library = "//:lib/x86_64-linux-gnu/libopus.so",
    visibility = ["//visibility:private"],
)

cc_import(
    name = "libopusfile_private",
    static_library = "//:lib/libopusfile.a",
    shared_library = "//:lib/libopusfile.so",
    visibility = ["//visibility:private"],
)

cc_import(
    name = "libopusurl_private",
    static_library = "//:lib/libopusurl.a",
    shared_library = "//:lib/libopusurl.so",
    visibility = ["//visibility:private"],
)
```
<div style="color:greenyellow"> · Method 3. Simple BUILD file, care only share object dynamic link </div>
WORKSPACE

```
new_local_repository(
    name = "system_libs",
    build_file_content = """
load("@rules_cc//cc:defs.bzl", "cc_library")

cc_library(
    name = "libvirt",
    srcs = ["libvirt.so", "libvirt-admin.so", "libvirt-lxc.so", "libvirt-qemu.so"],
    visibility = ["//visibility:public"],
)
""",
    path = "/usr/lib64",
)


```
BUILD.bazel
```
go_library(
  ....
  cdeps = ["@system_libs//:libvirt"],
  cgo = True
)  
```

<span style="color:orange"> · Step 3.</span> Update BUILD.bazel using cgo, cdeps

```ssh
# GET origin BUILD.bazel path, it usually located in /home/thuy/.cache/bazel/_bazel_xxx/xxxx/external
bazel query @github_com_zergon321_reisen//:reisen --output=location
# create a new update BUILD.patch witch cgo, cdeps, clinkopts
diff -Naur thirdparty/reisen/BUILD.orig thirdparty/reisen/BUILD.need > thirdparty/reisen/BUILD.patch
```
Next update first 2 lines to point to correct BUILD path
```shell
--- github_com_zergon321_reisen/BUILD.bazel
+++ github_com_zergon321_reisen/BUILD.bazel
```

<span style="color:orange"> · Step 4.</span> Rebuild with patch

WORKSPACE
```shell
go_repository(
    name = "github_com_zergon321_reisen",
    importpath = "github.com/zergon321/reisen",
    patch_args = ["-p1"],
    build_file_proto_mode = "disable",
    patches = ["//thirdparty/reisen:BUILD.patch"],  # keep
    sum = "h1:rflBVtSoc4xXR3tXN+b2Hvk/Fb3qaJhgj+CJ/YWOYmI=",
    version = "v0.1.8",
)
```
OR repositories.bzl
```shell
go_repository(
        name = "github_com_zergon321_reisen",
        importpath = "github.com/zergon321/reisen",
        patch_args = ["-p1"],
        build_file_proto_mode = "disable",
        patches = ["//thirdparty/reisen:BUILD.patch"],  # keep
        sum = "h1:rflBVtSoc4xXR3tXN+b2Hvk/Fb3qaJhgj+CJ/YWOYmI=",
        version = "v0.1.8",
    )
```
Finally
```shell
bazel build @github_com_zergon321_reisen//:reisen
```

