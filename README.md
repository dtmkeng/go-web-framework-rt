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
### Result after run test week-6-json
``` s
   Gin_github: 58480 Bytes
   Beego_github: 150920 Bytes
   Revel_github: 164448 Bytes
   Gorilamux_github: 1314160 Bytes
   Echo_github: 100168 Bytes
   Aero_github: 1403864 Bytes
   GoJSONREST_github: 142120 Bytes
   Buffalo_github: 1038888 Bytes
   Gin_static: 34064 Bytes
   Beego_staic: 98440 Bytes
   Echo_Static: 78704 Bytes
   Gorilamux_static: 583680 Bytes
   Aero_static: 784040 Bytes
   GoJSONRest_static: 136760 Bytes
   Buffalo_static: 629952 Bytes
   Revel: 111328 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/go-web-framework-rt
BenchmarkGin_Param               	 1000000	      1030 ns/op	     528 B/op	       6 allocs/op
BenchmarkBeego_Param             	  500000	      2917 ns/op	    1592 B/op	      19 allocs/op
BenchmarkRevel_Param             	   30000	     53281 ns/op	   35857 B/op	     822 allocs/op
BenchmarkGorillaMux_Param        	 1000000	      2181 ns/op	    1296 B/op	      11 allocs/op
BenchmarkEcho_Param              	 1000000	      1775 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param              	30000000	        45.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param        	  500000	      2843 ns/op	    1721 B/op	      26 allocs/op
BenchmarkBuffalo_Param           	 1000000	      1408 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_Param5              	 1000000	      1075 ns/op	     528 B/op	       6 allocs/op
BenchmarkBeego_Param5            	  500000	      3029 ns/op	    1592 B/op	      19 allocs/op
BenchmarkRevel_Param5            	   30000	     52821 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param5       	  500000	      2874 ns/op	    1360 B/op	      11 allocs/op
BenchmarkEcho_Param5             	 1000000	      1756 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param5             	50000000	        34.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param5       	  500000	      3735 ns/op	    2169 B/op	      29 allocs/op
BenchmarkBuffalo_Param5          	 1000000	      1440 ns/op	     992 B/op	      11 allocs/op
BenchmarkGin_Param20             	 1000000	      1479 ns/op	     528 B/op	       6 allocs/op
BenchmarkBeego_Param20           	  300000	      4320 ns/op	    1592 B/op	      19 allocs/op
BenchmarkRevel_Param20           	   30000	     52845 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param20      	  200000	      6253 ns/op	    3468 B/op	      13 allocs/op
BenchmarkEcho_Param20            	 1000000	      1739 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param20            	50000000	        35.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param20      	  200000	      9209 ns/op	    5559 B/op	      33 allocs/op
BenchmarkBuffalo_Param20         	 1000000	      1500 ns/op	    1008 B/op	      11 allocs/op
BenchmarkGin_GithubStatic        	 1000000	      1003 ns/op	     528 B/op	       6 allocs/op
BenchmarkGin_GithubParam         	 1000000	      1142 ns/op	     528 B/op	       6 allocs/op
BenchmarkGin_GithubAll           	    5000	    240588 ns/op	  107192 B/op	    1218 allocs/op
BenchmarkBeego_GithubStatic      	  500000	      2857 ns/op	    1592 B/op	      19 allocs/op
BenchmarkBeego_GithubParam       	  500000	      2931 ns/op	    1592 B/op	      19 allocs/op
BenchmarkBeego_GithubAll         	    2000	    617764 ns/op	  323194 B/op	    3857 allocs/op
BenchmarkRevel_GithubStatic      	   30000	     52183 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubParam       	   30000	     52611 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubAll         	     100	  10542315 ns/op	 7278931 B/op	  166866 allocs/op
BenchmarkGorilaMux_GithubStatic  	  200000	     10010 ns/op	     992 B/op	      10 allocs/op
BenchmarkGorilaMux_GithubParam   	  200000	      6357 ns/op	    1312 B/op	      11 allocs/op
BenchmarkGorilaMux_GithubAll     	     500	   3548119 ns/op	  254946 B/op	    2197 allocs/op
BenchmarkEcho_GithubStatic       	 2000000	       908 ns/op	     480 B/op	       5 allocs/op
BenchmarkEcho_GithubParam        	 1000000	      1046 ns/op	     480 B/op	       5 allocs/op
BenchmarkEcho_GithubAll          	   10000	    216711 ns/op	   97445 B/op	    1015 allocs/op
BenchmarkAero_GithubStatic       	 2000000	       790 ns/op	     448 B/op	       5 allocs/op
BenchmarkAero_GithubParam        	 2000000	       865 ns/op	     448 B/op	       5 allocs/op
BenchmarkAero_GithubAll          	   10000	    177469 ns/op	   90954 B/op	    1015 allocs/op
BenchmarkGoJSONREST_GithubStatic 	 1000000	      2482 ns/op	    1401 B/op	      24 allocs/op
BenchmarkGoJSONREST_GithubParam  	  500000	      3115 ns/op	    1785 B/op	      27 allocs/op
BenchmarkGoJSONREST_GithubAll    	    2000	    622440 ns/op	  352169 B/op	    5376 allocs/op
BenchmarkBuffalo_GithubStatic    	   10000	    149591 ns/op	   60987 B/op	     444 allocs/op
BenchmarkBuffalo_GithubParam     	  100000	     17106 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_GithubAll       	      50	  29901246 ns/op	12369745 B/op	   89963 allocs/op
BenchmarkGin_StaticAll           	   10000	    171971 ns/op	   82901 B/op	     942 allocs/op
BenchmarkBeego_StaticAll         	    3000	    445607 ns/op	  249962 B/op	    2983 allocs/op
BenchmarkEcho_StaticAll          	   10000	    158721 ns/op	   75363 B/op	     785 allocs/op
BenchmarkGorilaMux_StaticAll     	    2000	   1030629 ns/op	  155859 B/op	    1578 allocs/op
BenchmarkAero_StaticAll          	   10000	    127871 ns/op	   70343 B/op	     785 allocs/op
BenchmarkGoJSONRest_StaticAll    	    3000	    418410 ns/op	  219969 B/op	    3768 allocs/op
BenchmarkBuffalo_StaticAll       	      50	  22044599 ns/op	 9510993 B/op	   62822 allocs/op
PASS
ok  	github.com/dtmkeng/go-web-framework-rt	92.981s

#Air
   Air_github: 102304 Bytes
   Air_static: 102464 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/air
BenchmarkAir_Param        	  300000	      5112 ns/op	    2688 B/op	      28 allocs/op
BenchmarkAir_Param5       	  300000	      5327 ns/op	    2688 B/op	      28 allocs/op
BenchmarkAir_Param20      	  300000	      5808 ns/op	    2688 B/op	      28 allocs/op
BenchmarkAir_ParamWrite   	  300000	      5122 ns/op	    2688 B/op	      28 allocs/op
BenchmarkAir_GithubStatic 	  300000	      5043 ns/op	    2656 B/op	      27 allocs/op
BenchmarkAir_GithubParam  	  300000	      5300 ns/op	    2688 B/op	      28 allocs/op
BenchmarkAir_GithubAll    	    2000	   1084954 ns/op	  544554 B/op	    5648 allocs/op
BenchmarkAir_StaticAll    	    5000	    300140 ns/op	  113044 B/op	    2041 allocs/op
PASS
ok  	github.com/dtmkeng/air	18.354s


#Goa
   Goa_github: 86704 Bytes
   Goa_static: 86704 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/goa/http
BenchmarkGoa_Param        	 2000000	       818 ns/op	     800 B/op	       7 allocs/op
BenchmarkGoa_Param5       	 1000000	      1287 ns/op	    1024 B/op	      10 allocs/op
BenchmarkGoa_Param20      	  500000	      4043 ns/op	    3644 B/op	      14 allocs/op
BenchmarkGoa_ParamWrite   	 2000000	       927 ns/op	    1152 B/op	      12 allocs/op
BenchmarkGoa_GithubStatic 	 5000000	       327 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoa_GithubParam  	 1000000	      1081 ns/op	     832 B/op	       8 allocs/op
BenchmarkGoa_GithubAll    	   10000	    203994 ns/op	  141270 B/op	    1375 allocs/op
BenchmarkAir_StaticAll    	    5000	    357899 ns/op	  134447 B/op	    2355 allocs/op
PASS
ok  	github.com/dtmkeng/goa/http	15.748s



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