package readbody

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

const (
	contextBody = "github.com/creasty/gin-contrib/readbody.Body"
)

// Read reads all bytes from the request body
// and initializes c.Request.Body with a new buffer.
// It's marked as DEPRECATED and it will be unexported in the future;
// Use a combination of Recorder and Get.
func Read(c *gin.Context) (body []byte) {
	body, _ = ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}

// Recorder is a middleware that reads the request body and stores it to the context
func Recorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(contextBody, Read(c))
		c.Next()
	}
}

// Get retrives stored data of the request body from the context
func Get(c *gin.Context) []byte {
	if v, ok := c.Get(contextBody); ok && v != nil {
		if err, ok := v.([]byte); ok {
			return err
		}
	}
	return nil
}
