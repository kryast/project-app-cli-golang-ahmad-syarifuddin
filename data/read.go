package data

import "fmt"

func (d *DataPegawai) TampilkanSemuaData() {
	if len(DataDataPegawai) == 0 {
		fmt.Println("Tidak ada data yang tersedia.")
		fmt.Println("-----------------------------")
		return
	}

	fmt.Println("=============================================")
	fmt.Println("Daftar Semua Data:")
	fmt.Println("=============================================")
	for _, data := range DataDataPegawai {
		fmt.Printf("Nama: %s, Kota: %s, Negara: %s\n", data.Nama, data.Kota, data.Negara)
		fmt.Println("----------------------------------------------")
	}
}

func (k *DataKonsumen) TampilkanSemuaData() {
	if len(DataDataKonsumen) == 0 {
		fmt.Println("Tidak ada data yang tersedia.")
		fmt.Println("-----------------------------")
		return
	}

	fmt.Println("=============================================")
	fmt.Println("Daftar Semua Data Konsumen:")
	fmt.Println("=============================================")
	for _, data := range DataDataKonsumen {
		fmt.Printf("Nama: %s, Alamat: %s, Saldo: %d\n", data.Nama, data.Alamat, data.Saldo.Kas)
		fmt.Println("----------------------------------------------")
	}
}
