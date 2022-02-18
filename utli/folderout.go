package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Folderout
func Folderout(targetDir string, toDir string, wet bool) {
	// Set toDir to targetDir if not provided
	if toDir == "" {
		toDir = targetDir
	}

	// Get absolute path
	dirBase, _ := filepath.Abs(targetDir)
	toBase, _ := filepath.Abs(toDir)

	// Check dir
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Start checking files in the dir
	parentDir, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Organize todo
	var toDo []transaction
	for _, parentItem := range parentDir {
		if parentItem.IsDir() {
			todoPath := path.Join(targetDir, parentItem.Name())
			childDir, err := ioutil.ReadDir(todoPath)
			if err != nil {
				fmt.Printf("[ERR] %s\n", err)
				os.Exit(1)
			}
			for _, childItem := range childDir {
				oldPath := path.Join(parentItem.Name(), childItem.Name())
				newPath := childItem.Name()
				toDo = append(toDo, transaction{
					"MV",
					filepath.ToSlash(dirBase),
					filepath.ToSlash(oldPath),
					filepath.ToSlash(toBase),
					filepath.ToSlash(newPath)})
			}
			toDo = append(toDo, transaction{"RMDIR",
				filepath.ToSlash(dirBase),
				filepath.ToSlash(parentItem.Name()),
				"",
				""})
		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD] %s From \"%s\" -> To \"%s\"\n", "folderout",
		filepath.Join(filepath.ToSlash(dirBase)), filepath.Join(filepath.ToSlash(toBase)))
	endMsg := ""
	if wet {
		log.info("#" + startMsg)
		startMsg = "[WET]" + startMsg
		endMsg = fmt.Sprintf("[DONE] saved @ \"%s\"\n", filepath.Join(log.getLogDirName(), log.getFileName()))
	} else {
		startMsg = "[DRY]" + startMsg
	}

	fmt.Printf(startMsg)
	// Commit transactions
	for _, i := range toDo {
		fmt.Print(i.commit(wet))
	}
	fmt.Printf(endMsg)
}
