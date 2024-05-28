## `byteseq` Go Module
This is a simple Go Module that provides a means to iterate over `byte` values,
ensuring a given value is returned only once. The ability to specify a slice 
of specific byte values **_NOT_** be returned is also possible.

The main use case for this was randomly iterating over host byte values
in IPv4 addresses that represent a traditional "Class C" network block. In
this particular case, returning the `.0` and `.255` values was not desirable,
as these aren't valid host byte values in such addresses.

---

### How to use
The following is a very simply example that iterates over the full 0-255 range.
It is also available to experiment with on [Go Playground](https://go.dev/play/p/apb6kznVmZ9).
```go
package main

import (
	"fmt"

	"github.com/owenjklan/byteseq"
)

func main() {
	bseq := byteseq.NewRandomSeq(nil)

	for bseq.HasMore() {
		value, err := bseq.NextValue()
		fmt.Printf("Byte value: %02x\n", value)

		if err != nil {
			fmt.Println("Error obtaining value from sequence:", err)
		}
	}
}
```

We can extend the above example to demonstrate what happens if we attempt to
obtain values from an exhausted sequence. This is also available on [Go Playground](https://go.dev/play/p/70SSUy3IZJm).

```go
package main

import (
	"fmt"

	"github.com/owenjklan/byteseq"
)

func main() {
	bseq := byteseq.NewRandomSeq(nil)

	for bseq.HasMore() {
		value, err := bseq.NextValue()
		fmt.Printf("Byte value: %02x\n", value)

		if err != nil {
			fmt.Println("Error obtaining value from sequence:", err)
		}
	}

	// We should now receive an error because we've exhausted the sequence
	fmt.Println("Sequence has more bytes:", bseq.HasMore())

	value, err := bseq.NextValue()
	if err != nil {
		fmt.Println("Error obtaining value from sequence:", err)
		return
	}

	// The following should not be reached
	fmt.Printf("The very last value we got: %02x", value)
}
```

The final example demonstrates how to restrict the byte values that will be returned.
As with previous examples, a [Go Playground]() link is available.
```go
package main

import (
	"fmt"

	"github.com/owenjklan/byteseq"
)

func main() {
	var consumedBytes []byte

	// Create a slice of byte values
	for i := 0; i < 250; i++ {
		consumedBytes = append(consumedBytes, byte(i))
	}

	bseq := byteseq.NewRandomSeq(consumedBytes)

	for bseq.HasMore() {
		value, err := bseq.NextValue()
		fmt.Printf("Byte value: %02x\n", value)

		if err != nil {
			fmt.Println("Error obtaining value from sequence:", err)
		}
	}
}
```

For this last example, a reduced range of byte values is returned:
```text
Byte value: fb
Byte value: fe
Byte value: fc
Byte value: fd
Byte value: fa
Byte value: ff
```

Observe that the smallest value returned is 250 and the largest is 256 and
that the returned order is indeed random.

--- 

This package was also my first attempt at publishing (and then using) a Go
Module on Github. There is probably room for improvement.

