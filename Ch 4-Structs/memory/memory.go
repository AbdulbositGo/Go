package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

type empty struct{}

func main() {

	typ1 := reflect.TypeOf(contact{})
	typ2 := reflect.TypeOf(perms{})

	fmt.Printf("Contact struct is %d bytes\n", typ1.Size())
	fmt.Printf("perms struct is %d bytes\n", typ2.Size())
	fmt.Printf("perms struct is %d bytes\n", reflect.TypeOf(empty{}).Size())

}
