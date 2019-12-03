package main

import (
	"bytes"
	"fmt"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewLocalShell()
	cid, err := sh.Add(strings.NewReader("hello world! jihao gogo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(cid)
	// get block
	blo, err := sh.BlockGet("QmXnSYWiLndBSK69gboyVAsgpvJf1aE1SEfw7Mq65heiNB")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(blo))

	// cat obj
	obj, err := sh.Cat("Qmaisz6NMhDB51cCvNWa1GMS7LU1pAxdF4Ld6Ft9kZEP2a")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(obj)
	fmt.Println(buf.String())

	// id
	fmt.Println(sh.ID())
}
