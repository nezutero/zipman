package main

import (
	"main.go/core"
)

func main() {
	core.Zip("./yourfilename.zip", []string{"./yourfilename_1.txt", "./yourfilename_2.txt"})
}
