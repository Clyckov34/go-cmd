package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sudo", "apt", "update")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Run()
}

//https://evilinside.ru/kak-v-go-lang-zapustit-konsolnuyu-komandu-v-interaktivnom-rezhime/#
