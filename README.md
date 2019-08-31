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
### Result after run test week-3
``` s
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/go-web-framework-rt
BenchmarkGin_Param             	20000000	        70.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param           	 2000000	       779 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param           	   30000	     50407 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param      	 1000000	      1760 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_Param            	 1000000	      1729 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param            	30000000	        46.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param      	 1000000	      1079 ns/op	     649 B/op	      13 allocs/op
BenchmarkBuffalo_Param         	 1000000	      1343 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_Param5            	10000000	       129 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param5          	 1000000	      1222 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param5          	   30000	     53357 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param5     	  500000	      2736 ns/op	    1344 B/op	      10 allocs/op
BenchmarkEcho_Param5           	 1000000	      1801 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param5           	50000000	        33.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param5     	 1000000	      2117 ns/op	    1097 B/op	      16 allocs/op
BenchmarkBuffalo_Param5        	 1000000	      1387 ns/op	     992 B/op	      11 allocs/op
BenchmarkGin_Param20           	 5000000	       324 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param20         	 1000000	      2203 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param20         	   30000	     52427 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param20    	  200000	      6231 ns/op	    3452 B/op	      12 allocs/op
BenchmarkEcho_Param20          	 1000000	      1791 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param20          	50000000	        35.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param20    	  200000	      7657 ns/op	    4485 B/op	      20 allocs/op
BenchmarkBuffalo_Param20       	 1000000	      1478 ns/op	    1008 B/op	      11 allocs/op
BenchmarkGin_ParamWrite        	20000000	       116 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_ParamWrite      	 2000000	       977 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_ParamWrite      	   30000	     52407 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_ParamWrite 	 1000000	      1925 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_ParamWrite       	 1000000	      1796 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_ParamWrite       	30000000	        45.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_ParamWrite 	 1000000	      1594 ns/op	    1129 B/op	      18 allocs/op
BenchmarkBuffalo_ParamWrite    	 1000000	      1335 ns/op	     976 B/op	      11 allocs/op
PASS
ok  	github.com/dtmkeng/go-web-framework-rt	56.575s

#Air
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/air
BenchmarkAir_Param      	 2000000	       836 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_Param5     	 2000000	       958 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_Param20    	 1000000	      1410 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_ParamWrite 	  300000	      6358 ns/op	    2784 B/op	      31 allocs/op
PASS
ok  	github.com/dtmkeng/air	13.486s

# Goa
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