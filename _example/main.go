package main

import (
	"fmt"
	"github.com/pkg6/goenv"
	"os"
)

type Config struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
}

func main() {
	goenv.Load()
	AccessKeyId := os.Getenv("AccessKeyId")
	AccessKeySecret := os.Getenv("AccessKeySecret")
	fmt.Println(AccessKeyId, AccessKeySecret)
	c := new(Config)
	goenv.JsonUnmarshal(c)
	fmt.Println(c)
}
