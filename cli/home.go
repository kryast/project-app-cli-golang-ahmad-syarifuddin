package cli

import (
	"fmt"
	"project-app-cli-golang-ahmad-syarifuddin/auth"
)

func Home() {
	defer func() {
		if r := recover(); r != nil {
			ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Home()
		}
	}()

	var input int
	fmt.Println("========================")
	fmt.Println("Selamat Datang Di UdinDB")
	fmt.Println("========================")
	fmt.Println("1. Login")
	fmt.Println("2. Daftar")
	fmt.Println("99. Exit")
	fmt.Print("Masukkan nomor menu: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		ClearScreen()
		fmt.Println("Silahkan Login")
		// fmt.Println(auth.DataAkun)
		if err := auth.Login(); err != nil {
			fmt.Println("Error:", err)
		}
	case 2:
		ClearScreen()
		fmt.Println("Silahkan Daftar Akun")
		akun, dataAkun, err := auth.DaftarAkun()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("---------------------")
			fmt.Println("Akun berhasil dibuat!")
			fmt.Printf("Username: %s\n", akun.Username)
			fmt.Println("---------------------")
			fmt.Println("Daftar Akun saat ini:")
			for _, a := range dataAkun {
				fmt.Printf("Username: %s\n", a.Username)
			}
			fmt.Println("---------------------")
		}

		fmt.Println("=============================")
		fmt.Println("Silahkan Kembali ke Menu Home")
		fmt.Println("=============================")
		fmt.Println("0. Kembali")
		fmt.Print("Masukkan nomor menu: ")
		fmt.Scan(&input)
		if input == 0 {
			Home()
		}
	case 99:
		ClearScreen()

	default:
		ClearScreen()
		fmt.Println("Halaman yang kamu pilih tidak tersedia.")
	}

}
