package main

import "fmt"

type Doer interface {
	Do() string
	// Undo() string
}

// Oye Hoye ky scene he
type Maalik struct {
	Name string
	Work string
}

func (m Maalik) Do() string {
	return fmt.Sprintf("Maalik %s does %s", m.Name, m.Work)
}

// Oye Hoye ky scene he - 2
type Nokar struct {
	Name string
	Work string
}

func (n Nokar) Do() string {
	return fmt.Sprintf("Nokar %s does %s", n.Name, n.Work)
}

// Combo scene
func printDoerInfo(d Doer) {
	fmt.Printf("MSG: %s \n", d.Do())
}

func main() {
	var d Doer = Maalik{Name: "Karan", Work: "Management"}

	switch v := d.(type) {
	case Maalik:
		fmt.Println("It's a Maalik: ", v.Name)
	case Nokar:
		fmt.Println("It's a Nokar: ", v.Name)
	default:
		fmt.Println("unknown type")
	}

}
