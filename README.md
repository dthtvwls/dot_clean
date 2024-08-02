# dot_clean

A cross-platform utility to clean "[AppleSingle/AppleDouble](https://en.wikipedia.org/wiki/AppleSingle_and_AppleDouble_formats)" metadata files,
also known as "dot underscore" or "dot underbar" files.

For example:
```sh
$ ls -a | cat
> .
> ..
> ._favicon.ico
> ._index.html
> ._robots.txt
> favicon.ico
> index.html
> robots.txt
```

These files are created by macOS on non-Apple filesystems, but they can be a nuisance and in many cases it's desirable to get rid of them.

This implementation aims to mimic the specific behavior of Apple's `dot_clean` command using the `-m` switch,
which removes these files from a given directory and its children.

See [DOT_CLEAN(1)](https://keith.github.io/xcode-man-pages/dot_clean.1.html).

## Getting Started

The suggested method of installation is to use Go's tooling to download and build from source.

### Prerequisites

* [Go](https://go.dev/dl/)

### Installation

```sh
go install github.com/dthtvwls/dot_clean@latest
```

The `dot_clean` binary will be installed to `~/go/bin`.

## Usage

To recursively remove all dot underscore files from a directory:
```sh
dot_clean -m /path/to/directory
```

Windows paths also work:
```sh
dot_clean -m C:\path\to\directory
```
