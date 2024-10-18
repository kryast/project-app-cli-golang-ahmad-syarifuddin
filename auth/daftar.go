package auth

import "fmt"

type Akun struct {
	username, password string
}

var DataAkun []Akun

func DaftarAkun() (Akun, []Akun) {
	var username string
	var password string

	fmt.Print("Masukkan Username :")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password :")
	fmt.Scan(&password)

	DataAkun = append(DataAkun, Akun{username, password})
	return Akun{username, password}, DataAkun
}
