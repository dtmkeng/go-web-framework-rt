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
### Result after run test
``` s
   Gin_github: 58560 Bytes
   Beego_github: 150920 Bytes
   Revel_github: 164240 Bytes
   Gorilamux_github: 1314160 Bytes
   Echo_github: 95992 Bytes
   Aero_github: 1403352 Bytes
   GoJSONREST_github: 141880 Bytes
   Buffalo_github: 1038904 Bytes
   Gin_static: 34064 Bytes
   Beego_staic: 98440 Bytes
   Echo_Static: 75072 Bytes
   Gorilamux_static: 583680 Bytes
   Aero_static: 784040 Bytes
   GoJSONRest_static: 136600 Bytes
   Buffalo_static: 629952 Bytes
   Revel: 112160 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/project/http_router
BenchmarkGin_Param               	20000000	        73.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param5              	10000000	       134 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param20             	 5000000	       330 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParamWrite          	10000000	       129 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param             	 2000000	       959 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_Param5            	 1000000	      1200 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_Param20           	 1000000	      2393 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_ParamWrite        	 1000000	      1044 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_Param             	   20000	     62507 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_Param5            	   20000	     69658 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_Param20           	   20000	     66962 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_ParamWrite        	   20000	     67087 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param        	 1000000	      2210 ns/op	    1280 B/op	      10 allocs/op
BenchmarkGorillaMux_Param5       	  500000	      3081 ns/op	    1344 B/op	      10 allocs/op
BenchmarkGorillaMux_Param20      	  200000	      6785 ns/op	    3451 B/op	      12 allocs/op
BenchmarkGorillaMux_ParamWrite   	 1000000	      2269 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_Param              	 1000000	      2334 ns/op	    1040 B/op	      12 allocs/op
BenchmarkEcho_Param5             	 1000000	      2062 ns/op	    1040 B/op	      12 allocs/op
BenchmarkEcho_Param20            	 1000000	      2087 ns/op	    1040 B/op	      12 allocs/op
BenchmarkEcho_ParamWrite         	 1000000	      2212 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param              	30000000	        55.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param5             	50000000	        41.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param20            	30000000	        41.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParamWrite         	30000000	        53.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param        	 1000000	      1348 ns/op	     649 B/op	      13 allocs/op
BenchmarkGOJSONREST_Param5       	 1000000	      2539 ns/op	    1097 B/op	      16 allocs/op
BenchmarkGOJSONREST_Param20      	  200000	      9228 ns/op	    4485 B/op	      20 allocs/op
BenchmarkGOJSONREST_ParamWrite   	 1000000	      2129 ns/op	    1129 B/op	      18 allocs/op
BenchmarkBuffalo_Param           	 1000000	      1836 ns/op	     976 B/op	      11 allocs/op
BenchmarkBuffalo_Param5          	 1000000	      1945 ns/op	     992 B/op	      11 allocs/op
BenchmarkBuffalo_Param20         	 1000000	      1922 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_ParamWrite      	 1000000	      1782 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_GithubStatic        	20000000	        90.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubParam         	10000000	       183 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll           	   50000	     36632 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_GithubStatic      	 1000000	      1051 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubParam       	 1000000	      1167 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubAll         	   10000	    244268 ns/op	   71457 B/op	     609 allocs/op
BenchmarkRevel_GithubStatic      	   20000	     63297 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubParam       	   20000	     62178 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubAll         	     100	  12779233 ns/op	 7278768 B/op	  166866 allocs/op
BenchmarkGorilaMux_GithubStatic  	  200000	     10675 ns/op	     976 B/op	       9 allocs/op
BenchmarkGorilaMux_GithubParam   	  200000	      6686 ns/op	    1296 B/op	      10 allocs/op
BenchmarkGorilaMux_GithubAll     	     500	   3798438 ns/op	  251660 B/op	    1994 allocs/op
BenchmarkEcho_GithubStatic       	20000000	        90.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubParam        	10000000	       169 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubAll          	   50000	     34996 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubStatic       	20000000	        60.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubParam        	10000000	       129 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubAll          	   50000	     26321 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONREST_GithubStatic 	 2000000	       904 ns/op	     329 B/op	      11 allocs/op
BenchmarkGoJSONREST_GithubParam  	 1000000	      1649 ns/op	     713 B/op	      14 allocs/op
BenchmarkGoJSONREST_GithubAll    	    5000	    320771 ns/op	  134371 B/op	    2737 allocs/op
BenchmarkBuffalo_GithubStatic    	   10000	    102829 ns/op	   34791 B/op	     360 allocs/op
BenchmarkBuffalo_GithubParam     	  100000	     18904 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_GithubAll       	     100	  20695185 ns/op	 7048496 B/op	   72882 allocs/op
BenchmarkGin_StaticAll           	   50000	     23926 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_StaticAll         	   10000	    175810 ns/op	   55264 B/op	     471 allocs/op
BenchmarkEcho_StaticAll          	  100000	     20974 ns/op	       0 B/op	       0 allocs/op
BenchmarkGorilaMux_StaticAll     	    2000	   1071377 ns/op	  153332 B/op	    1421 allocs/op
BenchmarkAero_StaticAll          	  100000	     15232 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONRest_StaticAll    	   10000	    168768 ns/op	   51653 B/op	    1727 allocs/op
BenchmarkBuffalo_StaticAll       	     100	  14758570 ns/op	 5370557 B/op	   49525 allocs/op
PASS
ok  	github.com/dtmkeng/project/http_router	118.153s
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