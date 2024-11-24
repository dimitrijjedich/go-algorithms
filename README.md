# Sorting Algorithms in Go

This repository contains implementations of various sorting algorithms in Go as a hands-on learning project. The projects purpose is simply to learn the syntax, how to structure a Go project and implement and run tests.

This README is a simple writeup of the steps I took to get to the current point. When you try to follow along you may have to change some stuff according to you (the module name for example).

---

## Table of Contents

1. [Setup Instructions](#setup-instructions)
2. [Project Structure](#project-structure)
3. [Writing Code & Test](#writing-code--tests)
4. [Running Tests](#running-tests)

---

## Setup Instructions

1. **Install Go**:
    - Ensure you have Go installed. If not download it from [golang.org](https://golang.org/).
    - Verify the installation:
      ```bash
      go version
      ```

2. **Create the Project Directory**:
   ```bash
   mkdir go-algorithms && cd sorting-algorithms
   ```
   
3. **Initialize the Go Module**
   ```bash
   go mod init github.com/dimitrijjedich/go-algorithms
   ```

## Project Structure

Below the idea of the structure for this project:

```text
go-algorithms/
│
├── algorithms/
│   ├── bubble.go       # Bubble Sort implementation
│   └── ...             # Additional algorithms
│
├── algorithms_test/
│   ├── bubble_test.go  # Unit tests for Bubble Sort
│   └── ...             # Additional tests
│
├── main.go             # Main file for testing and demonstration
└── go.mod              # Go module file
```

## Writing Code & Tests

To write code and tests just create the files in the corresponding folders. And add the code according to the provided examples:

### Code

Using Bubble Sort as an example, create [algorithms/bubble.go](algorithms/bubble.go) and add the implementation. Repeat the same process for other sorting algorithms, by creating separate files under `algorithms/`

### Tests

Again using Bubble Sort as an example, create [algorithms_test/bubble_test.go](algorithms_test/bubble_test.go) for the test and add the implementation like provided by the example.

A simpler version of a tests would look like this:
```Go
package algorithms_test

import (
   // Import the algorithms package to access the BubbleSort function
   "github.com/dimitrijjedich/go-algorithms/algorithms"
   // Import the reflect package to compare slices for equality.
   "reflect"
   // Import the testing package to write and run unit tests.
   "testing"
)

// TestBubbleSort tests the BubbleSort function from the algorithms package.
func TestBubbleSort(t *testing.T) {
   // Define the input slice to be sorted.
   input := []int{5, 3, 8, 6, 2}
   // Define the expected result after sorting the input slice.
   expected := []int{2, 3, 5, 6, 8}

   // Call the BubbleSort function and store the result.
   result := algorithms.BubbleSort(input)

   // Use reflect.DeepEqual to check if the result matches the expected output.
   // If the result does not match, report a test failure with an error message.
   if !reflect.DeepEqual(result, expected) {
      t.Errorf("Expected %v, got %v", expected, result)
   }
}

```

#### Table-Driven Testing

To extend the basic tests Table-Driven testing can be used. You define all test cases in a slice of structs and iterate through them using t.Run.

The benefits are a centralized definition of test cases as combination with a description, input and expected result. Also, it makes it much easier to add test cases. 

An example implementation can be viewed in the before mentioned [bubble_test.go](algorithms_test/bubble_test.go) as it makes use of it.

## Running Tests

To run the tests execute the following command:
```bash
go test ./algorithms_test 
```

To make simple tests during the implementation process, you can always run the [main.go](main.go) file with a simpler dataset, that may be manipulated manually, as input. To run it execute the following command:
```bash
go run main.go
```