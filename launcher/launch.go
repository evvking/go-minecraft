package launcher

import (
	"fmt"
	"time"
)

type LaunchApp struct {
	Version   string
	VersionID int
}

func (mc LaunchApp) Run() { // TODO: Dodaj ładowanie zasobów osobno dla każdej wersji
	fmt.Println("\nLoading assets")
	time.Sleep(6 * time.Second)
	fmt.Println("Entering worlds directory")
	time.Sleep(3 * time.Second)
}
