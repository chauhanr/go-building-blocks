# Golang Interfaces

The examples in this package include
1. Sorter - generic way to sort int or string arrays. The sorting algorithm remains that same.  
2. Miner - generic way to determine the minimum element in the int or float array.

## Sorter Interface

```
   type Sorter interface {
       Len() int
       Less(i, j int) bool // This will give true if the element at index i is less than one at index j
       Swap(i, j int) // Swap the elements at i and j
   }
```
The sorter interface can now be used for array elements of different types

```
 type IntArray []int  
 // or
 type StringArray []string
 ```
There will be another function called the Sort Method that will have the algorithm to sort the data which will be of Sorter type.

```
func Sort( data Sorter){
    for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}
```

## Miner Interface

Similar to the pattern above we have the Miner interface. Here too we have Miner interface

```
type Element interface{}

type Miner interface {
      Less(i, j int) bool // this determines if element at i is less than j
      Len() int  // returns the array length
      Get(index i) Element
}
```

The Min method that has the algorithm to find the smallest element

```
func Min(data Miner) Element{
    if data.Len() != 0 {
		min := 0
		for i := 1; i < data.Len(); i++ {
			if data.Less(i, min) {
				min = i
			}
		}
		return data.Get(min)
	}
	return nil
}
```

This is an important pattern where we are trying to abstract an algorithm from the data type. This will be important pattern when we implement data structures like Tree, Stacks, Linked Lists and Graph.

## Reflection

Learned how to use the reflection library in Go.
There are two main methods in the reflect package
1. **reflect.ValueOf(x)** - returns value of the interface passed to the method e.g. it returns <float64 Value> for a float64, or float32 types.
2. **reflect.TypeOf(x)** - returns the typeof the variable passed like float64 or int16, int8 or even a structures.Person (struct type)
3. **reflect.TypeOf(x).Kind()** - this method will give the kind of type that we have. The difference between Type and Kind comes into the picture when we have a struct Type => "structures.Person" and Kind => "struct"

for more examples and method used use the reflect_test.go file in this package.
