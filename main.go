package main

import (
	"github.com/lancer2672/DandelionServer_Go/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	_, _ = config, err
}
