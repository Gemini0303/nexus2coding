package main

import (
	"nexus2coding/src/nexus"
)

func main() {
	n := nexus.New("http://43.254.44.133:58000", "admin", "admin123")
	n.ChooseRepository()

}
