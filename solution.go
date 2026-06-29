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
	m := Maalik{Name: "Karan", Work: "Aaram"}
	n := Nokar{Name: "Shivam", Work: "coding"}

	a := m.Do()
	fmt.Println(a)

	b := n.Do()
	fmt.Println(b)

	printDoerInfo(m)
	printDoerInfo(n)
}
