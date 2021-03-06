package main

import (
	"fmt"
	"container/list"
)

// 2 用map和链表实现LRU缓存
var cache map[int]string
var lst *list.List

const CAP = 10 // 定义缓存容量的上限

func init() {
	cache = make(map[int]string, CAP)
	lst = list.New()
}

func readFromDisk(key int) string {
	return "china"
}

func read(key int) string {
	if v, exists := cache[key]; exists {  // 命中缓存
		head := lst.Front()
		notFound := false
		for {
			if head == nil {
				notFound  = true
				break
			}
			if head.Value.(int) == key { // 从链表里找到相应的key
				lst.MoveToFront(head)  // 把key移动到链表的头部
				break
			} else {
				head = head.Next()
			}
		}
		if !notFound {  // 正常情况下不会发生这种情况
			lst.PushFront(key)
		}
		return v
	} else { // 如果没有命中缓存
		v = readFromDisk(key) // 从磁盘中读取数据
		cache[key] = v        // 放入缓存
		lst.PushFront(key)    // 放入链表头部
		if len(cache) >= CAP {// 缓存已满
			tail := lst.Back()
			delete(cache, tail.Value.(int))  // 从缓存移除很久不用的元素
			lst.Remove(tail)                 // 从链表中删除最后一个元素
			fmt.Printf("Remove %d from cache\n", tail.Value.(int))
		}
		return v
	}
}

func TraversList(lst *list.List) {
	head := lst.Front()
	for head.Next() != nil {
		fmt.Printf("%v ", head.Value)
		head = head.Next()
	}
	fmt.Println(head.Value)
}

func main() {
	for i := 0; i < 15; i++ {
		_ = read(i)
	}
	for k, v := range cache {
		fmt.Printf("%d: %s\n", k, v)
	}
	fmt.Println("*****************")
	TraversList(lst)
}