package cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func BackHome() {
	var input int
	fmt.Println("=============================")
	fmt.Println("Silahkan Kembali ke Menu Home")
	fmt.Println("=============================")
	fmt.Println("0. Kembali")
	fmt.Print("Masukkan nomor menu: ")
	fmt.Scan(&input)
	ClearScreen()
	if input == 0 {
		Home()
	} else {
		err := errors.New("")
		Panik(err)
	}
}
