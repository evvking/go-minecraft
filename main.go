package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"minecraft/libraries/commandline"
	"minecraft/libraries/launcher"
	"minecraft/libraries/world"
)

func main() {
	fmt.Println("\nGo! Minecraft Launcher")
	fmt.Println("\nSelected version: 1.0 (Inital Release)")

	var launcherApp launcher.LaunchApp = launcher.LaunchApp{
		VersionID: 1,
	}

	var commandLine commandline.CommandLine
	var cmd string

LaunchCommandLine:
	for {
		cmd = commandLine.Input()
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
		case "versionselect":
			fmt.Print("Enter Version ID: ")
			fmt.Scanln(&launcherApp.VersionID)
		}
	}

	fmt.Println("\nWorlds Manager")

	var world = world.World{}
WorldsManager:
	for {
		cmd = commandLine.Input()

		switch cmd {
		case "worldcreate":
			fmt.Print("Enter world name: ")
			fmt.Scanln(&world.Name)

			world.Make()
			world.Print()
			break WorldsManager
		}
	}
}
