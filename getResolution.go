package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func getResolution() (int64, int64) {
	cmd := "xdpyinfo | awk '/dimensions/{print $2}'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		log.Fatal("Falied to get Resolution: ", err)
	}

	resolution := strings.Trim(string(out), "\n")
	dimensions := strings.Split(resolution, "x")

	width, err1 := strconv.ParseInt(dimensions[0], 0, 64)
	height, err2 := strconv.ParseInt(dimensions[1], 0, 64)

	if err1 != nil || err2 != nil {
		log.Fatal("Falied to convert dimensions: ", err1, err2)
	}

	return width, height
}
