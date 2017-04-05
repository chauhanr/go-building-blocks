# go-bb [golang building blocks project.]

This project is for checking in:
1. Sample code.
2. Examples during topics mastery.
3. Components that can be used later.

Also I write test cases and classes for all of them so that can write better TDD code.

Topics covered.
1. Structs
2. Pointers
3. Files
4. Concurrency
5. OS module and command execution.

# Testing in Go.
The testing any component in go can be done using the testing package.
1. The tests must be in a file of format *_test.go
2. the function for test must be like name (t *testing.T)
3. We need to call the function we need to test in the testing function defined.

Every example in this project has a test file and a test case. This allows for the functions to be prepared individually and we can reuse them. We can also run the entire suite all at once or selectively like:

* go test -v -run=regrep    - this will try and match all the test cases that match the regular expression given

e.g.
$ go test -v -run 'Mul'   -- will run all the tests with have Mul in their name.

** Good Practices **
1. create your test classes before you write your components.
2. have testing table data to test different scenarios e.g. like the multiplier_test.go under the functions package.
3. You can benchmark your test cases too.

# Maps
maps are dictionaries in golang they hold key and value pairs. Maps are declared using **var m map[keytype]valuetype**

maps when not initialized using make will always be nil value as in the case with the m map above.

Key value in map are any variables that can be used with == and != operators and therefore keys can only be string, int, float but cannot be slices, arrays or structs. If you need to use struct as a key value then the struct must implement a Key() and Hash() function. However maps can take interfaces as values.

Looking up maps is faster than linear search but is slower than indexed search such as array and slices.

When accessing the map values using keys we get back a variable of type defined for the value type. If there is no value w.r.t the key in the map then the map returns a zero value of the value type defined.

## Patterns  
**Branching using functions as values in maps**

One of the patterns that we can use is to have a function being stored against a key in the map. So when we encounter a key we can get the function and execute it at the time. This helps us create branching logic based on the key we encounter.

**Checking the presence of the key**

There is a good way to do that using the if construct

```golang

   if _, isPresent := mapName[key]; isPresent {
      // do what we need to when a key is present
   }

```
