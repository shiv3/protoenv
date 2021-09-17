# protoenv

version manager for protocol buffer tools


# Install

```
go install github.com/shiv3/protoenv@v0.0.3
```

# Usage

## protoc

subcommands

```
Usage:
  protoenv protoc [command]

Available Commands:
  global      Set or show the global protoc version
  init        Set or show the global Go version%!(EXTRA string=protoc)
  install     install specified version
  local       Set or show the local protoc version
  version     Show the current protoc version and its origin
  versions    List all protoc versions available to protoenv

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
