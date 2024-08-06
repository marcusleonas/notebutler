package lib

import (
	"fmt"
	"os"
)

func Check() {
	if _, err := os.Stat(".notebutler"); os.IsNotExist(err) {
		fmt.Println("Config directory (.notebutler) not found. Notebutler not initialized. Run `notebutler init` to initialize.")
		os.Exit(1)
	}

	if _, err := os.Stat("notebutler.json"); os.IsNotExist(err) {
		fmt.Println("Config file (notebutler.json) not found. Notebutler not initialized. Run `notebutler init` to initialize.")
		os.Exit(1)
	}
}
