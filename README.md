# Aten Serializer Package

## Description

The Go Serializer Package provides a flexible and generic way to serialize and deserialize data, supporting various types, including integers, floats, booleans, slices, maps, and more. It aims to simplify data conversion tasks and enhance the interoperability of data between different systems.

## Goal of the Package

The goal of this package is to offer a convenient and extensible solution for converting data to and from string representations. The package includes a range of default supported types and allows users to easily extend the serialization and deserialization capabilities for custom types. Whether you're working with simple values or complex data structures, this package aims to streamline the process of data conversion.

## How to Use

### Serialization

To serialize a string representation of data, use the `Serialize` function provided by the package. This function attempts to convert the input string to various predefined types, returning the converted value if successful. If no suitable conversion is found, the original string is returned.

Example:

```go
package main

import (
	"fmt"
	"github.com/atenteccompany/aten-serializer"
)

func main() {
	inputString := "42"
	result, err := serializer.Serialize(inputString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Serialized Value:", result)
}
```

### Deserialization

To deserialize a string representation back into its original type, use the `Deserialize` function. The function analyzes the type of the input data and performs the appropriate deserialization. It supports basic types like integers, floats, booleans, as well as slices, maps, and structs.

Example:

```go
package main

import (
	"fmt"
	"github.com/atenteccompany/aten-serializer"
)

func main() {
	inputData := 42
	result, err := serializer.Deserialize(inputData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deserialized Value:", result)
}
```

### Generic Serialization

For more specialized use cases, the package also includes generic serialization using the `SerializeT` function. This function leverages Go's type system to provide a type-safe way to deserialize data into a specified type.
Example:

```go
package main

import (
	"fmt"
	"github.com/atenteccompany/aten-serializer"
)

type CustomStruct struct {
	Value int
}

func main() {
	inputString := `{"Value": 42}`
	result, err := serializer.SerializeT[CustomStruct](inputString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Deserialized Value: %+v\n", *result)
}
```

