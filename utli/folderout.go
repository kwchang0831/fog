package utli

import (
	"fmt"
	"os"
	"path/filepath"
)

// Folderout optimizes the directory organization operation
func Folderout(targetDir string, toDir string, wet bool) {
	// Set toDir to targetDir if not provided
	if toDir == "" {
		toDir = targetDir
	}

	// Get absolute path
	dirBase, _ := filepath.Abs(targetDir)
	toBase, _ := filepath.Abs(toDir)

	// Check if targetDir exists
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Start checking files in the dir
	parentDir, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Prepare paths for logging
	dirBaseSlash := filepath.ToSlash(dirBase)
	toBaseSlash := filepath.ToSlash(toBase)
	startMsg := fmt.Sprintf("[CMD] folderout From \"%s\" -> To \"%s\"\n", dirBaseSlash, toBaseSlash)
	endMsg := ""

	// Organize todo
	var toDo []transaction
	for _, parentItem := range parentDir {
		if parentItem.IsDir() {
			todoPath := filepath.Join(targetDir, parentItem.Name())
			childDir, err := os.ReadDir(todoPath)
			if err != nil {
				fmt.Printf("[ERR] %s\n", err)
				os.Exit(1)
			}
			for _, childItem := range childDir {
				oldPath := filepath.Join(parentItem.Name(), childItem.Name())
				newPath := childItem.Name()
				toDo = append(toDo, transaction{
					"MV",
					dirBaseSlash,
					filepath.ToSlash(oldPath),
					toBaseSlash,
					filepath.ToSlash(newPath)})
			}
			toDo = append(toDo, transaction{"RMDIR",
				dirBaseSlash,
				filepath.ToSlash(parentItem.Name()),
				"",
				""})
		}
	}

	// Log start message
	if wet {
		log.info("#" + startMsg)
		startMsg = "[WET]" + startMsg
		endMsg = fmt.Sprintf("[DONE] saved @ \"%s\"\n", filepath.Join(log.getLogDirName(), log.getFileName()))
	} else {
		startMsg = "[DRY]" + startMsg
	}

	fmt.Print(startMsg)

	// Commit transactions
	for _, i := range toDo {
		fmt.Print(i.commit(wet))
	}

	// Log end message
	fmt.Print(endMsg)
}
