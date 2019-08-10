package main

import (
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var benchRe *regexp.Regexp

const fiveColon = "/:a/:b/:c/:d/:e"
const fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
const fiveRoute = "/test/test/test/test/test"

const twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
const twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
const twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

func isTested(name string) bool {
	if benchRe == nil {
		// Get -test.bench flag value (not accessible via flag package)
		bench := ""
		for _, arg := range os.Args {
			if strings.HasPrefix(arg, "-test.bench=") {
				// ignore the benchmark name after an underscore
				bench = strings.SplitN(arg[12:], "_", 2)[0]
				break
			}
		}

		// Compile RegExp to match Benchmark names
		var err error
		benchRe, err = regexp.Compile(bench)
		if err != nil {
			panic(err.Error())
		}
	}
	return benchRe.MatchString(name)
}

func calcMem(name string, load func()) {
	if !isTested(name) {
		return
	}

	m := new(runtime.MemStats)

	// before
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	println("   "+name+":", after-before, "Bytes")
}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func benchRoutes(b *testing.B, router http.Handler, routes []route) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.method
			r.RequestURI = route.path
			u.Path = route.path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}
func BenchmarkGin_Param(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param5(b *testing.B) {
	router := loadGinSingle("GET", fiveColon, ginHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param20(b *testing.B) {
	router := loadGinSingle("GET", twentyColon, ginHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_ParamWrite(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// beego
func BenchmarkBeego_Param(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param5(b *testing.B) {
	router := loadBeegoSingle("GET", fiveColon, beegoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param20(b *testing.B) {
	router := loadBeegoSingle("GET", twentyColon, beegoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_ParamWrite(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkRevel_Param(b *testing.B) {
	router := loadRevelSingle("GET", "/user/:name", "RevelController.Handle")

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkRevel_Param5(b *testing.B) {
	router := loadRevelSingle("GET", fiveColon, "RevelController.Handle")

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkRevel_Param20(b *testing.B) {
	router := loadRevelSingle("GET", twentyColon, "RevelController.Handle")

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkRevel_ParamWrite(b *testing.B) {
	router := loadRevelSingle("GET", "/user/:name", "RevelController.HandleWrite")

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkGorillaMux_Param(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param5(b *testing.B) {
	router := loadGorillaMuxSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param20(b *testing.B) {
	router := loadGorillaMuxSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_ParamWrite(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", gorillaHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// echo
func BenchmarkEcho_Param(b *testing.B) {
	router := loadEchoSingle("GET", "/user/{name}", echoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param5(b *testing.B) {
	router := loadEchoSingle("GET", fiveBrace, echoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param20(b *testing.B) {
	router := loadEchoSingle("GET", twentyBrace, echoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_ParamWrite(b *testing.B) {
	router := loadEchoSingle("GET", "/user/{name}", echoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// Aero
func BenchmarkAero_Param(b *testing.B) {
	router := loadAeroSingle("GET", "/user/{name}", aeroHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkAero_Param5(b *testing.B) {
	router := loadAeroSingle("GET", fiveBrace, aeroHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkAero_Param20(b *testing.B) {
	router := loadAeroSingle("GET", twentyBrace, aeroHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkAero_ParamWrite(b *testing.B) {
	router := loadAeroSingle("GET", "/user/{name}", aeroHandlerTest)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// GoJSONREST
func BenchmarkGOJSONREST_Param(b *testing.B) {
	router := loadGOJSONRESTSingle("GET", "/user/:name", goJSONRESTHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGOJSONREST_Param5(b *testing.B) {
	handler := loadGOJSONRESTSingle("GET", fiveColon, goJSONRESTHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGOJSONREST_Param20(b *testing.B) {
	handler := loadGOJSONRESTSingle("GET", twentyColon, goJSONRESTHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGOJSONREST_ParamWrite(b *testing.B) {
	handler := loadGOJSONRESTSingle("GET", "/user/:name", goJSONRESTHandlerTest)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
}

// Buffalo

func BenchmarkBuffalo_Param(b *testing.B) {
	router := loadBuffaloSingle("GET", "/user/:name", buffaloHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkBuffalo_Param5(b *testing.B) {
	handler := loadBuffaloSingle("GET", fiveColon, buffaloHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkBuffalo_Param20(b *testing.B) {
	handler := loadBuffaloSingle("GET", twentyColon, buffaloHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkBuffalo_ParamWrite(b *testing.B) {
	handler := loadBuffaloSingle("GET", "/user/:name", buffaloHandlerTest)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
}

// Air
// func BenchmarkAir_Param(b *testing.B) {
// 	router := loadAirSingle("GET", "/user/:name", airHandler)

// 	r, _ := http.NewRequest("GET", "/user/gordon", nil)
// 	benchRequest(b, router, r)
// }
// func BenchmarkAir_Param5(b *testing.B) {
// 	handler := loadAirSingle("GET", fiveColon, airHandler)

// 	r, _ := http.NewRequest("GET", fiveRoute, nil)
// 	benchRequest(b, handler, r)
// }
// func BenchmarkAir_Param20(b *testing.B) {
// 	handler := loadAirSingle("GET", twentyColon, airHandler)

// 	r, _ := http.NewRequest("GET", twentyRoute, nil)
// 	benchRequest(b, handler, r)
// }
// func BenchmarkAir_ParamWrite(b *testing.B) {
// 	handler := loadAirSingle("GET", "/user/:name", airHandlerTest)

// 	r, _ := http.NewRequest("GET", "/user/gordon", nil)
// 	benchRequest(b, handler, r)
// }
