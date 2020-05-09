package lists

/*
 * MIT License
 *
 * Copyright (c) 2020
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
 * BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR  OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
import (
	"fmt"
)

type Node struct {
	next  *Node
	Value int
}

type List struct {
	first *Node
}

func NewList() *List {
	return &List{}
}

func (l *List) AddFirst(i int) {

}

func (l *List) AddLast(value int) {
	if l.first == nil {
		l.first = &Node{Value: value}
		return
	}
	temp := l.first
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &Node{Value: value}
}

func (l *List) Print() {
	temp := l.first
	for temp != nil {
		fmt.Printf("%v\n", temp.Value)
		temp = temp.next
	}
}
