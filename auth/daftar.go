package auth

type Akun struct {
	username, password string
}

func DaftarAkun(username string, password string) Akun {
	return Akun{username, password}
}

var DataAkun []Akun

func TampungAkun(Akun ...Akun) []Akun {

	DataAkun = append(DataAkun, Akun...)
	return DataAkun
}
