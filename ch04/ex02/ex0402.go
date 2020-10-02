package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func EncodeStringToSHA(b []byte, mode string) [32]uint8 {
	switch mode {
	case "256":
		sha := sha256.Sum256(b)
	case "512":
		sha := sha512.Sum512(b)
	default:
		log.Fatal("unexpected mode")
	}
	return sha
}

func main() {
	var mode string
	flag.Parse()
	flag.StringVar(&mode, "m", "sha256", "")
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	switch mode {
	case "256":
		sha := sha256.Sum256(b)
		fmt.Printf("%x\n", sha)

	case "512":
		sha := sha512.Sum512(b)
		fmt.Printf("%x\n", sha)

	default:
		log.Fatal("unexpected mode")
	}

}
