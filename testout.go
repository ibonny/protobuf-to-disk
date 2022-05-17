package main

import (
	"fmt"
	"io/ioutil"
	"log"
	pb "protobuf-to-disk/testrecord"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("OK, we're here.")

	p := &pb.TestRecord{
		Firstname: "Ian",
		Lastname:  "Bonnycastle",
		Age:       51,
	}

	// Write the new address book back to disk.
	var out []byte
	var err error

	if out, err = proto.Marshal(p); err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	fname := "test.dat"

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	var in []byte

	if in, err = ioutil.ReadFile(fname); err != nil {
		log.Fatalln("Error reading file:", err)
	}

	var newRec pb.TestRecord

	if err = proto.Unmarshal(in, &newRec); err != nil {
		log.Fatalln("Cannot process data:", err)
	}

	fmt.Println(newRec)
}
