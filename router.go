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

	// initBeego()
	initGin()

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
