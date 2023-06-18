package room

import "fmt"

type Cat struct {
	Color string
	Age   uint8
	Name  string
}

func (cat *Cat) Meow() {
	fmt.Println("Meoooooow")
}

func (cat *Cat) Rename(newName string) {
	cat.Name = newName
}

func (cat Cat) RenameV2(newName string) {
	cat.Name = newName
}

// PrintDetails use a capital letter on the first letter of the function
// to export the function
func PrintDetails(roomNumber, size, nights int) {
	fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}

func Zeroval(ival int) {
	ival = 0
}

func Zeroptr(iptr *int) {
	*iptr = 0
}
