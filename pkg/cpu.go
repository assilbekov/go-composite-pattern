package pkg

import "fmt"

type CPU struct {
	Name        string
	Description string
}

func (c *CPU) Search(name string) {
	if c.Name == name {
		fmt.Printf("CPU [%s] found, %s\n", name, c.Description)
		return
	}
}

func (c *CPU) GetName() string {
	return c.Name
}
