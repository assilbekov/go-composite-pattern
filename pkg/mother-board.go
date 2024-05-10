package pkg

import "fmt"

type MotherBoard struct {
	Name        string
	Description string
	Components  []Component
}

func (m *MotherBoard) Search(name string) {
	if m.Name == name {
		fmt.Printf("MotherBoard [%s] found, %s\n", name, m.Description)
	}

	for _, c := range m.Components {
		c.Search(name)
	}
}

func (m *MotherBoard) GetName() string {
	return m.Name
}
