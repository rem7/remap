# REMAP

Utility to assign EIP to instance or update a DNS entry in Route53 with the instance's IP.

## Compiling

```bash
$ git clone git@github.com:rem7/remap.git
$ cd remap
# the aws-go-sdk is vendored in since it hasn't reached a stable version
$ export GOPATH=$(pwd)/_vendor:$GOPATH 
# fetch other dependencies that support gopkg.in 
$ go get
$ go build
```

## Running

The application supports two modes: eip-mode and dns-mode. In eip-mode it will assign [forcefully] the EIP passed in through the parameters. In dns-mode it will update a record with the instance public-IP or private-IP. remap can also be run once or as a service. 

### example

```bash
ubuntu@remap-test:~$ remap dns-mode --hosted-zone-id "XXXXXXXXXXXX" --dns-name "test.example.com" --run-once "true" --from-userdata "false"
2015/08/18 16:31:25 [remap started]
2015/08/18 16:31:25 DNS Mode running
2015/08/18 16:31:25 Instance IP: 10.179.1.235 A records in test.example.com [10.179.1.235]
2015/08/18 16:31:25 DNS entry points to instance IP. No changes nessesary
```

Added libraries via subtrees
```bash
git subtree add --prefix _vendor/src/github.com/aws/aws-sdk-go git@github.com:aws/aws-sdk-go.git master --squash
git subtree add --prefix _vendor/src/github.com/vaughan0/go-ini git@github.com:vaughan0/go-ini.git master --squash
```

