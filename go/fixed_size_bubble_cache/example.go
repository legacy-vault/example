// example.go.

package main

import (
	"github.com/legacy-vault/library/go/fixed_size_bubble_cache"
	"log"
	"time"
)

func main() {

	var c *cache.Cache

	c = cache.New(5, 30)
	log.Println(c.GetRecordTTL())

	c.AddARecord(&cache.Record{UID: "U-1", Data: "John"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second)
	c.AddARecord(&cache.Record{UID: "U-2", Data: "Peter"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second)
	c.AddARecord(&cache.Record{UID: "U-3", Data: "Maria"})
	log.Println(c.EnlistAllRecordValues())

	printDataByUID(c, "U-2")
	c.SetRecordTTL(10)
	log.Println(
		c.EnlistAllRecordValues(),
		c.RecordUIDExists("U-1001"),
		c.GetRecordTTL(),
	)

	time.Sleep(time.Second * 2)
	c.AddARecord(&cache.Record{UID: "U-4", Data: "Fox"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddARecord(&cache.Record{UID: "U-5", Data: "Go"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddARecord(&cache.Record{UID: "U-6", Data: "Six"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddARecord(&cache.Record{UID: "U-7", Data: "777"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 5)
	log.Println(c.EnlistAllRecordValues())

	// Let's get an outdated Record and see that it is removed.
	printDataByUID(c, "U-3")
	log.Println(c.EnlistAllRecordValues())

	c.AddARecord(&cache.Record{UID: "U-4", Data: "Fox+"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddARecord(&cache.Record{UID: "U-7", Data: "777+"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.Clear()
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddARecord(&cache.Record{UID: "N-1", Data: "111"})
	log.Println(c.EnlistAllRecordValues())

	time.Sleep(time.Second * 2)
	c.AddARecord(&cache.Record{UID: "N-2", Data: "222"})
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
