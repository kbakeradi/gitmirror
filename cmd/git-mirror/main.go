package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Git Mirror Command")

	const GIT_DEST_PATH string = "/Users/kbaker/Code/Analog/vlab/apps/archive"

	dirList := []string{
		"vlab-app-magnetics",
		"vlab-app-service-magnetics",
	}

	cmd := "echo"
	flags := ""

	for _, dir := range dirList {
		cmdfull := exec.Command(cmd, flags, dir)
		stdout, err := cmdfull.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("------------------------------------------")
		fmt.Printf("%s %s", cmd, dir)
		fmt.Println("")
		fmt.Print(string(stdout))

	}

}
