package puppet

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3,pcs[:]) //

	var str strings.Builder
	str.WriteString(message + "\nTrace back:")
	for _,pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file,line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d",file,line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s",err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal server error")
			}
		}()
		c.Next()
	}
}
