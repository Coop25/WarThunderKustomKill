package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Coop25/WarThunderKustomKill/manager"
)

func main() {
	wtPath, err := manager.FindWarThunderPath()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to undo previous changes? (y/n): ")
	undoInput, _ := reader.ReadString('\n')
	undoInput = strings.TrimSpace(strings.ToLower(undoInput))

	configPath := filepath.Join(wtPath, "config.blk")
	menuPath := filepath.Join(wtPath, "menu.csv")

	if undoInput == "y" {
		// Undo config.blk
		cm, err := manager.NewConfigManager(configPath)
		if err != nil {
			panic(err)
		}
		err = cm.RemoveInsertedLine("inserted_by_tool = true")
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
		targets := []manager.CSVRowTarget{
			{Key: "aircraft1", HumanName: "Spitfire Mk I", OriginalVal: "Locked"},
			{Key: "aircraft2", HumanName: "BF-109", OriginalVal: "Locked"},
			{Key: "tank1", HumanName: "Tiger I Tank", OriginalVal: "Locked"},
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

	// Normal run
	cm, err := manager.NewConfigManager(configPath)
	if err != nil {
		panic(err)
	}
	err = cm.InsertAfterSection("SECTION_NAME", "inserted_by_tool = true")
	if err != nil {
		panic(err)
	}
	err = cm.Save()
	if err != nil {
		panic(err)
	}
	fmt.Println("config.blk updated and saved.")

	fmt.Println("Press ENTER to continue...")
	reader.ReadBytes('\n')

	csvm, err := manager.NewCSVManager(menuPath)
	if err != nil {
		panic(err)
	}
	targets := []manager.CSVRowTarget{
		{Key: "aircraft1", HumanName: "Spitfire Mk I", OriginalVal: "Locked"},
		{Key: "aircraft2", HumanName: "BF-109", OriginalVal: "Locked"},
		{Key: "tank1", HumanName: "Tiger I Tank", OriginalVal: "Locked"},
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
