package main

import "fmt"

type Counter struct {
	Value int
}

func main() {
	// Pointers store memory addresses of values.
	// &value gets the address, and *pointer dereferences it.
	number := 10
	numberPtr := &number

	fmt.Println("number:", number)
	fmt.Println("numberPtr points to:", *numberPtr)

	// Changing the dereferenced pointer changes the original value.
	*numberPtr = 25
	fmt.Println("number after pointer write:", number)

	// A pointer can also point to a struct.
	counter := Counter{Value: 1}
	increment(&counter)
	fmt.Println("counter after increment:", counter.Value)

	// Go is still call by value:
	// the pointer itself is copied into the parameter,
	// but both copies point to the same underlying value.
	swapLocalPointer(numberPtr)
	fmt.Println("number after swapLocalPointer:", number)

	// Slices and maps are both passed by value, but they behave differently.
	// A slice is a small header: pointer + len + cap.
	values := []int{1, 2, 3}
	changeSliceElement(values)
	fmt.Println("slice after element change:", values)

	// Appending may allocate a new backing array, so the caller's slice header
	// may not see the new length unless the returned slice is assigned back.
	values = appendValue(values, 4)
	fmt.Println("slice after append:", values)

	// A map value points to runtime-managed hash table data.
	// Mutating map contents inside a function is visible to the caller.
	counts := map[string]int{"go": 1}
	changeMap(counts)
	fmt.Println("map after change:", counts)

	// Slices as buffers:
	// reuse the same backing array instead of allocating every time.
	buffer := make([]byte, 0, 8)
	buffer = append(buffer, 'G', 'o')
	fmt.Println("buffer contents:", string(buffer), "len/cap:", len(buffer), cap(buffer))

	// Emptying with [:0] keeps the allocated backing array for reuse.
	buffer = buffer[:0]
	buffer = append(buffer, 'R', 'e', 'u', 's', 'e')
	fmt.Println("reused buffer:", string(buffer), "len/cap:", len(buffer), cap(buffer))

	// Copying only the active window can avoid holding on to a very large array.
	// This is a practical way to reduce accidental memory retention.
	large := make([]byte, 0, 1024)
	large = append(large, []byte("important")...)

	workingView := large[:9]
	tightCopy := append([]byte(nil), workingView...)

	// Now tightCopy only keeps the bytes it needs.
	fmt.Println("workingView len/cap:", len(workingView), cap(workingView))
	fmt.Println("tightCopy len/cap:", len(tightCopy), cap(tightCopy))
}

func increment(counter *Counter) {
	// Pointer parameters let the function mutate the caller's value.
	counter.Value++
}

func swapLocalPointer(ptr *int) {
	// Reassigning the local pointer variable does not retarget the caller's variable.
	other := 999
	ptr = &other
	*ptr = 1000
}

func changeSliceElement(values []int) {
	if len(values) > 0 {
		values[0] = 99
	}
}

func appendValue(values []int, next int) []int {
	return append(values, next)
}

func changeMap(values map[string]int) {
	values["go"]++
	values["memory"] = 2
}
