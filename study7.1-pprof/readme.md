## 性能分析
性能分析基础数据的获取有三种方式:

runtime/pprof 包 （不推荐）

net/http/pprof 包 （推荐）

go test 时添加收集参数


## 性能调优的两个方向：CPU 和内存