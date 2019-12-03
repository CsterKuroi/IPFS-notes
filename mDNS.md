# mDNS

```
路由层需要支持存储内容的查找以及IPFS节点的路由查找，为了实现这个目的，可以采用DHTS，mdns，snr甚至是dns协议来，具体根据设计的需要，动态的设计和配置所采用的路由协议，比如mdns在适合在局域网中发现节点并路由数据

libP2P定义了routing 接口，目前有2个实现，分别是KAD routing 和 MDNS routing

除了KAD routing 之外， IPFS 也实现了MDNS routing, 主要用来在局域网内发现节点, 这个功能相对比较独立， 由于用到了多播地址， 在一些公有云部署环境中可能无法工作。
```

```
Discovery
目前系统支持3种发现方式， 分别是：
bootstrap 通过配置的启动节点发现其他的节点
random walk 通过查询随机生成的peerID， 从而发现新的节点 ?
mdns 通过multicast 发现局域网内的节点
```

```
有三种方式可以找节点，一个是Bootstrap，也就是在config中配置的引导节点；一个是MDNS，搞网络底层的人应该都知道这个局域网组播协议；再一个就是Kad-DHT了，这个协议早年玩过eMule和BT的人肯定很熟，eMule里有一个专门的Kad网，用的就是这个协议。

```
