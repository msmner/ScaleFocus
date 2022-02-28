package main

import "fmt"

func main() {
	magicList := MagicList{}
	add(&magicList, 10)
	add(&magicList, 12)
	add(&magicList, 14)
	result := toSlice(&magicList)
	fmt.Println(result)
}

type Item struct {
	Value    int
	PrevItem *Item
}

type MagicList struct {
	LastItem *Item
}

func add(l *MagicList, value int) {
	if l.LastItem == nil {
		item := &Item{value, nil}
		l.LastItem = item
	} else {
		item := &Item{value, l.LastItem}
		l.LastItem = item
	}
}

func toSlice(l *MagicList) []int {
	result := []int{}
	for {
		if l.LastItem.PrevItem == nil {
			return append([]int{l.LastItem.Value}, result...)
		} else {
			result = append([]int{l.LastItem.Value}, result...)
			l.LastItem = l.LastItem.PrevItem
		}
	}
}
