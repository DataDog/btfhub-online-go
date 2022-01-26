# btfhub-online-go
Golang client to communicate with the BTFHub online server.

# Download from GitHub

```bash
go get github.com/seek-ret/btfhubonline
```

# Examples

All examples available at the [examples directory](./examples).

## Client initialization

```go
client, err := btfhubonline.New("<server address>", btfhubonline.ClientOptions{Secure: true})
```

## List available BTFs

```go
btfs, err := client.List()
if err != nil {
    log.Fatal(err)
}

for _, btf := range btfs {
    fmt.Printf("Server has: \"%s %s %s %s\"\n", btf.Distribution, btf.DistributionVersion, btf.Arch, btf.KernelVersion)
}
```

## Download BTF

```go
btfIdentifier := btfhubonline.BTFRecordIdentifier{
    Distribution:        "ubuntu",
    DistributionVersion: "20.04",
    KernelVersion:       "5.11.0-1021-gcp",
    Arch:                "x86_64",
}

btf, err := client.GetRawBTF(btfIdentifier)
```


## Download Customized BTF
```go
bpfByteCode, err := ioutil.ReadFile("<bpf.core.o>")
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
```

# License
This SDK is distributed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
