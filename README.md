# gomodplay
Test mulitple GO module example.
 - Three modules
    - sampler - https://github.com/rahulhegde/gomodplay
    - greetings - https://github.com/rahulhegde/gomodplay/greetings
    - finance - https://github.com/rahulhegde/gomodplay/finance

 - Single mainapp -  https://github.com/rahulhegde/gomodplay/mainapp


## Tag & publish the module version - https://go.dev/doc/modules/publishing

* Make code changes and commit to the repository
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay$ git tag v1.0.2
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay$ git push origin v1.0.2
```

* Tag is not visible to other developers 

```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay
github.com/rahulhegde/gomodplay v1.0.0 v1.0.1 v1.0.2
```

* Make the tag visible to other developers by updating go module index (proxy)  
```developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay$ GOPROXY=proxy.golang.org go list -m github.com/rahulhegde/gomodplay@v1.0.2
github.com/rahulhegde/gomodplay v1.0.3
```

* Tag is visible now, to other developers 
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay
github.com/rahulhegde/gomodplay v1.0.0 v1.0.1 v1.0.2 v1.0.3
```

* Change the mainapp to use v1.0.3 - existing go.mod
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ cat go.mod 
module rahulhegde/mainapp

go 1.15

require github.com/rahulhegde/gomodplay v1.0.2
```

* Pull the new sampler package 
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go get github.com/rahulhegde/gomodplay@v1.0.3
go: downloading github.com/rahulhegde/gomodplay v1.0.3
```

* mainapp starts using v1.0.3 - changed go.mod
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ cat go.mod 

module rahulhegde/mainapp

go 1.15

require github.com/rahulhegde/gomodplay v1.0.3
```

* Run the mainapp to see changes 
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go build
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ ./mainapp 
finance: version:  103
```

## Mulitple modules definition in a single repository. Tag definition for finance and greetings modules

* Commit changes related to finance module 
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git add -A
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git commit -m "gomodplay/finance module - v1.0.2"
...1678a87..6061f7a  main -> main
```

* Tag and update modules index in the form: modulename/semanticversion (https://go.dev/doc/modules/managing-source)
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git tag finance/v1.0.2

developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ git push origin finance/v1.0.2
... * [new tag]         finance/v1.0.2 -> finance/v1.0.2

developer@rahul-Inspiron-7573:~/workspace/go-ws/src/github.com/rahulhegde/gomodplay/finance$ GOPROXY=proxy.golang.org go list -m github.com/rahulhegde/gomodplay/finance@v1.0.2
```

* Check listing of finance modules to developers
```
developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay
github.com/rahulhegde/gomodplay v1.0.0 v1.0.1 v1.0.2 v1.0.3 v1.0.4

developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay/greetings
github.com/rahulhegde/gomodplay/greetings v1.1.0 v1.1.1

developer@rahul-Inspiron-7573:~/workspace/go-ws/hello$ go list -m -versions github.com/rahulhegde/gomodplay/finance
github.com/rahulhegde/gomodplay/finance v1.0.0 v1.0.1 v1.0.2
```