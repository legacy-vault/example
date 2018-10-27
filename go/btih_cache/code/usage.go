// usage.go.

package main

import (
	"log"
	"time"

	"github.com/legacy-vault/library/go/btih_cache"
)

const CacheCapacity = 3
const CacheRecordTTLSec = 15
const CacheRootFoler = "/tmp/btih"

func main() {

	var bc *btih_cache.BTIHCache
	var err error
	var ok bool

	bc, err = btih_cache.New(
		CacheCapacity,
		CacheRecordTTLSec,
		CacheRootFoler,
	)
	checkError(err)

	showFileByBTIH(bc, "AAABBB")
	showFileCachedStatus(bc, "AAABBB")
	showFileCachedStatus(bc, "AAAxxx")

	showFileByBTIH(bc, "CCCDDD")
	showFileByBTIH(bc, "OOOPPP")
	showFileCachedStatus(bc, "OOOPPP")

	showFileByBTIH(bc, "MMMNNN")
	showFileCachedStatus(bc, "MMMNNN")
	showFileCachedStatus(bc, "AAABBB") // Removed as a Tail Record.
	// Cache now has: [HEAD]=[M]-[O]-[C]=[TAIL].

	log.Println("Waiting 20 Seconds...")
	time.Sleep(time.Second * 20)
	// [C] Element still exists in the List but is now outdated.
	showFileCachedStatus(bc, "CCCDDD")

	// This Request requires File Storage to be used, as the Record is outdated.
	showFileByBTIH(bc, "CCCDDD")

	ok = bc.Stop()
	if !ok {
		panic("Stop Error")
	}
	// Returns an empty Array of Bytes, as the Cache has been stopped.
	showFileByBTIH(bc, "MMMNNN")
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func showFileByBTIH(bc *btih_cache.BTIHCache, btih string) {

	var ba []byte
	var err error

	ba, err = bc.GetFileByBTIH(btih)
	if err != nil {
		log.Println(err)
	}
	log.Println(btih, "->", len(ba), "Bytes.")
}

func showFileCachedStatus(bc *btih_cache.BTIHCache, btih string) {

	var isCached bool

	isCached = bc.IsCached(btih)
	log.Println(("'" + btih + "'"), "Cached Status:", isCached)
}
