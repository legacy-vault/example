// usage.go.

// Simple File Storage Usage.

package main

import (
	"fmt"
	"log"

	"github.com/legacy-vault/library/go/simple_file_storage"
)

func main() {

	var contents []byte
	var err error
	var storageA *sfstorage.Storage
	var storageB *sfstorage.Storage

	storageA, err = sfstorage.New("/tmp/abc////")
	checkError(err)

	fmt.Println("Root Path:", storageA.GetRootPath())

	contents, err = storageA.GetFileContents("///1.txt////")
	checkError(err)
	fmt.Println("Contents:", string(contents))

	storageB, err = sfstorage.New("/tmp/new_folder/////")
	checkError(err)

	fmt.Println("Root Path:", storageB.GetRootPath())
}

func checkError(err error) {

	if err != nil {
		log.Fatalln(err)
	}
}
