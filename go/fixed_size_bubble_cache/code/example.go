// example.go.

package main

import (
	"log"
	"time"

	"github.com/legacy-vault/library/go/fixed_size_bubble_cache"
)

func main() {

	var c *cache.Cache

	c = cache.New(5, 30)
	log.Println(c.GetRecordTTL())

	c.AddRecord("U-1", "John")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second)
	c.AddRecord("U-2", "Peter")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second)
	c.AddRecord("U-3", "Maria")
	log.Println(c.EnlistAllRecordValues())

	printDataByUID(c, "U-2")
	c.SetRecordTTL(10)
	log.Println(
		c.EnlistAllRecordValues(),
		c.RecordUIDExists("U-1001"),
		c.GetRecordTTL(),
	)

	time.Sleep(time.Second * 2)
	c.AddRecord("U-4", "Fox")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddRecord("U-5", "Go")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddRecord("U-6", "Six")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddRecord("U-7", "777")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 5)
	log.Println(c.EnlistAllRecordValues())

	// Let's get an outdated Record and see that it is removed.
	printDataByUID(c, "U-3")
	log.Println(c.EnlistAllRecordValues())

	c.AddRecord("U-4", "Fox+")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddRecord("U-7", "777+")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.Clear()
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddRecord("N-1", "111")
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddRecord("N-2", "222")
	log.Println(c.EnlistAllRecordValues())
}

func printDataByUID(cache *cache.Cache, uid cache.RecordUID) {

	var data interface{}
	var exists bool

	data, exists = cache.GetRecordDataByUID(uid)
	if exists {
		log.Printf("Record with ID='%v' is [%v].\r\n", uid, data)
	} else {
		log.Printf("Record with ID='%v' does not exist.\r\n", uid)
	}
}
