package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

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
