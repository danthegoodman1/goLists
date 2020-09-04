package main

import (
	"container/list"
	"fmt"
	"github.com/MauriceGit/skiplist"
	"github.com/collinglass/bptree"
	"github.com/google/btree"
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

type Element struct {
	key  int
	data string
}

func stringToNum(input string) (num int) {
	runes := []rune(input)
	e := 0
	for i := 0; i < len(runes); i++ {
		e += int(runes[i])
	}
	return e
}

func NewElement(data string) (e Element) {
	i := stringToNum(data)
	return Element{i, fmt.Sprintf("hey%d", i)}
}

// Implement the interface used in skiplist
func (e Element) ExtractKey() float64 {
	return float64(e.key)
}
func (e Element) String() string {
	return e.data
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

func testTheirs(testSize int, searchItem int) (taken time.Duration, traverseTime time.Duration) {
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
	return totalDiff, diff
}

func testSkip(testSize int, searchItem int) (taken time.Duration) {
	fmt.Println("Testing skiplist...")
	totalTime := time.Now()
	list := skiplist.New()
	// Insert some elements
	for i := 0; i < testSize; i++ {
		newE := NewElement(string(i))
		list.Insert(newE)
	}

	end := time.Now()
	diff := end.Sub(totalTime)
	log.Printf("LIST CONSTRUCTION COMPLETE - TIME SPENT: %v", diff)

	start := time.Now()
	// Find an element
	newE := NewElement(string(searchItem))
	if e, ok := list.Find(newE); ok {
		fmt.Println(e.GetValue())
	}

	end = time.Now()
	diff = end.Sub(start)

	log.Printf("LIST SEARCH COMPLETE - TIME SPENT: %v", diff)

	// Delete all elements
	// for i := 0; i < testSize; i++ {
	//     list.Delete(Element(i))
	// }
	totalDiff := end.Sub(totalTime)
	log.Printf("TEST SKIPLIST COMPLETE - TIME SPENT: %v", totalDiff)

	// smallest item:
	sm := list.Next((list.GetSmallestNode()))
	fmt.Println(sm.GetValue().String())

	return diff
}

func testBTree(testSize int, searchItem int) (taken time.Duration) {
	fmt.Println("Testing btree...")
	totalTime := time.Now()
	list := btree.New(2)
	// Insert some elements
	for i := 0; i < testSize; i++ {
		list.ReplaceOrInsert(btree.Int(i))
	}

	end := time.Now()
	diff := end.Sub(totalTime)
	log.Printf("LIST CONSTRUCTION COMPLETE - TIME SPENT: %v", diff)

	start := time.Now()
	// Find an element
	list.Get(btree.Int(searchItem))

	end = time.Now()
	diff = end.Sub(start)

	log.Printf("LIST SEARCH COMPLETE - TIME SPENT: %v", diff)

	// Delete all elements
	// for i := 0; i < testSize; i++ {
	//     list.Delete(Element(i))
	// }
	totalDiff := end.Sub(totalTime)
	log.Printf("TEST BTREE COMPLETE - TIME SPENT: %v", totalDiff)

	return diff
}

func testBPTree(testSize int, searchItem int) (taken time.Duration) {
	fmt.Println("Testing b+tree...")
	totalTime := time.Now()
	list := bptree.NewTree()
	// Insert some elements
	for i := 0; i < testSize; i++ {
		err := list.Insert(i, []byte(string(i)))
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	end := time.Now()
	diff := end.Sub(totalTime)
	log.Printf("LIST CONSTRUCTION COMPLETE - TIME SPENT: %v", diff)

	start := time.Now()
	// Find an element
	_, err := list.Find(searchItem, false)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}

	end = time.Now()
	diff = end.Sub(start)

	log.Printf("LIST SEARCH COMPLETE - TIME SPENT: %v", diff)

	// Delete all elements
	// for i := 0; i < testSize; i++ {
	//     list.Delete(Element(i))
	// }
	totalDiff := end.Sub(totalTime)
	log.Printf("TEST B+TREE COMPLETE - TIME SPENT: %v", totalDiff)

	return diff
}

func main() {
	log.Println("Going")
	testSize := 100000000
	searchItem := 453012
	theirTime, traverseTime := testTheirs(testSize, searchItem)
	myTime := testMine(testSize, searchItem)
	skipTime := testSkip(testSize, searchItem)
	bTime := testBTree(testSize, searchItem)
	bpTime := testBPTree(testSize, searchItem)
	log.Printf("My total speed advantage: %vx", float32(theirTime.Nanoseconds())/float32(myTime.Nanoseconds()))
	log.Printf("Skiplist search advantage: %vx", float32(traverseTime.Nanoseconds())/float32(skipTime.Nanoseconds()))
	log.Printf("BTree search advantage: %vx", float32(traverseTime.Nanoseconds())/float32(bTime.Nanoseconds()))
	log.Printf("B+Tree search advantage: %vx", float32(traverseTime.Nanoseconds())/float32(bpTime.Nanoseconds()))
}
