# grasp-go-lib
掌握go语言的官网库

[函数]
return 的理解

[功能]
理解这些坑的关键是这条语句：

return xxx
上面这条语句经过编译之后，变成了三条指令：

1. 返回值 = xxx
2. 调用defer函数
3. 空的return

1,3步才是Return 语句真正的命令，第2步是defer定义的语句，这里可能会操作返回值。

[扩展]

