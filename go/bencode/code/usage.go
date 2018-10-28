// usage.go.

// 'bencode' Package Usage.

package main

import (
	"fmt"
	"github.com/legacy-vault/library/go/bencode"
	"log"
	"time"
)

func main() {

	var ba []byte
	var dict []bencode.DictionaryItem
	var err error
	var fileName string
	var list []interface{}
	var obj *bencode.DecodedObject

	fileName = "../data/" +
		"xubuntu-17.10.1-desktop-amd64.iso.torrent"

	obj, err = bencode.ParseFile(fileName)
	checkError(err)
	fmt.Printf(
		"Date=[%s] File=['%s'] BTIH=[%s].\r\n",
		time.Unix(obj.DecodeTimestamp, 0).Format(time.RFC3339),
		obj.FilePath,
		obj.BTIH.Text,
	)

	// Try to encode a simple Integer.
	ba, err = bencode.Encode(uint64(123))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple Byte String.
	ba, err = bencode.Encode([]byte("Hello, World!"))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple List.
	list = []interface{}{
		[]byte("Hello, World!"),
		uint64(2000),
	}
	ba, err = bencode.Encode(list)
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple Dictionary.
	list = []interface{}{
		[]byte("Hello, World!"),
		uint64(2000),
	}
	dict = []bencode.DictionaryItem{
		bencode.DictionaryItem{
			Key:   []byte("bytes"),
			Value: []byte{'X', '\r', '\n', 'Y', 'Z'},
		},
		bencode.DictionaryItem{
			Key:   []byte("name"),
			Value: []byte("John"),
		},
		bencode.DictionaryItem{
			Key:   []byte("texts"),
			Value: list,
		},
	}
	ba, err = bencode.Encode(dict)
	checkError(err)
	fmt.Println(string(ba))
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
