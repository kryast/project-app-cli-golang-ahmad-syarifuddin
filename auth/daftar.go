package auth

import (
	"errors"
	"fmt"
)

type Akun struct {
	username, password string
}

var DataAkun []Akun

func DaftarAkun() (Akun, []Akun, error) {
	var username string
	var password string

	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	if err := validasiAkun(username, password); err != nil {
		return Akun{}, DataAkun, err
	}

	DataAkun = append(DataAkun, Akun{username, password})
	return Akun{username, password}, DataAkun, nil
}

func validasiAkun(username, password string) error {
	for _, akun := range DataAkun {
		if akun.username == username {
			return errors.New("username sudah digunakan")
		}
	}
	if len(username) < 3 {
		return errors.New("username minimal 3 karakter")
	}
	if len(password) < 6 {
		return errors.New("password minimal 6 karakter")
	}

	return nil
}
