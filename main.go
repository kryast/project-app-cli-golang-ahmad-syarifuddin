package main

import (
	"fmt"
	"project-app-cli-golang-ahmad-syarifuddin/auth"
)

func main() {
	auth.DaftarAkun()
	auth.DaftarAkun()
	auth.DaftarAkun()
	auth.DaftarAkun()

	fmt.Println(auth.DataAkun)

	// err := auth.Login("kryast", "contoh")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Berhasil, ")
	// }

	// auth.Login("kryast", "contoh")
}
