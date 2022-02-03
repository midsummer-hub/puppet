package puppet

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer      http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params map[string]string
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:w,
		Req:r,
		Path:r.URL.Path,
		Method:r.Method,
	}
}

func (c *Context)Param(key string) string {
	value ,_:= c.Params[key]
	return value
}

func (c Context) PostForm(key string)string {
	return c.Req.FormValue(key)
}

func (c Context) Query(key string)string {
	return c.Req.URL.Query().Get(key)
}
func (c Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c Context)String(code int,format string,values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,values...)))
}

func (c Context) JSON(code int ,obj interface{}) {
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c Context) Data(code int,data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c Context)HTML(code int,html string) {
	c.SetHeader("Content-Type","text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}



