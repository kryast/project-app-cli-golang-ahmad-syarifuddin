package auth

import (
	"errors"
	"fmt"
)

func Login(username string, password string) error {

	for _, akun := range DataAkun {
		if akun.username == username && akun.password == password {
			fmt.Printf("Data Akun %s Yaitu %+v\n ", username, password)
			return nil
		}

	}
	return errors.New("username dan password salah")
}
