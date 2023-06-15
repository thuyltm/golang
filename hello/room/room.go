package room

import "fmt"

// use a capital letter on the first letter of the function
// to export the function
func PrintDetails(roomNumber, size, nights int) {
  fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}