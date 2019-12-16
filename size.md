```
$ ipfs object links  /ipfs/QmdHY9TnXaXez2bSsQqanm2WaAineyiaFydwqqkEk6kyzs
QmbWqxBEKC3P8tqsKc98xmWNzrzDtRLMiMPL8wBuTGsMnR 119776 guardian.jpg

$ ipfs object stat  /ipfs/QmdHY9TnXaXez2bSsQqanm2WaAineyiaFydwqqkEk6kyzs
NumLinks: 1
BlockSize: 60
LinksSize: 58
DataSize: 2
CumulativeSize: 119836

$ ipfs object stat /ipfs/QmdHY9TnXaXez2bSsQqanm2WaAineyiaFydwqqkEk6kyzs/guardian.jpg
NumLinks: 0
BlockSize: 119776
LinksSize: 4
DataSize: 119772
CumulativeSize: 119776
119772是的大小guardian.jpg（仅包含文件数据的块）
119776是guardian.jpgDAG 的大小（文件数据+ DAG元数据）
119836是guardian.jpgDAG 的大小+ CID下包装目录的元数据QmdHY9TnXaXez2bSsQqanm2WaAineyiaFydwqqkEk6kyzs
```
