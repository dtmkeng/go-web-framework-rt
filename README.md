# http-router-benchmark-test
### Framework List
    - Gin
    - Beego
    - Echo
    - Go-JSON-REST
    - Aero
    - Buffalo
    - Gorila/Mux
    - Revel
### Result after run test week-7
``` s
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/go-web-framework-rt
BenchmarkGin_Error404        	 2000000	       681 ns/op	     400 B/op	       2 allocs/op
BenchmarkBeego_Error404      	 2000000	       728 ns/op	       8 B/op	       1 allocs/op
BenchmarkRevel_Error404      	   20000	     78786 ns/op	   35857 B/op	     822 allocs/op
BenchmarkGorillaMux_Error404 	 1000000	      1783 ns/op	     896 B/op	       8 allocs/op
BenchmarkEcho_Error404       	  500000	      2739 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Error404       	20000000	        59.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Error404 	  500000	      3879 ns/op	    1313 B/op	      21 allocs/op
BenchmarkBuffalo_Error404    	 1000000	      2261 ns/op	     976 B/op	      11 allocs/op
PASS
ok  	github.com/dtmkeng/go-web-framework-rt	15.593s

#Air
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/air
BenchmarkAir_Error404 	  500000	      3066 ns/op	     720 B/op	      13 allocs/op
PASS
ok  	github.com/dtmkeng/air	6.193s

#Goa
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/goa/http
BenchmarkGoa_Error404 	  500000	      3511 ns/op	     856 B/op	      15 allocs/op
PASS
ok  	github.com/dtmkeng/goa/http	1.918s


```
## Detail test
    - githubapi = 203
        GET = 131
        POST = 29
        PUT = 15
        PATCH = 28
    - staticapi
        GET =  157
        POST =  0
        PUT = 0
        PATCH = 0