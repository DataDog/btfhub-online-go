package main

import (
	"fmt"
	"github.com/seek-ret/btfhub-online-go/btfhubonline"
	"log"
)

func main() {
	client, err := btfhubonline.New("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	btfs, err := client.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, btf := range btfs {
		fmt.Printf("Server has: \"%s %s %s %s\"\n", btf.Distribution, btf.DistributionVersion, btf.Arch, btf.KernelVersion)
	}
}
