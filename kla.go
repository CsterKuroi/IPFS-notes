package main

import (
	"fmt"

	"github.com/klauspost/reedsolomon"
)

func main() {
	// To create an encoder with 10 data shards (where your data goes) and 3 parity shards (calculated)
	enc, err := reedsolomon.New(10, 3)
	if err != nil {
		panic(err)
	}

	// Data
	data := make([][]byte, 13)
	// Create all shards, size them at 50000 each
	for i := range data {
		data[i] = make([]byte, 256*1024)
	}

	// Fill some data into the data shards
	for i, in := range data[:10] {
		for j := range in {
			in[j] = byte((i + j) & 0xff)
		}
	}

	// Encode
	err = enc.Encode(data)
	if err != nil {
		panic(err)
	}

	//v1
	ok, err := enc.Verify(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)

	// Delete two data shards
	data[3] = nil
	data[7] = nil
	//data[8] = nil
	data[11] = nil

	//v2
	ok, err = enc.Verify(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok)

	// Reconstruct the missing shards
	err = enc.Reconstruct(data)
	if err != nil {
		panic(err)
	}

	//v3
	ok, err = enc.Verify(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
}
