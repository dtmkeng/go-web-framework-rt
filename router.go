package main

import (
	"echo"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	"github.com/aerogo/aero"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beego"
	bc "github.com/astaxie/beego/context"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/mux"
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
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
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

// gorilamux
func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params["name"])
}

func loadGorillaMux(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.path, "{$1}"),
			h,
		).Methods(route.method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// echo
func echoHandler(c echo.Context) error {
	return nil
}
func echoHandlerWrite(c echo.Context) error {
	io.WriteString(c.Response(), c.Param("name"))
	return nil
}
func echoHandlerTest(c echo.Context) error {
	io.WriteString(c.Response(), c.Request().RequestURI)
	return nil
}

func loadEcho(routes []route) http.Handler {
	var h echo.HandlerFunc = echoHandler
	if loadTestHandler {
		h = echoHandlerTest
	}

	e := echo.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			e.GET(r.path, h)
		case "POST":
			e.POST(r.path, h)
		case "PUT":
			e.PUT(r.path, h)
		case "PATCH":
			e.PATCH(r.path, h)
		case "DELETE":
			e.DELETE(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return e
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	case "POST":
		e.POST(path, h)
	case "PUT":
		e.PUT(path, h)
	case "PATCH":
		e.PATCH(path, h)
	case "DELETE":
		e.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return e
}

// aero
func aeroHandler(c aero.Context) error {
	return nil
}

func aeroHandlerWrite(ctx aero.Context) error {
	io.WriteString(ctx.Response().Internal(), ctx.Get("name"))
	return nil
}
func aeroHandlerTest(ctx aero.Context) error {
	io.WriteString(ctx.Response().Internal(), ctx.Request().Path())
	return nil
}
func loadAero(routes []route) http.Handler {
	var h aero.Handler = aeroHandler
	if loadTestHandler {
		h = aeroHandlerTest
	}
	app := aero.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			app.Get(r.path, h)
		case "POST":
			app.Post(r.path, h)
		case "PUT":
			app.Put(r.path, h)
		case "PATCH":
			app.Router().Add(http.MethodPatch, r.path, h)
		case "DELETE":
			app.Delete(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return app
}
func loadAeroSingle(method, path string, h aero.Handler) http.Handler {
	app := aero.New()
	switch method {
	case "GET":
		app.Get(path, h)
	case "POST":
		app.Post(path, h)
	case "PUT":
		app.Put(path, h)
	case "PATCH":
		app.Router().Add(http.MethodPatch, path, h)
	case "DELETE":
		app.Delete(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	// }
	return app
}

// GORESTJSON
func goJSONRESTHandler(w rest.ResponseWriter, req *rest.Request) {}

func gOJSONRESTHandlerWrite(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.PathParam("name"))
}

func goJSONRESTHandlerTest(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.RequestURI)
}

func loadGOJSONREST(routes []route) http.Handler {
	h := goJSONRESTHandler
	if loadTestHandler {
		h = goJSONRESTHandlerTest
	}

	api := rest.NewApi()
	restRoutes := make([]*rest.Route, 0, len(routes))
	for _, route := range routes {
		restRoutes = append(restRoutes,
			&rest.Route{route.method, route.path, h},
		)
	}
	router, err := rest.MakeRouter(restRoutes...)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api.MakeHandler()
}

func loadGOJSONRESTSingle(method, path string, hfunc rest.HandlerFunc) http.Handler {
	api := rest.NewApi()
	router, err := rest.MakeRouter(
		&rest.Route{method, path, hfunc},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api.MakeHandler()
}

// 	buffalo
//  ENV ...
// var ENV = envy.Get("GO_ENV", "development")

func buffaloHandler(c buffalo.Context) error {
	return nil
}
func buffaloHandlerWrite(c buffalo.Context) error {
	io.WriteString(c.Response(), c.Param("name"))
	return nil
}
func buffaloHandlerTest(c buffalo.Context) error {
	io.WriteString(c.Response(), c.Request().RequestURI)
	return nil
}
func loadBuffalo(routes []route) http.Handler {
	h := buffaloHandler
	if loadTestHandler {
		h = buffaloHandlerTest
	}
	app := buffalo.New(buffalo.Options{})
	for _, r := range routes {
		switch r.method {
		case "GET":
			app.GET(r.path, h)
		case "POST":
			app.POST(r.path, h)
		case "PUT":
			app.PUT(r.path, h)
		case "PATCH":
			app.PATCH(r.path, h)
		case "DELETE":
			app.DELETE(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return app
}
func loadBuffaloSingle(method, path string, h buffalo.Handler) http.Handler {
	app := buffalo.New(buffalo.Options{})
	// buffalo.NewOptions()
	switch method {
	case "GET":
		app.GET(path, h)
	case "POST":
		app.POST(path, h)
	case "PUT":
		app.PUT(path, h)
	case "PATCH":
		app.PATCH(path, h)
	case "DELETE":
		app.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	// app.PreHandlers = http.Handler.ServeHTTP(nil, nil)
	return app
}

// func airHandler(req *air.Request, res *air.Response) error {
// 	return nil
// }
// func airHandlerWrite(req *air.Request, res *air.Response) error {
// 	// io.WriteString(res, s)
// 	return res.WriteString(req.Param("name").Value().String())
// }
// func airHandlerTest(req *air.Request, res *air.Response) error {
// 	return res.WriteString(req.Path)
// }

// // func (a *air.Air) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// // 	// a := air.New()
// // 	a.server.ServeHTTP(r, w)
// // }
// func loadAir(routes []route) http.Handler {
// 	h := airHandler
// 	if loadTestHandler {
// 		h = airHandlerTest
// 	}
// 	app := air.Default
// 	for _, r := range routes {
// 		switch r.method {
// 		case "GET":
// 			app.GET(r.path, h)
// 		case "POST":
// 			app.POST(r.path, h)
// 		case "PUT":
// 			app.PUT(r.path, h)
// 		case "PATCH":
// 			app.PATCH(r.path, h)
// 		case "DELETE":
// 			app.DELETE(r.path, h)
// 		default:
// 			panic("Unknow HTTP method: " + r.method)
// 		}
// 	}
// 	return app
// }
// func loadAirSingle(method, path string, h air.Handler) http.Handler {

// 	app := air.Default
// 	switch method {
// 	case "GET":
// 		app.GET(path, h)
// 	case "POST":
// 		app.POST(path, h)
// 	case "PUT":
// 		app.PUT(path, h)
// 	case "PATCH":
// 		app.PATCH(path, h)
// 	case "DELETE":
// 		app.DELETE(path, h)
// 	default:
// 		panic("Unknow HTTP method: " + method)
// 	}

// 	return app
// }
