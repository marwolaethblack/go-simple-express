# Go simple express

Go simple express is a wrapper to the basic go http server to provide some express like syntax

Usage:
``` go
var app epress.App

//Get requests
app.Get("/", MiddlewareHandler)

//Post requests
app.Post("/post", MiddlewareHandler)

//Delete requests
app.Delete("/delete", MiddlewareHandler)

app.Run(":8080")
```
And so on for PUT PATCH...

You can chain middleware in your request handlers
```go
app.Delete("/delete", MiddlewareHandler1, MiddlewareHandler2, MiddlewareHandler3)
```


Your MiddlewareHandler func needs to be in the format
```go
func(w http.ResponseWriter, req *http.Request, stop func())
```
All these paramaters (w, req, stop) will be passed in for you and you will have access to them in your function

MiddleWareHandler Example
```go
type User struct {
  Fname string
  Lname string
}

func getUsers(w http.ResponseWriter, req *http.Request, stop func(message string)) {
  userdata := User{"John", "Smith"}
  express.GzipJSON(w, userdata)
}
```

`stop(message string)` Will stop the middleware chain execution but will not stop the current function, use an if else block

stop() Example
```go
type User struct {
  Fname string
  Lname string
}

func getUsers(w http.ResponseWriter, req *http.Request, stop func(message string)) {
  userdata := User{"John", "Smith"}
  
  
  
  if !authorized() {
    stop("Unauthorized access")
  } else {
    express.GzipJSON(w, userdata)
  }
  
}

func main() {
  var app express.App
  app.Run(":8080")
}
```

You can set headers, gzip, encode to json and send your data to the writer using the `GzipJSON(w http.ResponseWriter, data interface{})` function
it accepts anything encodable to JSON with json.Marshall

Lastly this was made just for fun and practice and should not be taken seriously
