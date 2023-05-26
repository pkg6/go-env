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
	// now do something with s3 or whatever
}
