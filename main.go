package main

import (
	"project-app-cli-golang-ahmad-syarifuddin/data"
)

func main() {
	// auth.DaftarAkun()
	// auth.DaftarAkun()

	// fmt.Println(auth.DataAkun)

	// err := auth.Login()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Berhasil Login")
	// }

	data.InputData()
	data.InputData()
	// data.TampilkanSemuaData()
	data.UpdateData()
	// data.TampilkanSemuaData()
	data.DeleteData()
	data.TampilkanSemuaData()

}
