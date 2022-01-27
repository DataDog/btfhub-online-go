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

## Download Custom BTF
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

## Checking distribution and kernel version

###Distribution:
```bash
cat /etc/os-release | grep "^ID="
```

Output:
```bash
ID=ubuntu
```

###Distribution Version:
```bash
cat /etc/os-release | grep "VERSION_ID="
```

Output:
```bash
VERSION_ID="20.04"
```

###Kernel Version:

Run:
```bash
uname -r
```

Output:
```bash
5.11.0-44-generic
```

###Arch:

Run:
```bash
uname -p
```

Output:
```bash
x86_64
```

# Contributing

We are welcome you to contribute in any manner there is.
We care a lot from the community, thus please read our [code of conduct](./CODE_OF_CONDUCT.md) before contributing.
You don't have to develop to contribute to the community, and you can do it in one of the following ways:

1. Open issues (bugs, feature requests)
2. Develop the code (read [contributing.md](./CONTRIBUTING.md))
3. Mention us (twitter, linkedin, blogs)

# Linters
In our project we use `golangci-linter`, and we are using the following linters:
- goimports
- gci
- nakedret
- golint
- misspell
- gosimple
- godot
- whitespace
- staticcheck
- bodyclose
- ineffassign
- unused
- varcheck
- unparam
- godox
- gosec

# License
This SDK is distributed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
