package main

import "fmt"

func main() {
	var hello string = "Hello world"
	fmt.Println(hello) // Hello world

	hello = "Hello Go"
	fmt.Println(hello) // Hello Go

	fmt.Println(hello) // Go Go Go Ole Ole Ole

	nums1 := [4] int {3, 4, 5, 6}
    nums2 := [4] int {3, 4, 5}
 
    fmt.Println("nums1 == nums2:", nums1==nums2) // false
 
    nums3 := [3][2] int { {2}, {5}, }
    nums4 := [3][2] int { {2}, {5}, }
    fmt.Println("nums3 == nums4:", nums3==nums4)   // false
 
    nums5 := [4] int{3, 4, 5, 0}
    fmt.Println("nums2 == nums5:", nums2==nums5)  // true
}
