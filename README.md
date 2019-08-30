### Week-2 HTTP routes test
``` s
   Gin_github: 58480 Bytes
   Beego_github: 150920 Bytes
   Revel_github: 164448 Bytes
   Gorilamux_github: 1314160 Bytes
   Echo_github: 99960 Bytes
   Aero_github: 1404248 Bytes
   GoJSONREST_github: 142088 Bytes
   Buffalo_github: 1038888 Bytes
   Gin_static: 34064 Bytes
   Beego_staic: 98440 Bytes
   Echo_Static: 78704 Bytes
   Gorilamux_static: 583680 Bytes
   Aero_static: 784040 Bytes
   GoJSONRest_static: 136488 Bytes
   Buffalo_static: 629952 Bytes
   Revel: 111328 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/go-web-framework-rt
BenchmarkGin_Param               	20000000	        68.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param             	 2000000	       846 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param             	   30000	     55382 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param        	 1000000	      1991 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_Param              	 1000000	      1892 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param              	30000000	        46.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param        	 1000000	      1132 ns/op	     649 B/op	      13 allocs/op
BenchmarkBuffalo_Param           	 1000000	      1453 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_Param5              	10000000	       127 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param5            	 1000000	      1111 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param5            	   30000	     55784 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param5       	  500000	      2743 ns/op	    1344 B/op	      10 allocs/op
BenchmarkEcho_Param5             	 1000000	      1861 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param5             	50000000	        33.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param5       	 1000000	      2095 ns/op	    1097 B/op	      16 allocs/op
BenchmarkBuffalo_Param5          	 1000000	      1494 ns/op	     992 B/op	      11 allocs/op
BenchmarkGin_Param20             	 5000000	       329 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_Param20           	 1000000	      2172 ns/op	     352 B/op	       3 allocs/op
BenchmarkRevel_Param20           	   30000	     54960 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_Param20      	  200000	      6332 ns/op	    3452 B/op	      12 allocs/op
BenchmarkEcho_Param20            	 1000000	      1851 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_Param20            	50000000	        33.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_Param20      	  200000	      7824 ns/op	    4484 B/op	      20 allocs/op
BenchmarkBuffalo_Param20         	 1000000	      1553 ns/op	    1008 B/op	      11 allocs/op
BenchmarkGin_ParamWrite          	20000000	       115 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_ParamWrite        	 2000000	       881 ns/op	     360 B/op	       4 allocs/op
BenchmarkRevel_ParamWrite        	   30000	     56010 ns/op	   35856 B/op	     822 allocs/op
BenchmarkGorillaMux_ParamWrite   	 1000000	      2023 ns/op	    1280 B/op	      10 allocs/op
BenchmarkEcho_ParamWrite         	 1000000	      1846 ns/op	    1040 B/op	      12 allocs/op
BenchmarkAero_ParamWrite         	30000000	        44.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGOJSONREST_ParamWrite   	 1000000	      1706 ns/op	    1129 B/op	      18 allocs/op
BenchmarkBuffalo_ParamWrite      	 1000000	      1460 ns/op	     976 B/op	      11 allocs/op
BenchmarkGin_GithubStatic        	20000000	        82.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubParam         	10000000	       156 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll           	   50000	     33381 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_GithubStatic      	 2000000	       846 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubParam       	 2000000	       994 ns/op	     352 B/op	       3 allocs/op
BenchmarkBeego_GithubAll         	   10000	    218792 ns/op	   71457 B/op	     609 allocs/op
BenchmarkRevel_GithubStatic      	   30000	     54938 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubParam       	   30000	     55219 ns/op	   35856 B/op	     822 allocs/op
BenchmarkRevel_GithubAll         	     100	  11227636 ns/op	 7278768 B/op	  166866 allocs/op
BenchmarkGorilaMux_GithubStatic  	  200000	      9670 ns/op	     976 B/op	       9 allocs/op
BenchmarkGorilaMux_GithubParam   	  300000	      5954 ns/op	    1296 B/op	      10 allocs/op
BenchmarkGorilaMux_GithubAll     	     500	   3462369 ns/op	  251660 B/op	    1994 allocs/op
BenchmarkEcho_GithubStatic       	20000000	        84.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubParam        	10000000	       159 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubAll          	   50000	     31601 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubStatic       	30000000	        52.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubParam        	10000000	       120 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubAll          	   50000	     25100 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONREST_GithubStatic 	 2000000	       786 ns/op	     329 B/op	      11 allocs/op
BenchmarkGoJSONREST_GithubParam  	 1000000	      1489 ns/op	     713 B/op	      14 allocs/op
BenchmarkGoJSONREST_GithubAll    	    5000	    300410 ns/op	  134371 B/op	    2737 allocs/op
BenchmarkBuffalo_GithubStatic    	   20000	     96404 ns/op	   34763 B/op	     360 allocs/op
BenchmarkBuffalo_GithubParam     	  100000	     17409 ns/op	    1008 B/op	      11 allocs/op
BenchmarkBuffalo_GithubAll       	     100	  19116636 ns/op	 7049600 B/op	   72885 allocs/op
BenchmarkGin_StaticAll           	  100000	     20382 ns/op	       0 B/op	       0 allocs/op
BenchmarkBeego_StaticAll         	   10000	    159033 ns/op	   55264 B/op	     471 allocs/op
BenchmarkEcho_StaticAll          	  100000	     19865 ns/op	       0 B/op	       0 allocs/op
BenchmarkGorilaMux_StaticAll     	    2000	   1007488 ns/op	  153332 B/op	    1421 allocs/op
BenchmarkAero_StaticAll          	  100000	     14406 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoJSONRest_StaticAll    	   10000	    152124 ns/op	   51653 B/op	    1727 allocs/op
BenchmarkBuffalo_StaticAll       	     100	  13618804 ns/op	 5369174 B/op	   49519 allocs/op
PASS
ok  	github.com/dtmkeng/go-web-framework-rt	119.506s

# Air

   Air_github: 101824 Bytes
   Air_static: 101504 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/air
BenchmarkAir_Param        	 2000000	       837 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_Param5       	 2000000	       951 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_Param20      	 1000000	      1377 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_ParamWrite   	  300000	      6178 ns/op	    2784 B/op	      31 allocs/op
BenchmarkAir_GithubStatic 	 2000000	       762 ns/op	     240 B/op	       6 allocs/op
BenchmarkAir_GithubParam  	 2000000	       958 ns/op	     272 B/op	       7 allocs/op
BenchmarkAir_GithubAll    	   10000	    185993 ns/op	   54064 B/op	    1385 allocs/op
BenchmarkAir_StaticAll    	    5000	    299588 ns/op	  113041 B/op	    2041 allocs/op
PASS
ok  	github.com/dtmkeng/air	22.024s

# Goa
   Goa_github: 86688 Bytes
   Goa_static: 86704 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/goa/http
BenchmarkGoa_Param        	 3000000	       432 ns/op	     784 B/op	       6 allocs/op
BenchmarkGoa_Param5       	 2000000	       829 ns/op	    1008 B/op	       9 allocs/op
BenchmarkGoa_Param20      	  500000	      3066 ns/op	    3628 B/op	      13 allocs/op
BenchmarkGoa_ParamWrite   	 2000000	       753 ns/op	    1152 B/op	      12 allocs/op
BenchmarkGoa_GithubStatic 	30000000	        46.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoa_GithubParam  	 2000000	       610 ns/op	     816 B/op	       7 allocs/op
BenchmarkGoa_GithubAll    	   10000	    110608 ns/op	  138000 B/op	    1172 allocs/op
BenchmarkAir_StaticAll    	    5000	    323994 ns/op	  134422 B/op	    2355 allocs/op
PASS
ok  	github.com/dtmkeng/goa/http	15.155s

```