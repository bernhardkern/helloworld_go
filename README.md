# helloworld_go

## Building
In any case:
```
go get -v github.com/bernhardkern/helloworld_go
go build github.com/bernhardkern/helloworld_go
```
then change dir to $GOPATH/src/github.com/bernhardkern/helloworld_go
### Windows
Use git bash (and avoid firewall permissions)

```
go build && ./helloworld_go.exe -user <mysql user> <mysql password> <mysql schema>
```
### Unix
```
go run main.go -user <mysql user> <mysql password> <mysql schema>
```
