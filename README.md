# Golang kata for a simple word count application

## Intro

This kata deals with use of standard tooling and project structure
The intent is less on the algorithm, but to show various stages of the
pipeline process. From formatting, linting, building a docker image,
to running unit tests to verify the library is working ok.

Also dealing with some proper input strings naughtiness.

The code was written following: https://github.com/uber-go/guide/blob/master/style.md.
Note some of the test cases were taken from: https://github.com/minimaxir/big-list-of-naughty-strings/blob/master/blns.txt.

## Prerequisites

This app assumes you are running it under a LINUX-like system.
We are assuming basic golang infra, makefile and docker were already
installed.

First thing to do is to modify your ~/.bashprofile to include:

```bash
export GOPATH=$HOME/go
export GOBIN=$HOME/go/bin
export PATH=$PATH:$HOME/go/bin
```

This is to ensure you have the right go env variables set.

Please execute after editing:
```bash
source ~/.bash_profile
```

Afterwards please ensure you have installed:
- golang linter
```bash
go get -u golang.org/x/lint/golint
```
- golang test library (testify)
```bash
go get github.com/stretchr/testify
```

## FAQ

##### How to run your application?

By command line:
```bash
make all
```

You can review the specific commands in [Makefile](Makefile) file.

## License

See [LICENSE](LICENSE) file.
