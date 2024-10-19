package data

type Database interface {
	UpdateData()
	DeleteData()
	TampilkanSemuaData()
}

type DataPegawai struct {
	Nama, Kota, Negara string
}

type DataKonsumen struct {
	Nama, Alamat string
	Saldo        Saldo
}

type Saldo struct {
	Kas int
}

var DataDataPegawai []DataPegawai
var DataDataKonsumen []DataKonsumen
