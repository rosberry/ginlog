# ginlog

Middleware logger for [gin](https://github.com/gin-gonic/gin).

Write the logs: client IP address, execution time, request and response.

### Output example:

```
[GIN] ::1             [2021/02/01 10:21:15] 200 POST /v1/auth   125.91203ms
[GIN-DEBUG] {"device_id":"s1"}
[GIN-DEBUG] RESPONSE: {"result":true,"token":"49f7174a-a2a7-4a17-5093-12w7879cc6b4","user":{"id":4,"name":"A","role":"admin","photo":null,"auth":"none"}}
```

## Usage

1. Download ginlog by using:
```sh
$ go get github.com/rosberry/ginlog
```
2. Import it in your code:
```go
import "github.com/rosberry/ginlog"
```
3. Add ginlog.Logger(debugMode) as middleware in gin, debugMode is *bool* parameter to show more info in output.

### Example

```go
import (
  "github.com/gin-gonic/gin"
  "github.com/rosberry/ginlog"
)

func main() {
  debugMode := true
  
  r := gin.Default()
	
  // Add ginlog as middleware
  r.GET("/ping", ginlog.Logger(debugMode), func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  
  r.Run()
}
```

## About

<img src="https://github.com/rosberry/Foundation/blob/master/Assets/full_logo.png?raw=true" height="100" />

This project is owned and maintained by [Rosberry](http://rosberry.com). We build mobile apps for users worldwide 🌏.

Check out our [open source projects](https://github.com/rosberry), read [our blog](https://medium.com/@Rosberry) or give us a high-five on 🐦 [@rosberryapps](http://twitter.com/RosberryApps).

## License

This project is available under the MIT license. See the LICENSE file for more info.
