# PersistLink

PersistLink is a Go library for working with immutable linked lists. It provides a simple, functional API to create, manipulate, and combine linked lists efficiently


## Features
- **Empty List:** Create an empty list of any type.
- **Append Elements:** Add elements to the list.
- **Concatenate Lists:** Combine two linked lists into one.
- **Convert to Slice:** Convert the linked list into a Go slice.
- **Check for Emptiness:** Determine if a list is empty.

## Usage

Quick usage example : 

```go

	// Create a list of strings
	list := lib.Empty[string]().Append("apple").Append("banana").Append("cherry")

	// Create another list of strings
	anotherList := lib.Empty[string]().Append("date").Append("elderberry")

	// Concatenate the two lists
	concatenatedList := list.Concat(anotherList)
```


## License

This project is licensed under the MIT License.
