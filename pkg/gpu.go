package pkg

import "fmt"

type GPU struct {
	Name        string
	Description string
}

func (g *GPU) Search(name string) {
	if g.Name == name {
		fmt.Printf("GPU [%s] found, %s\n", name, g.Description)
		return
	}
}

func (g *GPU) GetName() string {
	return g.Name
}
