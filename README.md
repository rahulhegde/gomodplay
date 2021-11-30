# gomodplay
Test mulitple GO module example 


# Publish the module version - https://go.dev/doc/modules/publishing

// make changes and commit the code to repository
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay$ git tag v1.0.2
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay$ git push origin v1.0.2
```

// tag is not visible 
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay
github.com/rahulhegde/gomodplay v1.0.0 v1.0.1 v1.0.2
```

// update the go module index (proxy)  
```developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay$ GOPROXY=proxy.golang.org go list -m github.com/rahulhegde/gomodplay@v1.0.2
github.com/rahulhegde/gomodplay v1.0.3
```

// tag is visible
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay
github.com/rahulhegde/gomodplay v1.0.0 v1.0.1 v1.0.2 v1.0.3
```

// change the mainapp to use v1.0.3
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ cat go.mod 
module rahulhegde/mainapp

go 1.15

require github.com/rahulhegde/gomodplay v1.0.2
```

developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go get github.com/rahulhegde/gomodplay@v1.0.3
go: downloading github.com/rahulhegde/gomodplay v1.0.3
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ cat go.mod 
```
module rahulhegde/mainapp

go 1.15

require github.com/rahulhegde/gomodplay v1.0.3
```

developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go build
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ ./mainapp 
finance: version:  103

# Mulitple modules definition in a single repository. Tag definition for finance and greetings modules
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git add -A
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git commit -m "gomodplay/finance module - v1.0.2"
...1678a87..6061f7a  main -> main
```

* tag format modulename/semanticversion (https://go.dev/doc/modules/managing-source)
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git tag finance/v1.0.2
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git push origin finance/v1.0.2
... * [new tag]         finance/v1.0.2 -> finance/v1.0.2
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ GOPROXY=proxy.golang.org go list -m github.com/rahulhegde/gomodplay/finance@v1.0.2
```

* all module versioning listed 
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay
github.com/rahulhegde/gomodplay v1.0.0 v1.0.1 v1.0.2 v1.0.3 v1.0.4
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay/greetings
github.com/rahulhegde/gomodplay/greetings v1.1.0 v1.1.1
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay/finance
github.com/rahulhegde/gomodplay/finance v1.0.0 v1.0.1 v1.0.2
```
