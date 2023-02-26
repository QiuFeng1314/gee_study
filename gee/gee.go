package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

var (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 支持中间件
	parent      *RouterGroup  // 支持嵌套
	engine      *Engine       // 所有的分组共享一个Engine
}

type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup // 存储所有的分组
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func New() (engine *Engine) {
	engine = &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return
}

// Use 使用中间件，就是把中间件加到队列中
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	var middlewares []HandlerFunc

	for _, v := range engine.groups {
		if strings.HasPrefix(req.URL.Path, v.prefix) {
			middlewares = append(middlewares, v.middlewares...)
		}
	}
	ctx := newContext(resp, req)
	ctx.handlers = middlewares
	ctx.engine = engine
	engine.router.handle(ctx)
}

func (group *RouterGroup) Group(prefix string) (newGroup *RouterGroup) {
	engine := group.engine
	newGroup = &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return
}

func (group *RouterGroup) addRouter(method, pattern string, handler HandlerFunc) {
	group.engine.router.addRouter(method, group.prefix+pattern, handler)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRouter(GET, pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRouter(POST, pattern, handler)
}

func (group *RouterGroup) Put(pattern string, handler HandlerFunc) {
	group.addRouter(PUT, pattern, handler)
}

func (group *RouterGroup) Delete(pattern string, handler HandlerFunc) {
	group.addRouter(DELETE, pattern, handler)
}

func (engine *Engine) Run(port string) (err error) {
	err = http.ListenAndServe(port, engine)
	if err == nil {
		log.Printf("gee start success! server port is %q...", port)
	}
	return
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(relativePath, group.prefix)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(ctx *Context) {
		param := ctx.Param("filepath")
		if _, err := fs.Open(param); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(ctx.Resp, ctx.Req)
	}
}

func (group *RouterGroup) Static(relativePath, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.Get(urlPattern, handler)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
