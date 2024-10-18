package cli

import (
	"errors"
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
	fmt.Println("3. Lupa Username")
	fmt.Println("4. Lupa Password")
	fmt.Println("99. Exit")
	fmt.Print("Masukkan nomor menu: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		ClearScreen()
		fmt.Println("Silahkan Login")
		// fmt.Println(auth.DataAkun)
		if err := auth.Login(); err != nil {
			Panik(err)
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

		BackHome()
	case 3:
		ClearScreen()
		fmt.Println("---------------------")
		fmt.Println("Daftar Akun saat ini:")
		for _, a := range auth.DataAkun {
			fmt.Printf("Username: %s\n", a.Username)
		}
		fmt.Println("---------------------")
		BackHome()

	case 4:
		ClearScreen()

		var input string
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&input)

		if input == "" {
			fmt.Println("Username tidak boleh kosong.")
			BackHome()
		}

		for _, akun := range auth.DataAkun {
			if akun.Username == input {
				fmt.Printf("Username: %s\nPassword: %s\n", akun.Username, akun.Password)
				BackHome()
			} else {
				fmt.Println("Username tidak ditemukan. Silakan coba lagi.")
				BackHome()
			}
		}

	case 99:
		ClearScreen()

	default:
		err := errors.New("")
		Panik(err)
	}

}
