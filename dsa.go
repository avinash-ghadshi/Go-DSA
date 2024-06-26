package main

import (
	ds "dsa/datastructures"
)

func main() {
	// ilist := algo.Ilist{
	// 	List:    []int{1, 2, 3, 4, 5, 6},
	// 	Element: 3,
	// }
	// msg, err := algo.LinearSearch(ilist)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(msg)
	// }

	// slist := algo.Slist{
	// 	List:    []string{"a", "b", "c", "d"},
	// 	Element: "p",
	// }
	// msg, err = algo.LinearSearch(slist)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(msg)
	// }

	head := ds.NewNode("HEAD")
	tail := head.Add(head)
	for i := 0; i < 5; i++ {
		x := ds.NewNode(i)
		tail = x.Add(tail)
	}

	ds.Traverse(head)

}
