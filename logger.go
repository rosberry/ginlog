package ginlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const DebugValuesKey = "Debug"

// Logger returns the middleware for Gin, which provides detailed logging of requests and responses from the Gin Web server.
func Logger(debug bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Copy body for debug
		body := new(bytes.Buffer)
		if debug && c.Request.Body != nil {
			c.Request.Body = setTeeReadCloser(c.Request.Body, body)
		}

		// Copy response for debug
		response := new(bytes.Buffer)
		if debug && c.Writer != nil {
			c.Writer = setTeeGinResponseWriter(c.Writer, response)
		}

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		output := fmt.Sprintf("[GIN] %-15s [%v] %3d %s %s %13v\n",
			clientIP,
			end.UTC().Format("2006/01/02 15:04:05"),
			statusCode,
			method,
			path,
			latency,
		)
		if debug {
			userAgent := userAgentInfo(c)

			compactJSON := new(bytes.Buffer)
			if c.ContentType() == gin.MIMEJSON && json.Compact(compactJSON, body.Bytes()) == nil {
				body = compactJSON
			}

			if !strings.Contains(c.ContentType(), gin.MIMEMultipartPOSTForm) {
				output += fmt.Sprintf("[GIN-DEBUG][%s]\n %s%s\n", userAgent, fmtDebugValues(c), string(body.Bytes()))
			}
			output += fmt.Sprintf("[GIN-DEBUG] RESPONSE: %s\n", string(response.Bytes()))
		}
		output += c.Errors.ByType(gin.ErrorTypePrivate).String()
		fmt.Print(output)
	}
}

type teeReadCloser struct {
	r io.ReadCloser
	w io.Writer
}

func (t *teeReadCloser) Read(p []byte) (n int, err error) {
	n, err = t.r.Read(p)
	if n > 0 {
		t.w.Write(p[:n])
	}
	return
}

func (t *teeReadCloser) Close() error {
	return t.r.Close()
}

func setTeeReadCloser(r io.ReadCloser, w io.Writer) io.ReadCloser {
	return &teeReadCloser{r, w}
}

type teeGinResponseWriter struct {
	gin.ResponseWriter
	w io.Writer
}

func (t *teeGinResponseWriter) Write(b []byte) (int, error) {
	t.w.Write(b)
	return t.ResponseWriter.Write(b)
}

func setTeeGinResponseWriter(gRespW gin.ResponseWriter, w io.Writer) gin.ResponseWriter {
	return &teeGinResponseWriter{gRespW, w}
}

// AddDebugValue adds the specified string to the detailed log
func AddDebugValue(c *gin.Context, value string) {
	var values []string
	if v, ex := c.Get(DebugValuesKey); ex {
		values = append(*v.(*[]string), value)
	} else {
		values = []string{value}
	}
	c.Set(DebugValuesKey, &values)
}

func fmtDebugValues(c *gin.Context) string {
	var res string
	if v, ex := c.Get(DebugValuesKey); ex {
		for _, s := range *v.(*[]string) {
			res = res + "[" + s + "] "
		}
	}
	return res
}

func userAgentInfo(c *gin.Context) string {
	userAgent := c.Request.UserAgent()

	if v := c.GetHeader("App-Version"); v != "" {
		userAgent = fmt.Sprintf("%s App-Version: '%s'", userAgent, v)
	}

	if v := c.GetHeader("OS-Version"); v != "" {
		userAgent = fmt.Sprintf("%s OS-Version: '%s'", userAgent, v)
	}

	if v := c.GetHeader("Device"); v != "" {
		userAgent = fmt.Sprintf("%s Device: '%s'", userAgent, v)
	}

	if v := c.GetHeader("User-Timezone"); v != "" {
		userAgent = fmt.Sprintf("%s User-Timezone: '%s'", userAgent, v)
	}

	return userAgent
}