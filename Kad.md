# Kad

```
KadDHT的节点寻址与内容寻址同构，在KAD基于异或来计算逻辑距离的基础上，节点倾向于连接与自己距离更近的节点，存储与自己距离更近的内容Key值，并在此之上加入了延时等具体物理距离的考量(Coral DHT)。引入S-KadDHT加入了一定的抗女巫攻击能力，因为生成一个公私钥需要计算难度值
```

```
KAD距离计算 ID number= A + 256*B + 256*256*C + 256*256*256*D
```

```
x ⊕ x = 0，节点与它本身的异或距离为0。
x ⊕ y > 0 , if x != y，不同节点异或距离一定大于0。
x ⊕ y = y ⊕ x，异或距离是对称的。
x ⊕ y + y ⊕ z >= x ⊕ z，类似于三角形不等式，第三边的距离小于等于另外两边的距离之和。
x ⊕ y ⊕ y ⊕ z = x ⊕ z
x + y >= x ⊕ y

```

```
IPFS DHT的数据存储是根据数据的大小进行的：??? 白皮书宣称，实际没看到
小于1KB的数据直接存储到DHT上面
大于1KB的数据在DHT中存储的是节点ID
```

```
3.3 Routing
IPFS nodes require a routing system that can find (a)
other peers’ network addresses and (b) peers who can serve
particular objects. IPFS achieves this using a DSHT based
on S/Kademlia and Coral, using the properties discussed in
2.1. The size of objects and use patterns of IPFS are similar
to Coral [5] and Mainline [16], so the IPFS DHT makes a
distinction for values stored based on their size. Small values
(equal to or less than 1KB) are stored directly on the DHT.
For values larger, the DHT stores references, which are the
NodeIds of peers who can serve the block.
```
