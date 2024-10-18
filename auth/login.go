package auth

import (
	"errors"
	"fmt"
)

func Login() error {
	var username string
	var password string

	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	for _, akun := range DataAkun {
		if akun.username == username {
			if akun.password == password {
				fmt.Println("Berhasil Login!")
				return nil
			}
			return errors.New("password salah")
		}
	}
	return errors.New("username salah")
}
