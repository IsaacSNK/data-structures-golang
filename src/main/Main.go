package main

import "lists"

func main() {
	var list *lists.List = lists.NewList()
	list.AddLast(1)
	list.AddLast(2)
	list.Print()
	list.AddFirst(0)
	list.Print()
}
