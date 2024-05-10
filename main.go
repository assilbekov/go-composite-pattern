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
	pc = pkg.PC{
		Name:        "AMD",
		Description: "AMD PC",
		Components: []pkg.Component{
			&motherBoard,
		},
	}
)

func main() {
	motherBoard.Search("Intel")
	motherBoard.Search("NVIDIA")
	motherBoard.Search("Asus")
	motherBoard.Search("AMD")
	motherBoard.Search("AMD")
	pc.Search("Intel")
	pc.Search("NVIDIA")
	pc.Search("Asus")
	pc.Search("AMD")
	pc.Search("AMD")
}
