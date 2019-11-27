# IPFS-notes

* IPFS是用于存储和访问文件，网站，应用程序和数据的分布式系统。
  * 支持弹性（有容错能力）互联网
  * 抗审查（...）
  * 就近加速访问

* IPFS生态系统为内容提供CID，并通过生成IPLD-Merkle-DAG将内容链接在一起，可以使用libp2p提供的DHT发现内容，并打开提供该内容的程序的连接，然后使用多路复用连接下载它。
  * Merkle-DAG 有向无环图
  * DHT 分布式哈希表

## [references](./reference.md)
