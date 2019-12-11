## libp2p的三次握手
代码在github.com/libp2p/go-libp2p-secio/protocol.go中的runHandshakeSync方法里，发起连接的peer和接收连接的peer均执行相同的逻辑，两者唯一的区别就是前者知道对端的peerID，而后者不知道。


## 数据结构
```
message Propose {
    optional bytes rand = 1;
    optional bytes pubkey = 2;
    optional string exchanges = 3;
    optional string ciphers = 4;
    optional string hashes = 5;
}

message Exchange {
    optional bytes epubkey = 1;
    optional bytes signature = 2;
}
```
## 第一次握手： Propose
双方均发起一个Hello数据包，其格式为Hello = (rand, PublicKey, Supported)
```
    proposeOut.Rand = nonceOut
    proposeOut.Pubkey = myPubKeyBytes
    proposeOut.Exchanges = "P-256,P-384,P-521"
    proposeOut.Ciphers = "AES-256,AES-128,Blowfish"
    proposeOut.Hashes = "SHA256,SHA512"
```
在接收到对端节点的hello包后

1. 获取对方的公钥并生成对应的peerID

2. 确定双方均认可的椭圆曲线加密算法、对称加密算法和hash算法

## 第二次握手：Exchange
1. 利用选择好的椭圆曲线加密算法生成一个临时公钥，并使用自己的私钥对自己发送的hello包、接收到的hello包和公钥进行签名，并组装好对应的exchange包

2. 使用对端节点的公钥对收到的exchange包进行验证

3. 生成一个自己和远端节点的hmac密钥和加密算法

## 第三次握手：Verify
1. 组装自己的ETM ReadWriter（加密并进行信息摘要）

2. 使用加密算法发送对方第一次握手生成的随机数并校验对端的发送过来的随机数，并判断是否和自己的随机数一致

## 加密通讯
三次握手成功后，那么双方节点后续的通讯均采用加密+信息摘要方式进行通讯，具体处理逻辑在

github.com/libp2p/go-libp2p-secio/rw.go中的

WriteMsg和macCheckThenDecryp方法中
