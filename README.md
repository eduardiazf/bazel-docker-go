# bazel-docker-go
Monorepo Golang with Bazel and Docker

You will learn to use Bazel and Docker as an option for creating monorail in Go

Binary building and dependency management is one of the most frustrating parts of code development as well as having to manage multiple projects with different building tools.

## Documentation

### What is Bazel?
Bazel is a free software tool that allows the automation of software construction and testing.

### Why Bazel?
<strong>Speed up your builds and tests:</strong>
With advanced local and distributed caching, optimized dependency analysis and parallel execution, you get fast and incremental compilations.

<strong>Scalable:</strong>
Bazel helps you scale your organization, codebase and Continuous Integration system. It handles codebases of any size, in multiple repositories or a huge monorepo.

<strong>One tool, multiple languages:</strong>
You can build and test multiple languages with a single tool in a monorail

---
### First Steps
>We will create a project in Go with these commands:
```
mkdir -p bazel-docker-go
go mod init bairesapp
```
>We will create the directories according to the image.

![directories](https://github.com/eduardiazf/bazel-docker-go/blob/main/images/bazel-golang-docker-directories.PNG?raw=true)

>packages/hello/main.go
```
package main

import (
    "bairesapp/packages/shared"

    "fmt"
)

func  main() {
	fmt.Printf("Hellooo!!! Sum: %d", shared.Sum(2, 2))
}
```
>packages/shared/sum.go
```
package shared

func Sum(a, b int64) int64 {
	return a + b
}
```
>packages/shared/sum_test.go
```
package shared

import  "testing"

func  TestSum(t *testing.T) {
	result := Sum(2, 2)

	if result == 4 {
		t.Logf("Sum: 2 + 2 = %d", result)
	}
}
```
>Configure the root BUILD.bazel
```
load("@bazel_gazelle//:def.bzl", "gazelle")
# gazelle:prefix bairesapp
gazelle(name = "gazelle")
```
>Configure the WORKSPACE
```
workspace(name = "bairesapp")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7904dbecbaffd068651916dce77ff3437679f9d20e1a7956bff43826e7645fcc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "222e49f034ca7a1d1231422cdb67066b885819885c356673cb1f72f748a3c9d4",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036",
    strip_prefix = "rules_docker-0.14.4",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.4/rules_docker-v0.14.4.tar.gz"],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.15.5")

gazelle_dependencies()

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")

pip_deps()

load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")

# docker image
container_pull(
    name = "alpine_linux_amd64",
    digest = "sha256:cf35b4fa14e23492df67af08ced54a15e68ad00cac545b437b1994340f20648c",
    registry = "index.docker.io",
    repository = "library/alpine",
    tag = "3.8",
)

load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")

_go_image_repos()
```
Now you can run the following code to generate BUILD.bazel

```
bazel run //:gazelle
```
This will create our BUILD.bazel and we can use it to update it.

>packages/hello/BUILD.bazel

```
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")

# Included binary in docker container
go_image(
	name = "docker",
	base = ":image",
	embed = [":hello_lib"],
)

# Docker container
container_image(
	name = "image",
	base = "@alpine_linux_amd64//image"
)

go_library(
	name = "hello_lib",
	srcs = ["main.go"],
	importpath = "bairesapp/packages/hello",
	visibility = ["//visibility:private"],
	deps = ["//packages/shared"],
)

go_binary(
	name = "hello",
	embed = [":hello_lib"],
	visibility = ["//visibility:public"],
)
```
###  Useful commands
> run go binary
```
bazel run //packages/hello:hello
```
>run go test
```
bazel test //packages/shared:shared_test
```
>run binary inside docker container
```
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //packages/hello:docker
```
## Sources
* [https://bazel.build/](https://bazel.build/)
* [https://github.com/bazelbuild/rules_go](https://github.com/bazelbuild/rules_go)
* [https://github.com/bazelbuild/rules_docker](https://github.com/bazelbuild/rules_docker)
* https://github.com/bazelbuild/bazel-gazelle

