package main

import (
	"fmt"

	"minecraft/launcher"
)

func main() {
	fmt.Println("\nGo! Minecraft Launcher")
	fmt.Println("\nSelected version: 1.0 (Inital Release)")

	var launcherApp = launcher.LaunchApp{
		Version:   "1.0",
		VersionID: 1,
	}
}
