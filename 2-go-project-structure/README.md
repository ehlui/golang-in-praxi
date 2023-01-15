# Creating a project
Using go v1.19 and its official [guides](https://go.dev/doc/code).

## Steps 

> go init MY_MODULE_NAME


> go install MODULE_NAME_FROM_GO.MOD

This way we build and install our **hey.go** program.

This command builds the hello command, producing an executable binary. 
It then installs that binary as **$HOME/go/bin/hello** (or, under Windows, **%USERPROFILE%\go\bin\hello.exe**).

The install directory is controlled by the GOPATH and GOBIN [environment variables](https://pkg.go.dev/cmd/go#hdr-Environment_variables).

- If GOBIN is set, binaries are installed to that directory. 
- If GOPATH is set, binaries are installed 
to the bin subdirectory of the first directory in the GOPATH list.
- Otherwise, binaries are installed to the bin subdirectory of the default GOPATH (**$HOME/go** or **%USERPROFILE%\go**).

## Run

Binaries destination can be set by the **GOBIN** environment variable.

Now we can copy our binary and run it anywhere. 

### Linux-like
> chmod +x gopath/basicus

> cd gopath && ./OUR_GO_BINARY

### Windows

>  %USERPROFILE%\go\bin\hey.exe
