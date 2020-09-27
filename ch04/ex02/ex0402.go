package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func EncodeStringToSHA(b []byte, mode string) ([32]uint8, error) {
	sha := sha256.Sum256(b)
	return sha, nil
}

func main() {
	var mode string
	flag.Parse()
	flag.StringVar(&mode, "m", "sha256", "")
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	sha, err := EncodeStringToSHA(b, mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", sha)

}
