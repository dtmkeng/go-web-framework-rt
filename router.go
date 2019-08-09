package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	"github.com/astaxie/beego"
	bc "github.com/astaxie/beego/context"
	"github.com/gin-gonic/gin"
	"github.com/revel/pathtree"
	"github.com/revel/revel"
)

type route struct {
	method string
	path   string
}

var nullLogger *log.Logger
var loadTestHandler = false

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}
func (m *mockResponseWriter) WriteHeader(int) {}
func init() {

	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initBeego()
	initGin()
	initRevel()

}

// Gin
func ginHandle(_ *gin.Context) {}

func ginHandleWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Params.ByName("name"))
}

func ginHandleTest(c *gin.Context) {

	io.WriteString(c.Writer, c.Request.RequestURI)
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []route) http.Handler {
	h := ginHandle
	if loadTestHandler {
		h = ginHandleTest
	}

	router := gin.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, handle)
	return router
}
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}

// sbeego
func beegoHandler(ctx *bc.Context) {}

func beegoHandlerWrite(ctx *bc.Context) {
	ctx.WriteString(ctx.Input.Param(":name"))
}

func beegoHandlerTest(ctx *bc.Context) {
	ctx.WriteString(ctx.Request.RequestURI)
}

func initBeego() {
	beego.BConfig.RunMode = beego.PROD
	// beego
}

func loadBeego(routes []route) http.Handler {
	h := beegoHandler
	if loadTestHandler {
		h = beegoHandlerTest
	}

	re := regexp.MustCompile(":([^/]*)")
	app := beego.NewControllerRegister()
	for _, route := range routes {
		route.path = re.ReplaceAllString(route.path, ":$1")
		switch route.method {
		case "GET":
			app.Get(route.path, h)
		case "POST":
			app.Post(route.path, h)
		case "PUT":
			app.Put(route.path, h)
		case "PATCH":
			app.Patch(route.path, h)
		case "DELETE":
			app.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return app
}

func loadBeegoSingle(method, path string, handler beego.FilterFunc) http.Handler {
	app := beego.NewControllerRegister()
	switch method {
	case "GET":
		app.Get(path, handler)
	case "POST":
		app.Post(path, handler)
	case "PUT":
		app.Put(path, handler)
	case "PATCH":
		app.Patch(path, handler)
	case "DELETE":
		app.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return app
}

// RevelController ...
// Revel (Router only)
// In the following code some Revel internals are modelled.
// The original revel code is copyrighted by Rob Figueiredo.
// See https://github.com/revel/revel/blob/master/LICENSE
type RevelController struct {
	*revel.Controller
	router *revel.Router
}

//Handle ...
func (rc *RevelController) Handle() revel.Result {
	return revelResult{}
}

//HandleWrite ...
func (rc *RevelController) HandleWrite() revel.Result {
	return rc.RenderText(rc.Params.Get("name"))
}

//HandleTest ...
func (rc *RevelController) HandleTest() revel.Result {
	return rc.RenderText(rc.Request.GetRequestURI())
}

type revelResult struct{}

func (rr revelResult) Apply(req *revel.Request, resp *revel.Response) {}

func (rc *RevelController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Dirty hacks, do NOT copy!
	revel.MainRouter = rc.router

	upgrade := r.Header.Get("Upgrade")
	if upgrade == "websocket" || upgrade == "Websocket" {
		panic("Not implemented")
	} else {
		context := revel.NewGoContext(nil)
		context.Response.SetResponse(w)
		context.Request.SetRequest(r)

		c := revel.NewController(context)

		c.Request.WebSocket = nil
		// revel.Filters[0](c, revel.Filters[1:])
		if c.Result != nil {
			c.Result.Apply(c.Request, c.Response)
		} else if c.Response.Status != 0 {
			panic("Not implemented")
		}
		// Close the Writer if we can
		if w, ok := c.Response.GetWriter().(io.Closer); ok {
			w.Close()
		}
	}
}
func initRevel() {
	// Only use the Revel filters required for this benchmark
	// revel.AppLog =
	revel.Filters = []revel.Filter{
		revel.RouterFilter,
		revel.ParamsFilter,
		revel.ActionInvoker,
	}

	revel.RegisterController((*RevelController)(nil),
		[]*revel.MethodType{
			{
				Name: "Handle",
			},
			{
				Name: "HandleWrite",
			},
			{
				Name: "HandleTest",
			},
		})
}

var (
	appModule = &revel.Module{Name: "App"}
)

func loadRevel(routes []route) http.Handler {
	h := "RevelController.Handle"
	if loadTestHandler {
		h = "RevelController.HandleTest"
	}

	router := revel.NewRouter("")

	// parseRoutes
	var rs []*revel.Route
	for _, r := range routes {
		// revel.NewRoute(revel.appmo, method, path, action, fixedArgs, routesPath, line)
		rs = append(rs, revel.NewRoute(appModule, r.method, r.path, h, "", "", 0))
	}
	router.Routes = rs

	// updateTree
	router.Tree = pathtree.New()
	for _, r := range router.Routes {
		err := router.Tree.Add(r.TreePath, r)
		// Allow GETs to respond to HEAD requests.
		if err == nil && r.Method == "GET" {
			err = router.Tree.Add("/HEAD"+r.Path, r)
		}
		// Error adding a route to the pathtree.
		if err != nil {
			panic(err)
		}
	}

	rc := new(RevelController)
	rc.router = router
	return rc
}
func loadRevelSingle(method, path, action string) http.Handler {
	router := revel.NewRouter("")

	route := revel.NewRoute(appModule, method, path, action, "", "", 0)
	if err := router.Tree.Add(route.TreePath, route); err != nil {
		panic(err)
	}

	rc := new(RevelController)
	rc.router = router
	return rc
}
