package main

import (
	"flag"
	"os/exec"
)

func main() {
	hideFlag := flag.Bool("hide", false, "hide hidden files")
	flag.Parse()

	showHidddenFilesCmd := exec.Command("defaults", "write", "com.apple.Finder", "AppleShowAllFiles", "true")
	hideHidddenFilesCmd := exec.Command("defaults", "write", "com.apple.Finder", "AppleShowAllFiles", "false")

	if *hideFlag {
		hideHidddenFilesCmd.Run()
	} else {
		showHidddenFilesCmd.Run()
	}

	killAllFinderCmd := exec.Command("killall", "Finder")
	killAllFinderCmd.Run()
}
