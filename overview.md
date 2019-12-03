# Overview

* IPFS是用于存储和访问文件，网站，应用程序和数据的分布式系统。 Global p2p merkle-dag filesystem
  * 支持弹性（有容错能力）互联网
  * 抗审查（...）
  * 就近加速访问

* IPFS生态系统为内容提供CID，并通过生成IPLD-Merkle-DAG将内容链接在一起，可以使用libp2p提供的DHT发现内容，并打开提供该内容的程序的连接，然后使用多路复用连接下载它。
  * Merkle-DAG 有向无环图
  * DHT 分布式哈希表

* 安装
  * prebuilt package https://dist.ipfs.io/#go-ipfs
  * IPFS-Desktop https://github.com/ipfs-shipyard/ipfs-desktop
  * Building from source https://docs.ipfs.io/guides/guides/install/#building-from-source

* 初始化
```
> ipfs init
initializing ipfs node at /Users/jbenet/.go-ipfs
generating 2048-bit RSA keypair...done
peer identity: Qmcpo2iLBikrdf1d6QU6vXuNb6P7hwrbNPW9kLAH8eG67z
to get started, enter:

  ipfs cat /ipfs/QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG/readme
```
```
ipfs init --profile server
```
* 配置 https://github.com/ipfs/go-ipfs/blob/v0.4.15/docs/config.md

* WEBUI http://localhost:5001/webui

* 私有网络 https://xiaozhuanlan.com/topic/2378914056

* github.com\ipfs\go-ipfs\core\node\libp2p\pnet.go swarm.key
