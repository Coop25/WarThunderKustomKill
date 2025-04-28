package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	manager "github.com/Coop25/WarThunderKustomKill/managers"
)

var configLine = "  testLocalization:b=yes"

var targets = []manager.CSVRowTarget{
	{Key: "exp_reasons/kill_gm", HumanName: "Target destroyed", OriginalVal: "Target destroyed"},
	{Key: "exp_reasons/kill", HumanName: "Aircraft destroyed", OriginalVal: "Aircraft destroyed"},
	{Key: "exp_reasons/capture", HumanName: "Zone captured", OriginalVal: "Zone captured"},
	{Key: "exp_reasons/hit", HumanName: "Hit", OriginalVal: "Hit"},
	{Key: "exp_reasons/critical_hit", HumanName: "Critical hit", OriginalVal: "Critical hit"},
	{Key: "exp_reasons/ineffective_hit", HumanName: "Target undamaged", OriginalVal: "Target undamaged"},
}

func main() {
	wtPath, err := manager.FindWarThunderPath()
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(wtPath, "config.blk")
	menuPath := filepath.Join(wtPath, `lang\menu.csv`)

	// Normal run
	cm, err := manager.NewConfigManager(configPath)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)
	
	if cm.HasInsertedLine(configLine) {
		fmt.Print("Do you want to undo previous changes? (y/n): ")
		undoInput, _ := reader.ReadString('\n')
		undoInput = strings.TrimSpace(strings.ToLower(undoInput))

		if undoInput == "y" {
			err = cm.RemoveInsertedLine(configLine)
			if err != nil {
				panic(err)
			}
			err = cm.Save()
			if err != nil {
				panic(err)
			}
			fmt.Println("Removed inserted line from config.blk.")

			// Undo menu.csv
			csvm, err := manager.NewCSVManager(menuPath)
			if err != nil {
				panic(err)
			}
			err = csvm.ResetSecondColumnForMatches(targets)
			if err != nil {
				panic(err)
			}
			err = csvm.Save()
			if err != nil {
				panic(err)
			}
			fmt.Println("Reset matching rows in menu.csv.")

			fmt.Println("Undo complete. Exiting.")
			return
		}
		fmt.Println("config.blk line already exists. Skipping insert.")
	} else {
		err = cm.InsertAfterSection("debug{", configLine)
		if err != nil {
			panic(err)
		}
		err = cm.Save()
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted new line into config.blk and saved.")
	}

	fmt.Println("Please Launch Warthunder and confirm the top Left says the below text")
	fmt.Println(`"Custom localization enabled"`)
	fmt.Println("If so fully close the game and Press ENTER to continue...")
	reader.ReadBytes('\n')

	csvm, err := manager.NewCSVManager(menuPath)
	if err != nil {
		panic(err)
	}
	err = csvm.UpdateSecondColumnInteractive(targets)
	if err != nil {
		panic(err)
	}
	err = csvm.Save()
	if err != nil {
		panic(err)
	}
	fmt.Println("menu.csv updated and saved.")
}
