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
   Revel_github: 164448 Bytes
   Gorilamux_github: 1314160 Bytes
   Echo_github: 95992 Bytes
   Aero_github: 1403736 Bytes
   GoJSONREST_github: 142088 Bytes
   Buffalo_github: 1038888 Bytes
   Gin_static: 34064 Bytes
   Beego_staic: 98440 Bytes
   Echo_Static: 75072 Bytes
   Gorilamux_static: 583680 Bytes
   Aero_static: 784040 Bytes
   GoJSONRest_static: 136728 Bytes
   Buffalo_static: 629952 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/project/http_router
BenchmarkGin_Param               	20000000	       119 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param5              	10000000	       215 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param20             	 3000000	       546 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParamWrite          	10000000	       202 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param             	 1000000	      1520 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_Param5            	 1000000	      1725 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_Param20           	  500000	      3525 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_ParamWrite        	 1000000	      1593 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_Param             	   20000	     84991 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_Param5            	   20000	     86201 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_Param20           	   20000	     85359 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_ParamWrite        	   20000	     85984 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param        	  500000	      3190 ns/op	    1280 B/op	      10 allocs/op
BenchmarkGorillaMux_Param5       	  300000	      4497 ns/op	    1344 B/op	      10 allocs/op
BenchmarkGorillaMux_Param20      	  200000	     10183 ns/op	    3451 B/op	      12 allocs/op
BenchmarkGorillaMux_ParamWrite   	  500000	      3348 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_Param              	  500000	      2964 ns/op	    1040 B/op	      12 allocs/op
BenchmarkEcho_Param5             	  500000	      2998 ns/op	    1040 B/op	      12 allocs/op
BenchmarkEcho_Param20            	  500000	      2954 ns/op	    1040 B/op	      12 allocs/op
BenchmarkEcho_ParamWrite         	  500000	      2970 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param              	20000000	        79.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param5             	20000000	        62.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param20            	20000000	        62.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParamWrite         	20000000	        79.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param        	 1000000	      1756 ns/op	     649 B/op	      13 allocs/op
BenchmarkGOJSONREST_Param5       	  500000	      3249 ns/op	    1097 B/op	      16 allocs/op
BenchmarkGOJSONREST_Param20      	  100000	     12181 ns/op	    4484 B/op	      20 allocs/op
BenchmarkGOJSONREST_ParamWrite   	  500000	      2787 ns/op	    1129 B/op	      18 allocs/op
BenchmarkBuffalo_Param           	 1000000	      2368 ns/op	     976 B/op	      11 allocs/op
BenchmarkBuffalo_Param5          	 1000000	      2467 ns/op	     992 B/op	      11 allocs/op
BenchmarkBuffalo_Param20         	  500000	      2574 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_ParamWrite      	 1000000	      2364 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_GithubStatic        	10000000	       138 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubParam         	 5000000	       258 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll           	   20000	     63586 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_GithubStatic      	 1000000	      1543 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubParam       	 1000000	      1683 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubAll         	    5000	    349840 ns/op	   71457 B/op	     609 allocs/op
BenchmarkRevel_GithubStatic      	   20000	     85292 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubParam       	   20000	     85856 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubAll         	     100	  17782407 ns/op	 7278768 B/op	  166866 allocs/op
BenchmarkGorilaMux_GithubStatic  	  100000	     16887 ns/op	     976 B/op	       9 allocs/op
BenchmarkGorilaMux_GithubParam   	  200000	     10185 ns/op	    1296 B/op	      10 allocs/op
BenchmarkGorilaMux_GithubAll     	     200	   5932309 ns/op	  251663 B/op	    1994 allocs/op
BenchmarkEcho_GithubStatic       	10000000	       150 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubParam        	 5000000	       276 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubAll          	   30000	     55859 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubStatic       	20000000	        93.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubParam        	10000000	       207 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubAll          	   30000	     43336 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONREST_GithubStatic 	 1000000	      1384 ns/op	     329 B/op	      11 allocs/op
BenchmarkGoJSONREST_GithubParam  	 1000000	      2473 ns/op	     713 B/op	      14 allocs/op
BenchmarkGoJSONREST_GithubAll    	    3000	    500322 ns/op	  134371 B/op	    2737 allocs/op
BenchmarkBuffalo_GithubStatic    	   10000	    155720 ns/op	   34781 B/op	     360 allocs/op
BenchmarkBuffalo_GithubParam     	   50000	     30747 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_GithubAll       	     100	  30116704 ns/op	 7048496 B/op	   72885 allocs/op
BenchmarkGin_StaticAll           	   50000	     38931 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_StaticAll         	   10000	    273595 ns/op	   55265 B/op	     471 allocs/op
BenchmarkEcho_StaticAll          	   50000	     37584 ns/op	       0 B/op	       0 allocs/op
BenchmarkGorilaMux_StaticAll     	    1000	   1793488 ns/op	  153332 B/op	    1421 allocs/op
BenchmarkAero_StaticAll          	   50000	     25993 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONRest_StaticAll    	   10000	    248556 ns/op	   51653 B/op	    1727 allocs/op
BenchmarkBuffalo_StaticAll       	     100	  21897045 ns/op	 5369652 B/op	   49523 allocs/op
PASS
ok  	github.com/dtmkeng/project/http_router	124.128s
``