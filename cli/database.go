package cli

import (
	"errors"
	"fmt"
	"project-app-cli-golang-ahmad-syarifuddin/data"
)

func Database() {
	defer func() {
		if r := recover(); r != nil {
			ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Database()
		}
	}()

	var input int
	fmt.Println("========================")
	fmt.Println("Selamat Datang Di UdinDB")
	fmt.Println("========================")
	fmt.Println("1. Input Data Pegawai")
	fmt.Println("2. Input Data Konsumen")
	fmt.Println("3. Update Data Pegawai")
	fmt.Println("4. Update Data Konsumen")
	fmt.Println("5. Delete Data Pegawai")
	fmt.Println("6. Delete Data Konsumen")
	fmt.Println("7. Tampilkan Semua Data")
	fmt.Println("0. Log Out")
	fmt.Println("99. Exit")
	fmt.Print("Masukkan nomor menu: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		ClearScreen()
		fmt.Println("Input Data Pegawai")
		data.InputDataPegawai()
		BackDatabase()

	case 2:
		ClearScreen()
		fmt.Println("Input Data Konsumen")
		data.InputDataKonsumen()
		BackDatabase()

	case 3:
		ClearScreen()
		fmt.Println("Update Data Pegawai")
		var pegawai data.DataPegawai
		pegawai.UpdateData()
		BackDatabase()

	case 4:
		ClearScreen()
		fmt.Println("Update Data Konsumen")
		var konsumen data.DataKonsumen
		konsumen.UpdateData()
		BackDatabase()

	case 5:
		ClearScreen()
		fmt.Println("Delete Data")
		var pegawai data.DataPegawai
		pegawai.DeleteData()
		BackDatabase()

	case 6:
		ClearScreen()
		fmt.Println("Delete Data Konsumen")
		var konsumen data.DataKonsumen
		konsumen.DeleteData()
		BackDatabase()

	case 7:
		ClearScreen()
		fmt.Println("Tampilkan Semua Data")
		fmt.Println("====================")
		var pegawai data.DataPegawai
		pegawai.TampilkanSemuaData()
		var konsumen data.DataKonsumen
		konsumen.TampilkanSemuaData()
		BackDatabase()

	case 0:
		ClearScreen()
		Home()

	case 99:
		ClearScreen()

	default:
		err := errors.New("")
		Panik(err)

	}
}
