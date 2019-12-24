```
aliyun组播禁止了，需要bootstrap/swarm connect

go get github.com/Kubuxu/go-ipfs-swarm-key-gen/ipfs-swarm-key-gen
ipfs-swarm-key-gen > ~/.ipfs/swarm.key
 
 
ipfs bootstrap list
ipfs bootstrap rm --all
ipfs bootstrap add  /ip4/xxxxxxxxxx/tcp/4001/ipfs/QmabYFoBnCwUYHCx5wU75wtQ57YxqNdHRbLwpX8rVA2Bzo


ipfs swarm peers
ipfs swarm connect /ip4/xxxxxxxxxx/tcp/4001/ipfs/QmabYFoBnCwUYHCx5wU75wtQ57YxqNdHRbLwpX8rVA2Bzo

ipfs daemon
```

```
ipfs files cp /ipfs/QmeDqoRmdHa4DAy4N9KZezuc7XJXp1oeaGcPBeEtfFPDht /nohup.out
ipfs files read /nohup.out
```

显示所有数据块hash
```
ipfs refs local|wc -l
```
