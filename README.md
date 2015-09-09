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

###
```bash
# for eip-mode
Variables supported
REMAP_ELASTIC_IP=54.88.113.3
REMAP_EIP_ALLOCATION_ID=eipalloc-123abc

# for dns-mode
REMAP_HOSTED_ZONE_ID="ZSDOFPKO23-0K"
REMAP_DNS_NAME="my-route53.domain.com"
REMAP_USE_PUBLIC_IP="false"
```

## IAM Role/User credentials needed:
###eip-mode only
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Stmt1441837139000",
            "Effect": "Allow",
            "Action": [
                "ec2:AssociateAddress"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
```

### dns-mode
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Stmt1439781739000",
            "Effect": "Allow",
            "Action": [
                "route53:ChangeResourceRecordSets"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
```

### example

```bash
ubuntu@remap-test:~$ remap dns-mode --hosted-zone-id "XXXXXXXXXXXX" --dns-name "test.example.com" --run-once "true" --from-userdata "false"
2015/08/18 16:31:25 [remap started]
2015/08/18 16:31:25 DNS Mode running
2015/08/18 16:31:25 Instance IP: 10.179.1.235 A records in test.example.com [10.179.1.235]
2015/08/18 16:31:25 DNS entry points to instance IP. No changes nessesary
```

```bash
ubuntu@remap-test:~$ remap eip-mode --run-once true --from-userdata true
```


Added libraries via subtrees
```bash
git subtree add --prefix _vendor/src/github.com/aws/aws-sdk-go git@github.com:aws/aws-sdk-go.git master --squash
git subtree add --prefix _vendor/src/github.com/vaughan0/go-ini git@github.com:vaughan0/go-ini.git master --squash
```

