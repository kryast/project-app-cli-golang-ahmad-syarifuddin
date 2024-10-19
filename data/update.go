package data

import "fmt"

func (d *DataPegawai) UpdateData() {
	var Nama string
	fmt.Print("Masukkan Nama yang ingin diperbarui: ")
	fmt.Scan(&Nama)

	for i, data := range DataDataPegawai {
		if data.Nama == Nama {
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scan(&data.Nama)
			fmt.Print("Masukkan Kota baru: ")
			fmt.Scan(&data.Kota)
			fmt.Print("Masukkan Negara baru: ")
			fmt.Scan(&data.Negara)

			DataDataPegawai[i] = data
			fmt.Println("Data berhasil diperbarui!")
			fmt.Println("-------------------------")
			return
		}
	}

	fmt.Println("Data tidak ditemukan!")
	fmt.Println("---------------------")
}

func (k *DataKonsumen) UpdateData() {
	var Nama string
	fmt.Print("Masukkan Nama Konsumen yang ingin diperbarui: ")
	fmt.Scan(&Nama)

	for i, data := range DataDataKonsumen {
		if data.Nama == Nama {
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scan(&data.Nama)
			fmt.Print("Masukkan Alamat baru: ")
			fmt.Scan(&data.Alamat)
			fmt.Print("Masukkan Saldo baru: ")
			fmt.Scan(&data.Saldo.Kas)

			DataDataKonsumen[i] = data
			fmt.Println("Data berhasil diperbarui!")
			fmt.Println("-------------------------")
			return
		}
	}

	fmt.Println("Data tidak ditemukan!")
	fmt.Println("---------------------")
}
