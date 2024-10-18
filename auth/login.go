package auth

import (
	"errors"
	"fmt"
)

func Login() error {
	var username string
	var password string

	fmt.Print("Masukkan Username :")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password :")
	fmt.Scan(&password)

	for _, akun := range DataAkun {
		if akun.username != username {
			return errors.New("username salah")
		} else if akun.password != password {
			return errors.New("password salah")
		}

	}
	return nil
}
