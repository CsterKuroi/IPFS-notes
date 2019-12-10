# 分析 ipfs daemon

```
Initializing daemon...
go-ipfs version: 0.4.22-
Repo version: 7
System version: amd64/windows
Golang version: go1.12.7
Swarm listening on /ip4/10.0.130.239/tcp/4001
Swarm listening on /ip4/127.0.0.1/tcp/4001
Swarm listening on /ip4/169.254.148.101/tcp/4001
Swarm listening on /ip4/169.254.167.68/tcp/4001
Swarm listening on /ip4/169.254.212.74/tcp/4001
Swarm listening on /ip4/169.254.253.152/tcp/4001
Swarm listening on /ip6/::1/tcp/4001
Swarm listening on /p2p-circuit
Swarm announcing /ip4/10.0.130.239/tcp/4001
Swarm announcing /ip4/127.0.0.1/tcp/4001
Swarm announcing /ip4/169.254.148.101/tcp/4001
Swarm announcing /ip4/169.254.167.68/tcp/4001
Swarm announcing /ip4/169.254.212.74/tcp/4001
Swarm announcing /ip4/169.254.253.152/tcp/4001
Swarm announcing /ip6/::1/tcp/4001
API server listening on /ip4/127.0.0.1/tcp/5001
WebUI: http://127.0.0.1:5001/webui
Gateway (readonly) server listening on /ip4/127.0.0.1/tcp/8080
Daemon is ready
```

1. 打印版本
```
func printVersion() {
	v := version.CurrentVersionNumber
	if version.CurrentCommit != "" {
		v += "-" + version.CurrentCommit
	}
	fmt.Printf("go-ipfs version: %s\n", v)
	fmt.Printf("Repo version: %d\n", fsrepo.RepoVersion)
	fmt.Printf("System version: %s\n", runtime.GOARCH+"/"+runtime.GOOS)
	fmt.Printf("Golang version: %s\n", runtime.Version())
}
```

2. 检查fdlimit
3. 检查unencrypted
4. 检查初始化
5. 检查repo lock
6. 检查routing Option
7. 检查private network
8. 打印addresses of the host
9. 启动插件
10. 启动serveHTTPApi (corehttp)
11. 检查mount
12. 检查GC
13. 启动http gateway (corehttp)
