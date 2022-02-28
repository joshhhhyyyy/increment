# increment
![Yep](https://socialify.git.ci/joshhhhyyyy/Increment/image?font=Source%20Code%20Pro&language=1&name=1&owner=1&pattern=Overlapping%20Hexagons&theme=Dark)

## What is this?
**Needlessly complex program to parse, increment, and push the latest version number tag on github.**

Err handling by **[Sentry](sentry.io)** 

Made with 😫 , 😓 &amp; 😭

## Installation
### Via Go (ALL Platforms)
```go get github.com/joshhhhyyyy/increment```

```go install github.com/joshhhhyyyy/increment```

```export PATH=$PATH:$(go env GOPATH)/bin``` (Add gopath to path)

### Via apt (Debian derivatives only)
```echo "deb [trusted=yes] https://apt.joseos.com/ ./" | sudo tee /etc/apt/sources.list.d/joseos.list```

```sudo apt update```

```sudo apt install increment```

## Usage
```increment [OPTIONS]```

## Options
**note: both single minus "-" and double minus "--" work fine

```-key=""``` // Required but should be set with environment variables instead, Type: string, Sentry dsn/key

```-nfpm``` // Not needed at all, Type: bool, Whether to build a deb with nfpm or not

```--dont-push-main=""``` // Optional, Type: bool, Optional to not push to main (make a new tag only)
