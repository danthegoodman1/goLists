package main

import (
	"container/list"
	"fmt"
	"log"
	"time"
)

type listItem struct {
	next *listItem
	prev *listItem
	data interface{}
}

type bigList struct {
	len   int32
	first *listItem
	last  *listItem
}

func testMine(testSize int, searchItem int) (taken time.Duration) {
	fmt.Println("Testing mine...")
	totalTime := time.Now()
	l := bigList{0, nil, nil}
	l.first = &listItem{nil, nil, 0}
	cursor := l.first
	// Generate large list
	for i := 1; i < testSize; i++ {
		cursor.next = &listItem{nil, nil, i}
		l.last = cursor.next
		l.len++
		cursor = cursor.next
	}
	end := time.Now()
	diff := end.Sub(totalTime)
	log.Printf("LIST CONSTRUCTION COMPLETE - TIME SPENT: %v", diff)
	// Iterate over entire list
	start := time.Now()
	for cursor = l.first; cursor != nil; cursor = cursor.next {
		// Do nothing
	}
	end = time.Now()
	diff = end.Sub(start)
	log.Printf("LIST TRAVERSAL COMPLETE - TIME SPENT: %v", diff)

	totalDiff := end.Sub(totalTime)
	log.Printf("TEST MINE COMPLETE - TIME SPENT: %v", totalDiff)
	return totalDiff
}

func testTheirs(testSize int, searchItem int) (taken time.Duration) {
	fmt.Println("Testing theirs...")
	totalTime := time.Now()
	l := list.New()
	// Generate large list:
	for i := 0; i < testSize; i++ {
		l.PushBack(i)
	}
	end := time.Now()
	diff := end.Sub(totalTime)
	log.Printf("LIST CONSTRUCTION COMPLETE - TIME SPENT: %v", diff)
	// Iterate over entire list
	start := time.Now()
	for cursor := l.Front(); cursor != nil; cursor = cursor.Next() {
		// Do nothing
	}
	end = time.Now()
	diff = end.Sub(start)
	log.Printf("LIST TRAVERSAL COMPLETE - TIME SPENT: %v", diff)

	totalDiff := end.Sub(totalTime)
	log.Printf("TEST THEIRS COMPLETE - TIME SPENT: %v", totalDiff)
	return totalDiff
}

func main() {
	log.Println("Going")
	testSize := 10000000
	searchItem := 453012
	theirTime := testTheirs(testSize, searchItem)
	myTime := testMine(testSize, searchItem)
	log.Printf("My speed advantage: %vx", float32(theirTime.Nanoseconds())/float32(myTime.Nanoseconds()))
}
