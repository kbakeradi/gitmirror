package main

import (
	"fmt"
	"strings"
)

type gitData struct {
	name string
	src  string
	dest string
}

func writeCmd(cmd string, args []string) {
	fmt.Printf("%s %s\n", cmd, strings.Join(args, " "))
}
func main() {
	fmt.Println("## Git Mirror Command")
	fmt.Println("##########################")
	fmt.Println(" ")
	// Settings
	gitList := []gitData{
		{
			name: "vlab-app-magnetics",
			src:  "git@ssh.dev.azure.com:v3/adi-vlab/vlab-app-magnetics/vlab-app-magnetics",
			dest: "ssh://git@bitbucket.analog.com:7999/vlabarchive/vlab-app-magnetics.git",
		},
		{
			name: "vlab-app-service-magnetics",
			src:  "git@ssh.dev.azure.com:v3/adi-vlab/vlab-app-service-magnetics/vlab-app-service-magnetics",
			dest: "ssh://git@bitbucket.analog.com:7999/vlabarchive/vlab-app-service-magnetics.git",
		},
	}

	const GIT_LOCAL_PATH = "/Users/kbaker/Code/Analog/vlab/apps/archive"

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("########## COPY AND RUN THE BELOW CODE ##########")
	fmt.Println("")
	fmt.Println("")

	for _, git := range gitList {
		workingDir := GIT_LOCAL_PATH + "/" + git.name
		writeCmd("git", []string{"clone", git.src, workingDir})
		writeCmd("cd", []string{workingDir})
		writeCmd("git", []string{"remote", "add", "bitbucket", git.dest})
		writeCmd("git", []string{"branch", "-M", "main"})
		writeCmd("git", []string{"push", "bitbucket", "main"})
		writeCmd("cd", []string{"-"})
		fmt.Println("")
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("########## COPY AND RUN THE ABOVE CODE ##########")
}
