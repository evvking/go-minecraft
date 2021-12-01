package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"minecraft/launcher"
	commandline "minecraft/libraries"
)

func main() {
	fmt.Println("\nGo! Minecraft Launcher")
	fmt.Println("\nSelected version: 1.0 (Inital Release)")

	var launcherApp launcher.LaunchApp = launcher.LaunchApp{
		Version:   "1.0",
		VersionID: 1,
	}

	var commandLine commandline.CommandLine

LaunchCommandLine:
	for {
		var cmd string = commandLine.Input()
		// FIXME: Rozbijanie na CommandSplit nie działa (błąd przy [case "versionlist"]), błąd tablic

		switch cmd {
		case "launch":
			launcherApp.Run()
			break LaunchCommandLine
		case "versionlist": // TODO: Rozbij "versionlist" na zagnieżdzone SWITCH po naprawieniu błędu z tablicami
			file, err := os.Open("assets/versionlist.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
