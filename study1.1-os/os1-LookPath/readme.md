[函数]
func exec.LookPath()

[功能]
查找环境变量中指定的文件名路径

[扩展]
还能用在什么地方

[疑惑]
函数： splitList
list := []string{}
start := 0
quo := false
for i := 0; i < len(path); i++ {
switch c := path[i]; {
case c == '"':
quo = !quo
case c == ListSeparator && !quo:
list = append(list, path[start:i])
start = i + 1
}
}
list = append(list, path[start:])

为什么不使用 strings.Split 进行分割呢
可能因为手动分割效率更高