package main

import "composite/pkg"

var (
	cpu = pkg.CPU{
		Name:        "Intel",
		Description: "Intel CPU",
	}
	gpu = pkg.GPU{
		Name:        "NVIDIA",
		Description: "NVIDIA GPU",
	}
	motherBoard = pkg.MotherBoard{
		Name:        "Asus",
		Description: "Asus MotherBoard",
		Components: []pkg.Component{
			&cpu,
			&gpu,
		},
	}
)

func main() {
	motherBoard.Search("Intel")
	motherBoard.Search("NVIDIA")
	motherBoard.Search("Asus")
	motherBoard.Search("AMD")
	motherBoard.Search("AMD")
}
