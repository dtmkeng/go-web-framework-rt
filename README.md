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
### Result after run test week-6
``` s
      Gin_github: 58480 Bytes
   Beego_github: 150920 Bytes
   Revel_github: 164448 Bytes
   Gorilamux_github: 1314160 Bytes
   Echo_github: 99960 Bytes
   Aero_github: 1418936 Bytes
   GoJSONREST_github: 141656 Bytes
   Buffalo_github: 1038904 Bytes
   Gin_static: 34064 Bytes
   Beego_staic: 98440 Bytes
   Echo_Static: 78704 Bytes
   Gorilamux_static: 583680 Bytes
   Aero_static: 784040 Bytes
   GoJSONRest_static: 136520 Bytes
   Buffalo_static: 629952 Bytes
   Revel: 111120 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/go-web-framework-rt
BenchmarkGin_Param               	 2000000	       728 ns/op	     448 B/op	       3 allocs/op
BenchmarkBeego_Param             	 2000000	       999 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_Param             	   30000	     55746 ns/op	   35857 B/op	     822 allocs/op
BenchmarkGorillaMux_Param        	 1000000	      1995 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_Param              	 1000000	      1792 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param              	30000000	        45.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param        	  500000	      2871 ns/op	    1721 B/op	      26 allocs/op
BenchmarkBuffalo_Param           	 1000000	      1412 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_Param5              	 2000000	       733 ns/op	     448 B/op	       3 allocs/op
BenchmarkBeego_Param5            	 1000000	      1123 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_Param5            	   30000	     52946 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param5       	  500000	      2642 ns/op	    1344 B/op	      10 allocs/op
BenchmarkEcho_Param5             	 1000000	      1780 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param5             	50000000	        34.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param5       	  500000	      3771 ns/op	    2169 B/op	      29 allocs/op
BenchmarkBuffalo_Param5          	 1000000	      1450 ns/op	     992 B/op	      11 allocs/op
BenchmarkGin_Param20             	 2000000	       893 ns/op	     448 B/op	       3 allocs/op
BenchmarkBeego_Param20           	 1000000	      2089 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_Param20           	   30000	     52654 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param20      	  200000	      6085 ns/op	    3452 B/op	      12 allocs/op
BenchmarkEcho_Param20            	 1000000	      1761 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param20            	50000000	        34.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param20      	  200000	      9228 ns/op	    5558 B/op	      33 allocs/op
BenchmarkBuffalo_Param20         	 1000000	      1504 ns/op	    1008 B/op	      11 allocs/op
BenchmarkGin_GithubStatic        	 2000000	       635 ns/op	     448 B/op	       3 allocs/op
BenchmarkGin_GithubParam         	 2000000	       728 ns/op	     448 B/op	       3 allocs/op
BenchmarkGin_GithubAll           	   10000	    153410 ns/op	   90947 B/op	     609 allocs/op
BenchmarkBeego_GithubStatic      	 2000000	       933 ns/op	     360 B/op	       4 allocs/op
BenchmarkBeego_GithubParam       	 1000000	      1049 ns/op	     360 B/op	       4 allocs/op
BenchmarkBeego_GithubAll         	   10000	    219078 ns/op	   73084 B/op	     812 allocs/op
BenchmarkRevel_GithubStatic      	   30000	     51799 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubParam       	   30000	     51757 ns/op	   35857 B/op	     822 allocs/op
BenchmarkRevel_GithubAll         	     100	  10485501 ns/op	 7278849 B/op	  166866 allocs/op
BenchmarkGorilaMux_GithubStatic  	  200000	      9335 ns/op	     976 B/op	       9 allocs/op
BenchmarkGorilaMux_GithubParam   	  300000	      6096 ns/op	    1296 B/op	      10 allocs/op
BenchmarkGorilaMux_GithubAll     	     500	   3392809 ns/op	  251715 B/op	    1994 allocs/op
BenchmarkEcho_GithubStatic       	 2000000	       623 ns/op	     424 B/op	       4 allocs/op
BenchmarkEcho_GithubParam        	 2000000	       697 ns/op	     424 B/op	       4 allocs/op
BenchmarkEcho_GithubAll          	   10000	    144642 ns/op	   86075 B/op	     812 allocs/op
BenchmarkAero_GithubStatic       	20000000	        63.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubParam        	10000000	       130 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubAll          	   50000	     27581 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONREST_GithubStatic 	  500000	      2493 ns/op	    1401 B/op	      24 allocs/op
BenchmarkGoJSONREST_GithubParam  	  500000	      3102 ns/op	    1785 B/op	      27 allocs/op
BenchmarkGoJSONREST_GithubAll    	    2000	    628031 ns/op	  352008 B/op	    5376 allocs/op
BenchmarkBuffalo_GithubStatic    	    1000	   1921119 ns/op	 1175074 B/op	   12768 allocs/op
BenchmarkBuffalo_GithubParam     	  100000	     17307 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_GithubAll       	       3	 390258046 ns/op	238578090 B/op	 2591831 allocs/op
BenchmarkGin_StaticAll           	   10000	    106738 ns/op	   70338 B/op	     471 allocs/op
BenchmarkBeego_StaticAll         	   10000	    161885 ns/op	   56523 B/op	     628 allocs/op
BenchmarkEcho_StaticAll          	   10000	    105107 ns/op	   66570 B/op	     628 allocs/op
BenchmarkGorilaMux_StaticAll     	    2000	   1014661 ns/op	  153348 B/op	    1421 allocs/op
BenchmarkAero_StaticAll          	  100000	     16837 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONRest_StaticAll    	    5000	    422567 ns/op	  219967 B/op	    3768 allocs/op
BenchmarkBuffalo_StaticAll       	       5	 237889922 ns/op	131495598 B/op	 1617641 allocs/op
PASS
ok  	github.com/dtmkeng/go-web-framework-rt	100.621s

#Air
   Air_github: 102304 Bytes
   Air_static: 102144 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/air
BenchmarkAir_Param        	  300000	      4710 ns/op	    2629 B/op	      26 allocs/op
BenchmarkAir_Param5       	  300000	      4811 ns/op	    2629 B/op	      26 allocs/op
BenchmarkAir_Param20      	  300000	      5400 ns/op	    2629 B/op	      26 allocs/op
BenchmarkAir_ParamWrite   	  300000	      6008 ns/op	    2784 B/op	      31 allocs/op
BenchmarkAir_GithubStatic 	  300000	      4603 ns/op	    2597 B/op	      25 allocs/op
BenchmarkAir_GithubParam  	  300000	      4863 ns/op	    2629 B/op	      26 allocs/op
BenchmarkAir_GithubAll    	    2000	    990228 ns/op	  532637 B/op	    5242 allocs/op
BenchmarkAir_StaticAll    	    5000	    299019 ns/op	  113046 B/op	    2041 allocs/op
PASS
ok  	github.com/dtmkeng/air	17.720s

#Goa
   Goa_github: 86704 Bytes
   Goa_static: 86704 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/goa/http
BenchmarkGoa_Param        	 3000000	       588 ns/op	     784 B/op	       6 allocs/op
BenchmarkGoa_Param5       	 1000000	      1060 ns/op	    1008 B/op	       9 allocs/op
BenchmarkGoa_Param20      	  500000	      3797 ns/op	    3629 B/op	      13 allocs/op
BenchmarkGoa_ParamWrite   	 2000000	       924 ns/op	    1152 B/op	      12 allocs/op
BenchmarkGoa_GithubStatic 	20000000	        96.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoa_GithubParam  	 2000000	       830 ns/op	     816 B/op	       7 allocs/op
BenchmarkGoa_GithubAll    	   10000	    155437 ns/op	  138023 B/op	    1172 allocs/op
BenchmarkAir_StaticAll    	    5000	    359089 ns/op	  134441 B/op	    2355 allocs/op
PASS
ok  	github.com/dtmkeng/goa/http	16.238s


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