package main

import (
	"fmt"
	"project-app-cli-golang-ahmad-syarifuddin/auth"
)

func main() {
	// auth.DaftarAkun()
	// auth.DaftarAkun()

	_, _, err := auth.DaftarAkun()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil Daftar")
	}

	fmt.Println(auth.DataAkun)
	err = auth.Login()
	if err != nil {
		fmt.Println("Error:", err)
	}
	// data.InputData()
	// data.InputData()
	// data.TampilkanSemuaData()
	// data.UpdateData()
	// data.TampilkanSemuaData()
	// data.DeleteData()
	// data.TampilkanSemuaData()

}
