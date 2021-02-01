# ginlog
logger for gin

## Usage
```go
import (
  "github.com/gin-gonic/gin"
  "github.com/rosberry/ginlog"
)

func main() {
  debugMode := true
  
  r := gin.Default()
	
  //Add ginlog as middleware
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
