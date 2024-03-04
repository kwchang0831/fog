package utli

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Folderin
func Folderin(targetDir string, toDir string, wet bool) {
	// Not provide? Same as dir
	if toDir == "" {
		toDir = targetDir
	}

	// Get absolute path
	dirBase, _ := filepath.Abs(targetDir)
	toBase, _ := filepath.Abs(toDir)

	// Check targetDir
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Start checking files in the dir
	workDir, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Organize todo
	var toDo []transaction
	for _, item := range workDir {
		if !item.IsDir() {
			fileName := item.Name()
			oldPath := fileName
			fileExt := filepath.Ext(fileName)
			fileNameOnly := fileName[0 : len(fileName)-len(fileExt)]
			newPath := path.Join(fileNameOnly, fileName)
			toDo = append(toDo, transaction{
				"MKDIR",
				filepath.ToSlash(toBase),
				filepath.ToSlash(fileNameOnly),
				"",
				""})
			toDo = append(toDo, transaction{
				"MV",
				filepath.ToSlash(dirBase),
				filepath.ToSlash(oldPath),
				filepath.ToSlash(toBase),
				filepath.ToSlash(newPath)})
		}
	}
	// Start doing work
	startMsg := fmt.Sprintf("[CMD] %s From \"%s\" -> To \"%s\"\n", "folderin",
		filepath.Join(filepath.ToSlash(dirBase)), filepath.Join(filepath.ToSlash(toBase)))
	endMsg := ""
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
	fmt.Print(endMsg)
}
