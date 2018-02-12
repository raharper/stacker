# stacker [![Build Status](https://travis-ci.org/anuvu/stacker.svg?branch=master)](https://travis-ci.org/anuvu/stacker)

Stacker is a tool for building OCI images via a declarative yaml format.

### Hacking

Stacker requires go 1.9, on Ubuntu you can get that with:

    sudo apt-add-repository ppa:gophers/archive
    sudo apt update
    sudo apt install golang-1.9
    export PATH=$PATH:/usr/lib/go-1.9/bin

And uses glide for dependency management, which unfortunately is installed with
curl piped to sh as follows:

    curl https://glide.sh/get | sh

Stacker also has the following build dependencies:

    sudo apt install libacl1-dev

The test suite requires `jq`, which can be installed on Ubuntu:

    sudo apt install jq

And `umoci`, which can be installed with:

    go get github.com/openSUSE/umoci/cmd/umoci

And can be run with

    cd test
    sudo -E ./main.sh

It will exit 0 on failure. There are several environment variables available:

1. `STACKER_KEEP` keeps the base layers and files, so you don't need to keep
   downloading them
1. `STACKER_INSPECT` stops the test suite before cleanup, so you can inspect
   the failure

### Install

You'll need both the stacker and stackermount binaries:

    go install github.com/anuvu/stacker/stacker
    go install github.com/anuvu/stacker/stackermount

Stacker also depends on skopeo for some operations; you can install skopeo on
ubuntu with:

    sudo apt-add-repository ppa:projectatomic/ppa
    sudo apt update
    sudo apt install skopeo

## The `stacker.yaml` file

The basic driver of stacker is the stackerfile, canonically named
`stacker.yaml`. The stackerfile describes a set of OCI manifests to build. For
example:

    centos:
		from:
			type: docker
			url: docker://centos:latest
        run: echo meshuggah rocks > /etc/motd

Describes a manifest which is the latest Centos image from the Docker hub.
There are various other directives which users can use to 

1. `from`: this describes the base image that stacker will start from. You can
   either start from some other image in the same stackerfile, a Docker image,
   or a tarball.
1. `import`: A set of files to download or copy into the container. Stacker
   will put these files at `/stacker`, which will be automatically cleaned up
   after the commands in the `run` section are run and the image is finalized.
   URLs (`http://example.com/file.txt`), files (`/path/to/file`), and files
   from other stacker layers (`stacker:///output/my.exe`) are all supported.
1. `run`: This is the set of commands to run in order to build the image; they
   are run in a user namespaced container, with the set of files imported
   available in `/stacker`.
1. `environment`, `labels`, `working_dir`, `volumes`, `cmd`, `entrypoint`:
   these all correspond exactly to the similarly named bits in the [OCI image
   config spec](https://github.com/opencontainers/image-spec/blob/master/config.md#properties),
   and are available for users to pass things through to the runtime environment
   of the image.
1. `full_command`: because of the odd behavior of `cmd` and `entrypoint` (and
   the inherited nature of these from previous stacker layers), `full_command`
   provides a way to set the full command that will be executed in the image,
   clearing out any previous `cmd` and `entrypoint` values that were set in the
   image.
1. `build_only`: indicates whether or not to include this layer in the final
   OCI image. This can be useful in conjunction with an import from this layer
   in another image, if you want to isolate the build environment for a binary
   but not include all of its build dependencies.
