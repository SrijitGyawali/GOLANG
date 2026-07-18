package main

import "fmt"

func main() {
	// --- Arrays: FIXED size. You rarely use these directly in Go. ---
	var arr [3]int      // an array of 3 ints, all zero
	arr[0] = 10         // set by index
	arr[1] = 20
	fmt.Println("array:", arr, "length:", len(arr))

	// --- Slices: DYNAMIC size. This is what you'll use 99% of the time. ---
	nums := []int{1, 2, 3} // note: no size inside []
	fmt.Println("slice:", nums)

	// append adds elements (and grows the slice as needed).
	nums = append(nums, 4, 5)
	fmt.Println("after append:", nums, "length:", len(nums))

	// Slicing: nums[start:end] -> from start up to (but NOT including) end.
	fmt.Println("nums[1:3]:", nums[1:3]) // -> [2 3]
	fmt.Println("nums[:2]:", nums[:2])   // from start -> [1 2]
	fmt.Println("nums[3:]:", nums[3:])   // to the end  -> [4 5]

	// make() creates a slice with a length up front.
	scores := make([]int, 3) // [0 0 0]
	scores[0] = 100
	fmt.Println("made slice:", scores)

	// Loop over a slice with range.
	for i, v := range nums {
		fmt.Printf("index %d = %d\n", i, v)
	}
}
