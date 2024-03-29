package utli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// ReplaceName
func ReplaceName(targetDir, toDir, find, replace, exclude string, mode int, wet bool) {
	// Set toDir to targetDir if not provided
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

	// Start checking files in the targetDir
	parentDir, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Organize todo
	matchName := regexp.MustCompile(find)
	excludeName := regexp.MustCompile(exclude)

	var toDo []transaction
	for _, item := range parentDir {

		isDir := item.IsDir()
		oldFileName := item.Name()

		// Handle File
		if mode == FileAndFolder ||
			(!isDir && mode == FileOnly) ||
			(isDir && mode == FolderOnly) {

			// If match and not excluded
			if matchName.MatchString(oldFileName) &&
				((exclude == "") || (exclude != "" && !excludeName.MatchString(oldFileName))) {
				newFileName := matchName.ReplaceAllString(oldFileName, replace)
				toDo = append(toDo, transaction{
					"MV",
					filepath.ToSlash(dirBase),
					filepath.ToSlash(oldFileName),
					filepath.ToSlash(toBase),
					filepath.ToSlash(newFileName)})
			}

		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD][Mode%d] %s Find \"%s\" Replace with \"%s\" From \"%s\" -> toDir \"%s\"\n", mode, "replacename", find, replace,
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
