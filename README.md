
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

## <i class="icon-file"></i> REST-ful APIs


### <i class="icon-file"></i>  Project Structure

![enter image description here](https://github.com/VickyFengYu/go-restful/blob/master/image/project-structure.jpg?raw=true)


### <i class="icon-file"></i>  Run App

```
$ cd demo/
$ go run main.go
Endpoint Hit: indexPage
Endpoint Hit: returnAllArticles
Route variables map[id:1]

```

### <i class="icon-file"></i> IndexPage

http://localhost:10000/

![enter image description here](https://github.com/VickyFengYu/go-restful/blob/master/image/index-page.jpg?raw=true)



### <i class="icon-file"></i> ReturnAllArticles

http://localhost:10000/articles

![enter image description here](https://github.com/VickyFengYu/go-restful/blob/master/image/all-articles.jpg?raw=true)


### <i class="icon-file"></i> ReturnSingleArticle by id

http://localhost:10000/article/1

![enter image description here](https://github.com/VickyFengYu/go-restful/blob/master/image/query-by-id.jpg?raw=true)



## <i class="icon-file"></i> Problems Encountered

### <i class="icon-file"></i>  1

```
$ go build
go: finding github.com/VickyFengYu/go-restful/article latest
go: downloading github.com/VickyFengYu/go-restful/article v0.0.0-20200223040213-949d1e9f127a
verifying github.com/VickyFengYu/go-restful/article@v0.0.0-20200223040213-949d1e9f127a: github.com/VickyFengYu/go-restful/article@v0.0.0-20200223040213-949d1e9f127a: reading https://sum.golang.org/lookup/github.com/!vicky!feng!yu/go-restful/article@v0.0.0-20200223040213-949d1e9f127a: 410 Gone
```

> **Soulution**

Add the following codes in DEMO/go.mod

```
require github.com/VickyFengYu/go-restful/article v0.0.0
replace github.com/VickyFengYu/go-restful/article v0.0.0 => ./article
```

and add 

```
article "github.com/VickyFengYu/go-restful/article"

```

in DEMO/main.go

> Note: use  ```article "./article"``` in DEMO/main.go doesn't work, will cause error
>
> ```
> $ go build
> build _/Users/vicky/Documents/go/demo/article: cannot find module for path _/Users/vicky/> Documents/go/demo/article
> ```


### <i class="icon-file"></i>  2. call function in another package


> NOTE: This func must be Exported, <MARK>Capitalized</MARK>, and comment added.


```
//In DEMO/article/article.go
func IndexPage(w http.ResponseWriter, r *http.Request) {}

```

```
//In DEMO/main.go
myRouter.HandleFunc("/", article.IndexPage)

```


