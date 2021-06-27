# protoenv

version manager for protocol buffer tools



# Usage

## protoc

subcommands

```
Usage:
  protoenv protoc [command]

Available Commands:
  global      Set or show the global Go version
  install     install specified version
  version     Show the current protoc version and its origin
  versions    List all Go versions available to protoenv

Flags:
  -h, --help   help for protoc

Global Flags:
      --config string   config file (default is $HOME/.protoenv/config.yaml)
```

- show protoc versions

```
$ protoenv protoc install -l
v3.17.3
v3.17.2
v3.17.1
...
```

- install version

```
$ protoenv protoc install v3.17.3
downloading https://github.com/protocolbuffers/protobuf/releases/download/v3.17.3/protoc-3.17.3-osx-x86_64.zip ..
installed protoc ~/.protoenv/protoc/versions/v3.17.3/protoc
set global version: v3.17.3
```

- set global version

```
$ protoenv protoc global v3.17.3
set global version: v3.17.3
```
