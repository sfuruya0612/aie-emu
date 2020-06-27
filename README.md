# aie-emu

[![Build Status](https://travis-ci.org/sfuruya0612/aie-emu.svg?branch=master)](https://travis-ci.org/sfuruya0612/aie-emu)

## Description

aie-emu is CLI tools for get a list of AWS IAM Users.  
Using [aws-sdk-go](https://docs.aws.amazon.com/sdk-for-go/api/service/iam/) written by golang.  

aie-emu output some format.  
Now supported csv, Markdown, and Markdown Extra.  

## Install

- Go get

```bash
go get github.com/sfuruya0612/aie-emu
```

- Build from source

```bash
git clone https://github.com/sfuruya0612/aie-emu.git
cd aie-emu
make install
```

## Usage

### Help

```bash
$ aie-emu -h
NAME:
   aie-emu

USAGE:
   aie-emu [global options] command [command options] [arguments...]

VERSION:
   20.2.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --profile value, -p value  AWS credential (~/.aws/config) or read AWS_PROFILE environment variable (default: "default") [$AWS_PROFILE]
   --output value, -o value   (default: "stdout")
   --help, -h                 show help
   --version, -v              print the version
```

### Example

Set `AWS_PROFILE`  

- Export environment variable.(`export AWS_PROFILE=<Your Credential>`)
- Commands Option.(`aie-emu -p <Your Credential>`)

```bash
# stdout
$ aie-emu
```

```bash
# csv
$ aie-emu -o csv
```

```bash
# Markdown
$ aie-emu -o md
```

```bash
# Markdown Extra
$ aie-emu -o ex
```

## License

[MIT License](./LICENSE)
