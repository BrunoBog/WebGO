package main

import (
	"net"
	"io/ioutil"
	"log"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil{
		panic(err)
	}

	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil{
		panic(err)
	}

	log.Print(string(bs))
}
