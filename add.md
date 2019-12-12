# IPFS ： Add命令解析

这份文档的目的是捕捉用于添加文件的代码流，在这个过程中探索一些数据结构和包像 ipld.Node（又名dagnode）FSNode，MFS等等。

```
# Convert a file to the IPFS format.
echo "Hello World" > new-file
ipfs add new-file
added QmWATWQ7fVPP2EFGu71UkfnqhYXDYH566qy47CnJDgvs8u new-file
12 B / 12 B [=========================================================] 100.00%

# Add a file to the MFS.
NEW_FILE_HASH=$(ipfs add new-file -Q)
ipfs files cp /ipfs/$NEW_FILE_HASH /new-file

# Get information from the file in MFS.
ipfs files stat /new-file
# QmWATWQ7fVPP2EFGu71UkfnqhYXDYH566qy47CnJDgvs8u
# Size: 12
# CumulativeSize: 20
# ChildBlocks: 0
# Type: file

# Retrieve the contents.
ipfs files read /new-file
# Hello World
```

```
UnixfsAPI.Add()- 进入Unixfs包装的入口
```

```
adder.add(io.Reader) - Create and return the root DAG node
https://github.com/ipfs/go-ipfs/blob/v0.4.18/core/coreunix/add.go#L115
此方法通过使用将输入数据（io.Reader）分成多个块并将其组织到DAG中（通过点滴或平衡的布局。有关更多信息，请参见平衡），将输入数据（）转换为DAG树。Chunker
该方法返回根 ipld.Node所述的DAG。
```
