package pkg

type Component interface {
	Search(component Component)
	GetName() string
}
