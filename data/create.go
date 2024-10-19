package data

import "fmt"

func InputDataPegawai() (DataPegawai, []DataPegawai) {
	var Nama, Kota, Negara string

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&Nama)
	if len(Nama) < 3 {
		fmt.Println("minimal 3 karakter")
		return DataPegawai{}, DataDataPegawai
	}

	fmt.Print("Masukkan Kota asal: ")
	fmt.Scan(&Kota)
	if len(Kota) < 3 {
		fmt.Println("minimal 3 karakter")
		return DataPegawai{}, DataDataPegawai
	}

	fmt.Print("Masukkan Negara asal: ")
	fmt.Scan(&Negara)
	if len(Negara) < 3 {
		fmt.Println("minimal 3 karakter")
		return DataPegawai{}, DataDataPegawai
	}

	DataDataPegawai = append(DataDataPegawai, DataPegawai{Nama, Kota, Negara})
	fmt.Println("Data Pegawai berhasil ditambah!")
	fmt.Println("-----------------------")
	return DataPegawai{Nama, Kota, Negara}, DataDataPegawai
}

func InputDataKonsumen() (DataKonsumen, []DataKonsumen) {
	var Nama, Alamat string
	var Saldo Saldo

	fmt.Print("Masukkan Nama Konsumen: ")
	fmt.Scan(&Nama)
	if len(Nama) < 3 {
		fmt.Println("minimal 3 karakter")
		return DataKonsumen{}, DataDataKonsumen
	}

	fmt.Print("Masukkan Alamat Konsumen: ")
	fmt.Scan(&Alamat)
	if len(Alamat) < 3 {
		fmt.Println("minimal 3 karakter")
		return DataKonsumen{}, DataDataKonsumen
	}

	Saldo.Kas = 0
	DataDataKonsumen = append(DataDataKonsumen, DataKonsumen{Nama, Alamat, Saldo})
	fmt.Println("Data Konsumen berhasil ditambah!")
	fmt.Println("-----------------------")
	return DataKonsumen{Nama, Alamat, Saldo}, DataDataKonsumen
}
