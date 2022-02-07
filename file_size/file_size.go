package filesize

import "fmt"

/*
		x << y
	is x times 2 to the power of y

		x >> y
	is x divided by 2 to the power of y (fractional part discarded)
*/

func FileSizeMB() {
	// Convert to int64
	var MB1, MB5, MB20, MB800 int64
	MB20 = 1 << 20    // 1 MB
	MB5 = 5 << 20     // 5 MB
	MB20 = 20 << 20   // 20 MB
	MB800 = 800 << 20 // 800 MB
	fmt.Println("MB:", MB1, MB5, MB20, MB800)

	// Convert back to MB
	fmt.Println(MB20 >> 20) // 20
}
