// Package puppet TODO
package puppet

import (
	"net/http"
)

// HandlerFunc Request handler
type HandlerFunc func(*Context)

type (
	// RouterGroup RouterGroup
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc //support middleware
		parent      *RouterGroup  //support nesting
		engine      *Engine       // all groups share a Engine instance
	}
	// Engine Engine
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// Group Create a new RouterGroup
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// New constructor
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET GET request method
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)

}

// POST POST request method
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run Start the HTTP server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP A Handler responds to an HTTP request.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
