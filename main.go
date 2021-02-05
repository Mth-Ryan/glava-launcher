package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"os/user"
	"strings"
)

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {

	//Kill Any Glava Session
	run("killall", "glava")

	//User Configuration
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configPath := user.HomeDir + "/.config/glava/rc.glsl"

	//Open Config file
	imp, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal("Falied to read the rc.glsl: ", err)
	}

	//Search and Replace Resolution
	lines := strings.Split(string(imp), "\n")
	replaceLine := "#request setgeometry"
	width, height := getResolution()

	for i, line := range lines {
		if strings.Contains(line, replaceLine) {
			lines[i] = fmt.Sprintf("%s -15 0 %d %d",
				replaceLine, (width + 30), height)
		}
	}

	//Write Config file
	out := strings.Join(lines, "\n")
	err = ioutil.WriteFile(configPath, []byte(out), 0644)

	if err != nil {
		log.Fatal("Falied write the rc.glsl: ", err)
	}

	//Run Glava
	err = run("glava", "--desktop")

	if err != nil {
		log.Fatal("Falied to run Glava: ", err)
	}
}
