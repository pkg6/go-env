# GoEnv

A Go (golang) port of the Ruby [go-env](https://github.com/pkg6/go-env) project

From the original Library:

It can be used as a library (for loading in env for your own daemons etc.) or as a bin command.

There is test coverage and CI for both linuxish and Windows environments, but I make no guarantees about the bin version working on Windows.

## Installation

As a library

```shell
go get github.com/pkg6/goenv
```

## Usage

Add your application configuration to your `.env` and `.json` file in the root of your project:

```shell
AccessKeyId=yourAccessKeyId
AccessKeySecret=yourAccessKeySecret
```

Then in your Go app you can do something like

```go
package main

import (
	"fmt"
	"github.com/pkg6/goenv"
	"os"
)

func main() {
	goenv.Load()
	AccessKeyId := os.Getenv("AccessKeyId")
	AccessKeySecret := os.Getenv("AccessKeySecret")
	fmt.Println(AccessKeyId, AccessKeySecret)
	// now do something with aliyun or whatever
}
```

## Contributing

*code changes without tests and references to peer dotenv implementations will not be accepted*

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

