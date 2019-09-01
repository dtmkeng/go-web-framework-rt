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
### Result after run test week-5
``` s
   Gin_github: 58480 Bytes
   Beego_github: 150920 Bytes
   Revel_github: 164448 Bytes
   Gorilamux_github: 1314160 Bytes
   Echo_github: 99960 Bytes
   Aero_github: 1403656 Bytes
   GoJSONREST_github: 142120 Bytes
   Buffalo_github: 1038888 Bytes
   Gin_static: 34064 Bytes
   Beego_staic: 98440 Bytes
   Echo_Static: 78704 Bytes
   Gorilamux_static: 583680 Bytes
   Aero_static: 784040 Bytes
   GoJSONRest_static: 136952 Bytes
   Buffalo_static: 629952 Bytes
   Revel: 111952 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/go-web-framework-rt
BenchmarkGin_Param               	20000000	        71.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param             	 2000000	       900 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param             	   30000	     55713 ns/op	   35857 B/op	     822 allocs/op
BenchmarkGorillaMux_Param        	 1000000	      2003 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_Param              	 1000000	      1859 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param              	30000000	        48.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param        	 1000000	      1123 ns/op	     649 B/op	      13 allocs/op
BenchmarkBuffalo_Param           	 1000000	      1484 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_Param5              	10000000	       132 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param5            	 1000000	      1127 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param5            	   30000	     54963 ns/op	   35859 B/op	     822 allocs/op
BenchmarkGorillaMux_Param5       	  500000	      2684 ns/op	    1344 B/op	      10 allocs/op
BenchmarkEcho_Param5             	 1000000	      1834 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param5             	50000000	        35.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param5       	 1000000	      2048 ns/op	    1097 B/op	      16 allocs/op
BenchmarkBuffalo_Param5          	 1000000	      1500 ns/op	     992 B/op	      11 allocs/op
BenchmarkGin_Param20             	 5000000	       330 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param20           	 1000000	      2120 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param20           	   30000	     54588 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param20      	  200000	      6245 ns/op	    3452 B/op	      12 allocs/op
BenchmarkEcho_Param20            	 1000000	      1820 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param20            	50000000	        35.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param20      	  200000	      7538 ns/op	    4485 B/op	      20 allocs/op
BenchmarkBuffalo_Param20         	 1000000	      1561 ns/op	    1008 B/op	      11 allocs/op
BenchmarkGin_ParamWrite          	10000000	       123 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_ParamWrite        	 2000000	       966 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_ParamWrite        	   30000	     52968 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_ParamWrite   	 1000000	      1940 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_ParamWrite         	 1000000	      1749 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_ParamWrite         	30000000	        45.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_ParamWrite   	 1000000	      1668 ns/op	    1129 B/op	      18 allocs/op
BenchmarkBuffalo_ParamWrite      	 1000000	      1387 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_GithubStatic        	20000000	        83.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubParam         	10000000	       163 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll           	   50000	     34404 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_GithubStatic      	 2000000	       876 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubParam       	 2000000	       999 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubAll         	   10000	    211400 ns/op	   71460 B/op	     609 allocs/op
BenchmarkRevel_GithubStatic      	   30000	     51945 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubParam       	   30000	     51876 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubAll         	     100	  10662776 ns/op	 7278952 B/op	  166866 allocs/op
BenchmarkGorilaMux_GithubStatic  	  200000	      9385 ns/op	     976 B/op	       9 allocs/op
BenchmarkGorilaMux_GithubParam   	  300000	      5792 ns/op	    1296 B/op	      10 allocs/op
BenchmarkGorilaMux_GithubAll     	     500	   3366160 ns/op	  252349 B/op	    1994 allocs/op
BenchmarkEcho_GithubStatic       	20000000	        81.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubParam        	10000000	       153 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubAll          	   50000	     31417 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubStatic       	30000000	        53.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubParam        	10000000	       120 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubAll          	   50000	     25091 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONREST_GithubStatic 	 2000000	       799 ns/op	     329 B/op	      11 allocs/op
BenchmarkGoJSONREST_GithubParam  	 1000000	      1404 ns/op	     713 B/op	      14 allocs/op
BenchmarkGoJSONREST_GithubAll    	    5000	    286008 ns/op	  134375 B/op	    2737 allocs/op
BenchmarkBuffalo_GithubStatic    	   20000	     90092 ns/op	   34767 B/op	     360 allocs/op
BenchmarkBuffalo_GithubParam     	  100000	     17002 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_GithubAll       	     100	  17842777 ns/op	 7049415 B/op	   72883 allocs/op
BenchmarkGin_StaticAll           	  100000	     20711 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_StaticAll         	   10000	    153871 ns/op	   55267 B/op	     471 allocs/op
BenchmarkEcho_StaticAll          	  100000	     20039 ns/op	       0 B/op	       0 allocs/op
BenchmarkGorilaMux_StaticAll     	    2000	    977675 ns/op	  153343 B/op	    1421 allocs/op
BenchmarkAero_StaticAll          	  100000	     14965 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONRest_StaticAll    	   10000	    146491 ns/op	   51655 B/op	    1727 allocs/op
BenchmarkBuffalo_StaticAll       	     100	  12891991 ns/op	 5370019 B/op	   49521 allocs/op
PASS
ok  	github.com/dtmkeng/go-web-framework-rt	117.060s

# Air
   Air_github: 101664 Bytes
   Air_static: 102144 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/air
BenchmarkAir_Param        	 2000000	       837 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_Param5       	 2000000	       909 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_Param20      	 1000000	      1374 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_ParamWrite   	  300000	      6079 ns/op	    2784 B/op	      31 allocs/op
BenchmarkAir_GithubStatic 	 2000000	       719 ns/op	     240 B/op	       6 allocs/op
BenchmarkAir_GithubParam  	 2000000	       945 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_GithubAll    	   10000	    185314 ns/op	   54066 B/op	    1385 allocs/op
BenchmarkAir_StaticAll    	    5000	    295637 ns/op	  113046 B/op	    2041 allocs/op
PASS
ok  	github.com/dtmkeng/air	17.533s

#Goa
   Goa_github: 86704 Bytes
   Goa_static: 86704 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/goa/http
BenchmarkGoa_Param        	 3000000	       524 ns/op	     784 B/op	       6 allocs/op
BenchmarkGoa_Param5       	 2000000	       997 ns/op	    1008 B/op	       9 allocs/op
BenchmarkGoa_Param20      	  500000	      3766 ns/op	    3628 B/op	      13 allocs/op
BenchmarkGoa_ParamWrite   	 2000000	       944 ns/op	    1152 B/op	      12 allocs/op
BenchmarkGoa_GithubStatic 	30000000	        48.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoa_GithubParam  	 2000000	       755 ns/op	     816 B/op	       7 allocs/op
BenchmarkGoa_GithubAll    	   10000	    138339 ns/op	  138002 B/op	    1172 allocs/op
BenchmarkAir_StaticAll    	    5000	    359019 ns/op	  134441 B/op	    2355 allocs/op
PASS
ok  	github.com/dtmkeng/goa/http	17.033s


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