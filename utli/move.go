package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// Movematches
func Movematches(targetDir, toDir, find, exclude string, mode int, wet bool) {
	// Not provide? Same as dir
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
	workDir, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Organize todo
	matchName := regexp.MustCompile(find)
	excludeName := regexp.MustCompile(exclude)

	var toDo []transaction
	for _, item := range workDir {

		isDir := item.IsDir()
		fileName := item.Name()

		// Handle File
		if mode == FileAndFolder ||
			(!isDir && mode == FileOnly) ||
			(isDir && mode == FolderOnly) {

			// If match and not excluded
			if matchName.MatchString(fileName) &&
				((exclude == "") || (exclude != "" && !excludeName.MatchString(fileName))) {
				toDo = append(toDo, transaction{
					"MV",
					filepath.ToSlash(dirBase),
					filepath.ToSlash(fileName),
					filepath.ToSlash(toBase),
					filepath.ToSlash(fileName)})
			}
		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD][Mode%d] %s Find \"%s\" From \"%s\" -> To \"%s\"\n", mode, "move", find,
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
