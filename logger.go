package puppet

import (
	"fmt"
	"log"
	"time"
)

// Logger global midlleware
func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		fmt.Println(c.StatusCode)
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
