# Rusty Chain OSS Security Assessment Tool

Rusty Chain is an OSS security assessment tool for assessing supply chain security issues in standalone OSS and libraries

## Usage

### CLI
```shell
$ ./rusty-chain -h

Usage of ./rusty-chain:
  -repo string
        Github repository to analyze
```

### Docker Image
```shell
$ docker build -t rusty-chain .
$ docker run rusty-chain -h

Usage of /app/rusty-chain:
  -repo string
        Github repository to analyze
```