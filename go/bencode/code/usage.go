// usage.go.

// 'bencode' Package Usage.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/legacy-vault/library/go/bencode"
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

	// Try to encode a simple unsigned Integer.
	ba, err = bencode.Encode(uint(111))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple signed Integer.
	ba, err = bencode.Encode(int(-222))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple unsigned 64-Bit Integer.
	ba, err = bencode.Encode(uint64(111))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple signed 64-Bit Integer.
	ba, err = bencode.Encode(int64(-222))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple unsigned 32-Bit Integer.
	ba, err = bencode.Encode(uint32(111))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple signed 32-Bit Integer.
	ba, err = bencode.Encode(int32(-222))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple unsigned 16-Bit Integer.
	ba, err = bencode.Encode(uint16(16001))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple signed 16-Bit Integer.
	ba, err = bencode.Encode(int16(-16001))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple unsigned 8-Bit Integer.
	ba, err = bencode.Encode(uint8(8))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple signed 8-Bit Integer.
	ba, err = bencode.Encode(int8(-8))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple Byte String.
	ba, err = bencode.Encode([]byte("Hello, World!"))
	checkError(err)
	fmt.Println(string(ba))

	// Try to encode a simple String.
	ba, err = bencode.Encode("I am alive")
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
