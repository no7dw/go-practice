# Capitals in struct
    because exported struct is need to be uppercase, so 
    
    ```
        type Person0 struct {
            name string
            age  int
        }}
    ```

    ```
        type Person struct {
            Name string
            Age  int
        }}
    ```
 the above struct is quite different:
    you can not access Person0 instance name

[private and public struct](https://til.hashrocket.com/posts/3512417fb0-private-vs-public-struct-members)    
[Exported identifiers](https://golang.org/ref/spec#Exported_identifiers)

