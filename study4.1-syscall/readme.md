exec.LookPath() -> os.Getenv(`PATHEXT`) -> syscall.Getenv(key) -> n, err = GetEnvironmentVariable(keyp, &b[0], uint32(len(b))) -> Syscall

系统调用的学习，持续添加

网络连接中，系统调用的使用