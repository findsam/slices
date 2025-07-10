```go
package main

import "fmt"

func main() {
    bands := []string{"oasis", "oasis", "rush", "ac/dc"}
	filtered := slices.Filter(bands, func(name string) bool {
		return name == "oasis"
	})
    // returns [oasis oasis]
	fmt.Printf("Filtered: %v\n", filtered)

	count := slices.Reduce(filtered, 0, func(acc int, name string) int {
		if name == "oasis" {
			return acc + 1
		}
		return acc
	})
    // returns 2
	fmt.Printf("Count: %d\n", count)

	mapped := slices.Map(filtered, func(name string) string {
		return strings.ToUpper(string(name[0])) + name[1:]
	})
    // returns [Oasis Oasis]
	fmt.Printf("Mapped: %v\n", mapped)
}
```
