package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// range
// map
func main() {
	for i, v := range []int{1, 2, 3, 4, 5} {
		println(i, v)
	}

	fmt.Println(m["Bell Labs"])
	delete(m, "Bell Labs")
	e, ok := m["Bell labs"]
	fmt.Println(e, ok)

}
