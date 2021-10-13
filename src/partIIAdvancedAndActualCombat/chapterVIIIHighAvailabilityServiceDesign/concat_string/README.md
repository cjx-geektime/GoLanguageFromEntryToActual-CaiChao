# ⾼效的字符串连接
分析
```
➜  concat_string git:(main) ✗ go test -bench=.
goos: linux
goarch: amd64
pkg: geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIIHighAvailabilityServiceDesign/concat_string
cpu: Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz
BenchmarkSprintf-20          	  105298	     13728 ns/op
BenchmarkStringBuilder-20    	 2005072	       587.2 ns/op
BenchmarkBytesBuf-20         	 1733978	       757.8 ns/op
BenchmarkStringAdd-20        	  307875	      3641 ns/op
PASS
ok  	geek-GoLanguageFromEntryToActualCombat/src/partIIAdvancedAndActualCombat/chapterVIIIHighAvailabilityServiceDesign/concat_string	6.522s
```
StringBuilder 的性能最好
