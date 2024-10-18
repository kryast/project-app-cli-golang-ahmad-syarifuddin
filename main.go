package main

import (
	"fmt"
	"project-app-cli-golang-ahmad-syarifuddin/auth"
)

func main() {

	akun1 := auth.DaftarAkun("kryast", "contoh")
	akun2 := auth.DaftarAkun("kryast2", "contoh")
	akun3 := auth.DaftarAkun("kryast3", "contoh")
	akun4 := auth.DaftarAkun("kryast4", "contoh")

	auth.TampungAkun(akun1)
	auth.TampungAkun(akun2)
	auth.TampungAkun(akun3)
	auth.TampungAkun(akun4)

	fmt.Println(auth.DataAkun)

	err := auth.Login("tes", "contoh")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil, ")
	}

	auth.Login("kryast", "contoh")
}
