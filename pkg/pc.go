package pkg

import "fmt"

type PC struct {
	Name        string
	Description string
	Components  []Component
}

func (pc *PC) Search(name string) {
	if pc.Name == name {
		fmt.Println("PC found" + pc.Description + "\n")
	}

	for _, c := range pc.Components {
		c.Search(name)
	}
}

func (pc *PC) GetName() string {
	return pc.Name
}
