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
	client, err := btfhubonline.New("localhost:8080", btfhubonline.ClientOptions{Secure: false})
	if err != nil {
		log.Fatal(err)
	}

	btfIdentifier := btfhubonline.BTFRecordIdentifier{
		Distribution:        "ubuntu",
		DistributionVersion: "20.04",
		KernelVersion:       "5.11.0-1021-gcp",
		Arch:                "x86_64",
	}

	btf, err := client.GetRawBTF(btfIdentifier)
	if err != nil {
		log.Fatal(err)
	}

	localBTFPath := fmt.Sprintf("%s_%s_%s_%s.tar.xz", btfIdentifier.Distribution, btfIdentifier.DistributionVersion, btfIdentifier.Arch, btfIdentifier.KernelVersion)
	if err := ioutil.WriteFile(localBTFPath, btf, defaultFilePermissions); err != nil {
		log.Fatal(err)
	}
}
