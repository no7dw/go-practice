go
### how to run

    go build xxx.go
    ./xxx
### build and run
    
    go run xxx.go

### unit test

    go test hello_test.go 
    go test

    unit-test$ go test
    good
    PASS
    ok      _/Users/dengwei/projects/github/go-practice/unit-test   0.004s


The suffix **_test** has special meaning in Go
be careful of naming your go file

### error handling
  
    wrong way (DRY!!):
 
    ```
    _, err = fd.Write(p0[a:b])
    if err != nil {
        return err
    }
    _, err = fd.Write(p1[c:d])
    if err != nil {
        return err
    }
    _, err = fd.Write(p2[e:f])
    if err != nil {
        return err
    }

    ```

    ref:
  - https://morsmachine.dk/error-handling
  - https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/11.1.html
  - https://www.goinggo.net/2014/10/error-handling-in-go-part-i.html
  - https://blog.golang.org/defer-panic-and-recover
  - https://blog.golang.org/errors-are-values

### concurrency

    best documents about concurrency for starter:  
    ref:
    http://studygolang.com/articles/1615 

### struct & pointer

    对于receiver 是指针的object 调用, Go会自动转换成指针

    [struct ref](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md)
    