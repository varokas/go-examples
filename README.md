# go-examples

Example template for building and testing golang project.

## Prerequisite
1. Read about where to put the source code [here](https://golang.org/doc/code.html). It is very important that the source is in the right place for Go to work correctly.
1. Read about the (very confusing) state of [package management](https://github.com/golang/go/wiki/PackageManagementTools) in go.

## Tools
1. Find a decent editor. Visual Studio Code with Go Plugin from [lukehoban](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go) works well. Provided that you install go correctly.
1. Install [dep](https://github.com/golang/dep) 

```
go get -u github.com/golang/dep/cmd/dep
```


## Developing 
1. First, clone the project to the right path 
(e.g. ~/go/github.com/varokas/go-examples)

2. Then Run a command to populate vendor/ folder
```
dep ensure
```

3. Run Unit Test
```
go test ./...
```

## Building in CI environment
We will just build in docker to avoid fiddling with $GOPATH.

## Misc
Adding a dependencies to the vendor/ project by

```
dep ensure -add "github.com/urfave/cli"
```