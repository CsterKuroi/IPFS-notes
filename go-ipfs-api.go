package main

import (
	"bytes"
	"fmt"
	
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewLocalShell()
	// get block
	blo,err:=sh.BlockGet("Qmaisz6NMhDB51cCvNWa1GMS7LU1pAxdF4Ld6Ft9kZEP2a")
	if err !=nil{
		panic(err)
	}
	fmt.Println(string(blo))

	// cat obj
	obj,err:=sh.Cat("Qmaisz6NMhDB51cCvNWa1GMS7LU1pAxdF4Ld6Ft9kZEP2a")
	if err !=nil {
		panic(err)
	}
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(obj)
	fmt.Println(buf.String())

	// id
	fmt.Println(sh.ID())
}
