package main

import "fmt"

var head *node
var length int

type node struct {
	data  int
	left  *node
	right *node
}

func printStruct() {
	if head != nil {
		temp := head
		for temp.right != nil {
			fmt.Print(temp.data)
			fmt.Print(" ")
			temp = temp.right
		}
		fmt.Print(temp.data)
		fmt.Println(" ")
	}
}

func stackAdd(data int) {
	if head == nil {
		temp := node{data, nil, nil}
		head = &temp
		length = 1
		return
	}
	temp := node{data, nil, head}
	head.left = &temp
	head = &temp
	length++
}

func stackRem() {
	if head == nil {
		return
	}
	if head.right == nil {
		head = nil
		length = 0
		return
	}
	head = head.right
	head.left.right = nil
	head.left = nil
	length--
}

func queueAdd(data int) {
	if head == nil {
		temp := node{data, nil, nil}
		head = &temp
		length = 1
		return
	}
	temp := node{data, nil, head}
	head.left = &temp
	head = &temp
	length++
}

func queueRem() {
	if head == nil {
		return
	}
	if head.right == nil {
		head = nil
		length = 0
		return
	}
	temp := head
	for temp.right != nil {
		temp = temp.right
	}
	temp.left.right = nil
	temp.left = nil
	length--
}

func decPushFront(data int) {
	if head == nil {
		temp := node{data, nil, nil}
		head = &temp
		length = 1
		return
	}
	temp := node{data, nil, head}
	head.left = &temp
	head = &temp
	length++
}

func decPushBack(data int) {
	if head == nil {
		temp := node{data, nil, nil}
		head = &temp
		length = 1
		return
	}

	tempPointer := head
	for tempPointer.right != nil {
		tempPointer = tempPointer.right
	}
	temp := node{data, tempPointer, nil}
	tempPointer.right = &temp
	length++
}

func decPopFront() {
	if head == nil {
		return
	}
	if head.right == nil {
		head = nil
		length = 0
		return
	}
	head = head.right
	head.left.right = nil
	head.left = nil
	length--
}

func decPopBack() {
	if head == nil {
		return
	}
	if head.right == nil {
		head = nil
		length = 0
		return
	}
	temp := head
	for temp.right != nil {
		temp = temp.right
	}
	temp.left.right = nil
	temp.left = nil
	length--
}

func main() {

	stackRem()
	stackAdd(10)
	stackAdd(20)
	printStruct()
	stackRem()
	printStruct()
	stackRem()
	printStruct()

	queueRem()
	queueAdd(10)
	queueAdd(20)
	printStruct()
	queueRem()
	printStruct()
	queueRem()
	printStruct()

	decPopFront()
	decPopBack()
	decPushBack(20)
	decPushFront(10)
	decPushBack(30)
	decPushFront(0)
	printStruct()
	decPopBack()
	printStruct()
	decPopFront()
	printStruct()
}
