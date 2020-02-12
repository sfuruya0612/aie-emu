# aie-emu

## Description

aie-emu is CLI tools for get a list of AWS IAM Users.  
Using [aws-sdk-go](https://docs.aws.amazon.com/sdk-for-go/api/service/iam/) written by golang.  

aie-emu output some format.  
Now supported csv, Markdown, and Markdown Extra.  

## Install

- Go get

``` sh
go get github.com/sfuruya0612/aie-emu
```

- Build from source

``` sh
git clone https://github.com/sfuruya0612/aie-emu.git
cd aie-emu
make install
```

## Usage

### Help

``` sh
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

#### Example

Set `AWS_PROFILE`  
- Export environment variable.(`export AWS_PROFILE=<Your Credential>`)
- Commands Option.(`aie-emu -p <Your Credential>`)

``` sh
# stdout
$ aie-emu
```

```
# csv
$ aie-emu -o csv
```

```
# Markdown
$ aie-emu -o md
```

```
# Markdown Extra
$ aie-emu -o ex
```
