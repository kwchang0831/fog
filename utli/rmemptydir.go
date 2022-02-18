package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// RmEmptyDir
func RmEmptyDir(targetDir string, wet bool) {
	// Get absolute path
	dirBase, _ := filepath.Abs(targetDir)

	// Check targetDir
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Start checking files in the targetDir
	parentDir, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Organize todo
	var toDo []transaction
	for _, parentItem := range parentDir {
		if parentItem.IsDir() {
			childDir, err := ioutil.ReadDir(path.Join(targetDir, parentItem.Name()))
			if err != nil {
				fmt.Printf("[ERR] %s\n", err)
				os.Exit(1)
			}
			if len(childDir) == 0 {
				toDo = append(toDo, transaction{"RMDIR",
					filepath.ToSlash(dirBase),
					filepath.ToSlash(parentItem.Name()),
					"",
					""})
			}
		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD] %s In \"%s\"\n", "rmemptydir", filepath.Join(filepath.ToSlash(dirBase)))
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
