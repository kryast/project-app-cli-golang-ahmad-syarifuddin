package main

import (
	"fmt"
	"project-app-cli-golang-ahmad-syarifuddin/auth"
	"project-app-cli-golang-ahmad-syarifuddin/cli"
)

func main() {
	auth.DaftarAkun()
	cli.ClearScreen()
	// auth.DaftarAkun()

	_, _, err := auth.DaftarAkun()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil Daftar")
	}
	cli.ClearScreen()

	fmt.Println(auth.DataAkun)
	err = auth.Login()
	if err != nil {
		fmt.Println("Error:", err)
	}
	cli.ClearScreen()
	// data.InputData()
	// data.InputData()
	// data.TampilkanSemuaData()
	// data.UpdateData()
	// data.TampilkanSemuaData()
	// data.DeleteData()
	// data.TampilkanSemuaData()

}
