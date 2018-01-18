package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// sigCheck accepts two byte arrays and returns a bool it will compare input
// and sig byte by byte to see if we have a match
func sigCheck(input []byte, sig []byte) bool {
	for i := range input {
		if input[i] != sig[i] {
			return false
		}
	}
	return true
}

// doPatch accepts two byte arrays and an int
func doPatch(output []byte, patch []byte, found int) {
	c := 0
	for i := found; i < found+len(patch); i++ {
		output[i] = patch[c]
		c++
	}
	log.Printf("Wrote %v bytes!", c)
}

func main() {

	// Prep our CLI flags
	inPtr := flag.String("in", "", "input filename")
	outPtr := flag.String("out", "", "output filename")
	sigPtr := flag.String("sig", "", "signature hex as string")
	patchPtr := flag.String("patch", "", "patch hex as string")
	flag.Parse()

	if len(os.Args) < 4 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// "707172" []byte("pqr")
	sig, err := hex.DecodeString(*sigPtr)
	if err != nil {
		log.Fatal("unable decode signature: ", err.Error())
	}

	// "313233" []byte("123")
	patch, err := hex.DecodeString(*patchPtr)
	if err != nil {
		log.Fatal("unable decode patch: ", err.Error())
	}

	if len(sig) != len(patch) {
		log.Fatal("signature and patch are not the same legnth")
	}

	input, err := ioutil.ReadFile(*inPtr)
	if err != nil {
		log.Fatal("unable to open input file: ", err.Error())
	}
	output := input

	log.Printf("Signature: %x", sig)
	log.Printf("Patch: %x", patch)
	for i := 0; i < len(input); i++ {
		found := sigCheck(input[i:i+len(sig)], sig)
		if found {
			log.Printf("Found signature at %#x!", i)
			doPatch(output, patch, i)
			break
		}
	}
	out, err := os.Create(*outPtr)
	if err != nil {
		log.Fatal("unable to create output file: ", err.Error())
	}
	defer out.Close()

	_, err = io.Copy(out, bytes.NewReader(output))
	if err != nil {
		log.Fatal("unable to save file: ", err.Error())
	}
	log.Println("Saved output!")
}
