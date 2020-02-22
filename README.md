
go-restful
===================


## <i class="icon-file"></i> Install Go

Take Mac as example, you can choose 

Building from source

or

macOS package installer


> [Download](https://golang.org/dl/)


### <i class="icon-file"></i> Test installation

Open Visual Studio Code (or other IDEs), 

```
$ mkdir hello
$ cd hello/
$ go mod init example.com/user/hello
go: creating new go.mod: module example.com/user/hello
$ cat go.mod
module example.com/user/hello

go 1.13
```

Then run

```
$ go run hello.go
Hello, world.
```

