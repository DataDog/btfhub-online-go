package main

import (
	"fmt"
	"github.com/seek-ret/btfhub-online-go/btfhubonline"
	"io/ioutil"
	"log"
)

const (
	defaultFilePermissions = 0600
)

func main() {
	client, err := btfhubonline.New("https://btfhub.seekret.io")
	if err != nil {
		log.Fatal(err)
	}

	bpfByteCode, err := ioutil.ReadFile("/home/guy/code/ebpf-sniffer/cmd/sniffer/bpf.core.o")
	if err != nil {
		log.Fatal(err)
	}

	btfIdentifier := btfhubonline.BTFRecordIdentifier{
		Distribution:        "ubuntu",
		DistributionVersion: "20.04",
		KernelVersion:       "5.11.0-1021-gcp",
		Arch:                "x86_64",
	}

	btf, err := client.GetCustomBTF(btfIdentifier, bpfByteCode)
	if err != nil {
		log.Fatal(err)
	}

	localBTFPath := fmt.Sprintf("%s_%s_%s_%s.btf", btfIdentifier.Distribution, btfIdentifier.DistributionVersion, btfIdentifier.Arch, btfIdentifier.KernelVersion)
	if err := ioutil.WriteFile(localBTFPath, btf, defaultFilePermissions); err != nil {
		log.Fatal(err)
	}
}
