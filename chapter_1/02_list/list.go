package main

import (
	"container/list"
	"fmt"
)

func main() {
	var list list.List
	list.PushBack("Alpha")
	list.PushBack("Beta")
	list.PushBack(1)

	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
