package main

import (
	"fmt"
	"project-app-cli-golang-ahmad-syarifuddin/auth"
)

func main() {
	auth.DaftarAkun()
	auth.DaftarAkun()

	fmt.Println(auth.DataAkun)

	err := auth.Login()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil Login")
	}

}
