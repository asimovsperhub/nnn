package mysqldb

import (
	"fmt"
	"log"
	"testing"
)

type Res struct {
	Addr string `json:"addr"`
}

func TestName(t *testing.T) {
	cc := GetMdb()
	var c Res
	log.Println("MDB", cc)
	err := cc.Mysql.QueryRow(fmt.Sprintf("SELECT addr FROM Address WHERE addr='%s' LIMIT 1;", "ll")).Scan(&c.Addr)
	if err != nil {
		log.Println(err)
	}
	if c.Addr == "" {
		log.Println("not in")
	}
	log.Println(fmt.Sprintf("SELECT 1 FROM Address WHERE addr='%s' LIMIT 1;", "ll"))
	log.Println(c)
}
