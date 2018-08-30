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

### 自动下载项目的依赖包

    go get 项目会自动下载依赖包。但在git clone 项目后，具体的依赖包没安装，没有类似glide or 其他包管理的配置情况下（很多项目也并没有），可以

    ```
        go get -d -v ./...
    ```    
        -d标志只下载代码包，不执行安装命令；
        -v打印详细日志和调试日志。这里加上这个标志会把每个下载的包都打印出来；
        ./...这个表示路径，代表当前目录下所有的文件。

    [ref](http://blog.cyeam.com/golang/2018/06/19/gogetdep)
    [ref](https://golangtc.com/download/package)


### convert between x and y

    - convert between struct and byte array
    case : socket WriteMessage() ReadMessage()
    
    ```
        import "bytes"
        import "encoding/binary"

        var bin_buf bytes.Buffer
        binary.Write(&bin_buf, binary.LittleEndian, message)
    ```

    ```
        import "encoding/json"
        binArray, err := json.Marshal(messageStruct)

    ```

    ```
        import "encoding/json"
        type struct OrderBook {
            ...
        }
        var emp OrderBook
        err = json.Unmarshal(message, &emp)
    ```
    [convert struct to byte array](https://stackoverflow.com/questions/16330490/in-go-how-can-i-convert-a-struct-to-a-byte-array)
    (convert struct to byte array)[https://stackoverflow.com/questions/50660200/golang-convert-bytes-array-to-struct-with-a-field-of-type-byte]
    [json and go](https://blog.golang.org/json-and-go)

    - convert bewteen string and x

        - string to int/float/boolean

        ```
            i, err := strconv.Atoi("-42")
            s := strconv.Itoa(-42)
        ```

            strconv.parseInt
        ```
            import "strconv"
            b, err := strconv.ParseBool("true")
            f, err := strconv.ParseFloat("3.1415", 64)
            i, err := strconv.ParseInt("-42", 10, 64)
            u, err := strconv.ParseUint("42", 10, 64)
        ```

        - int/float/boolean to string
            
            ```
                import "fmt"
                a := 123
	            s := fmt.Sprintf("%d", a)
            ```

            ```
                s := strconv.Itoa(96)
                var n int64  = 97 
                s := strconv.Format(n, 10)

            ```
         [string convert]
         [convert](https://yourbasic.org/golang/convert-int-to-string/)   
         


