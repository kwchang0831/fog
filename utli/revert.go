package utli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Revert
func Revert(logPath string, wet bool) {
	f, _ := os.Open(logPath)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var toDo []transaction
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		a := actiontaken{}
		json.Unmarshal(scanner.Bytes(), &a)

		switch a.Action {
		case "MV":
			toDo = append(toDo, transaction{
				"MV",
				"",
				filepath.ToSlash(a.To),
				"",
				filepath.ToSlash(a.From)})
		case "RMDIR":
			toDo = append(toDo, transaction{
				"MKDIR",
				"",
				filepath.ToSlash(a.From),
				"",
				""})
		case "MKDIR":
			toDo = append(toDo, transaction{
				"RMDIR",
				"",
				filepath.ToSlash(a.From),
				"",
				""})
		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD] %s \"%s\"\n", "revert", filepath.Join(logPath))
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
	for i := len(toDo) - 1; i >= 0; i-- {
		msg := toDo[i].commit(wet)
		fmt.Print(msg)
	}
	fmt.Print(endMsg)
}
