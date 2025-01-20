package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Basic struct definition
type Rectangle struct {
	width  float64
	height float64
}

// Method for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Interface definition
type Shape interface {
	Area() float64
}

// Function that accepts an interface
func PrintArea(s Shape) {
	fmt.Printf("Area: %0.2f\n", s.Area())
}

// Function with multiple return values
func divideAndRemainder(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, a % b, nil
}

// Function that demonstrates goroutines and channels
func concurrentSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // Send sum to channel
}

func main() {
	// Variables and basic data types
	var a int = 10
	b := 20 // Short variable declaration
	c := 3.14
	d := "Hello, Go!"
	fmt.Printf("a = %d, b = %d, c = %0.2f, d = %s\n", a, b, c, d)

	// Arrays and slices
	arr := [3]int{1, 2, 3}
	slice := []int{4, 5, 6}
	slice = append(slice, 7)
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	// Maps
	m := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	m["cherry"] = 3
	fmt.Println("Map:", m)

	// Control structures
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// Switch statement
	switch day := time.Now().Weekday(); day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Functions with multiple return values
	result, remainder, err := divideAndRemainder(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 divided by 3 is %d with remainder %d\n", result, remainder)
	}

	// Structs and interfaces
	r := Rectangle{width: 3, height: 4}
	PrintArea(r)

	// Anonymous struct
	circle := struct {
		radius float64
	}{
		radius: 5,
	}

	// Anonymous function (closure) that satisfies the Shape interface
	PrintArea(Shape(struct{ radius float64 }{radius: circle.radius}))

	// Goroutines and channels
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go concurrentSum(s[:len(s)/2], c)
	go concurrentSum(s[len(s)/2:], c)
	part1, part2 := <-c, <-c // Receive from channel
	fmt.Println("Concurrent sum result:", part1+part2)

	// WaitGroup for synchronization
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines completed")

	// Defer statement
	defer fmt.Println("This will be printed last")

	// Panic and recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// This will cause a panic
	panic("This is a panic situation!")
}