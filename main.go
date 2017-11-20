// test project main.go
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/CrowdSurge/banner"
	"github.com/alyu/configparser"
)

func main() {
	fmt.Println("Hello World!")
	banner.Print("hello")
	//test
	config, err := configparser.Read("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	// Print the full configuration
	fmt.Println(config)
	//Encode nd Decode
	a := "73616d706c65"
	bs, err := hex.DecodeString(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	src := []byte("sample")
	encodedStr := hex.EncodeToString(src)
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.ToUpper(encodedStr))
}
