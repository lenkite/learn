package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	cet, _ := time.LoadLocation("CET")
	t1 := time.Now()
	t1CET := t1.In(cet)
	t2CET := t1.AddDate(0, 0, 10).In(cet)
	log.Println("now       CET:", t1CET)
	log.Println("now       UTC:", t1.UTC())
	log.Println("next week CET:", t2CET)
	log.Println("next week UTC:", t2CET.UTC())
}
