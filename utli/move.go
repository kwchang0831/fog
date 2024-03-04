package utli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Movematches optimizes the file moving operation
func Movematches(targetDir, toDir, find, exclude string, mode int, wet bool) {
	// Set toDir to targetDir if not provided
	if toDir == "" {
		toDir = targetDir
	}

	// Get absolute paths
	dirBase, _ := filepath.Abs(targetDir)
	toBase, _ := filepath.Abs(toDir)

	// Check if targetDir exists
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Read directory
	workDir, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Prepare paths for logging
	dirBaseSlash := filepath.ToSlash(dirBase)
	toBaseSlash := filepath.ToSlash(toBase)
	startMsg := fmt.Sprintf("[CMD][Mode%d] move Find \"%s\" From \"%s\" -> To \"%s\"\n", mode, find, dirBaseSlash, toBaseSlash)
	endMsg := ""

	// Prepare regular expressions
	matchName := regexp.MustCompile(find)
	var excludeName *regexp.Regexp
	if exclude != "" {
		excludeName = regexp.MustCompile(exclude)
	}

	// Preallocate capacity for toDo slice
	toDo := make([]transaction, 0, len(workDir))

	// Iterate through files and folders
	for _, item := range workDir {
		fileName := item.Name()
		isDir := item.IsDir()

		// Combine conditions for handling files and folders
		if (mode == FileAndFolder) || (mode == FileOnly && !isDir) || (mode == FolderOnly && isDir) {
			if matchName.MatchString(fileName) && (excludeName == nil || !excludeName.MatchString(fileName)) {
				toDo = append(toDo, transaction{"MV", dirBaseSlash, filepath.ToSlash(fileName), toBaseSlash, filepath.ToSlash(fileName)})
			}
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

	// Execute transactions
	for _, t := range toDo {
		fmt.Print(t.commit(wet))
	}

	// Log end message
	fmt.Print(endMsg)
}
