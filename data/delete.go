package data

import "fmt"

func (d *DataPegawai) DeleteData() {

	if len(DataDataPegawai) <= 0 {
		fmt.Println("Data Masih kosong")
		return
	}

	var Nama string
	fmt.Print("Masukkan Nama yang ingin dihapus: ")
	fmt.Scan(&Nama)

	for i, data := range DataDataPegawai {
		if data.Nama == Nama {
			DataDataPegawai = append(DataDataPegawai[:i], DataDataPegawai[i+1:]...)
			fmt.Println("Data berhasil dihapus!")
			fmt.Println("----------------------")
			return
		}
	}

	fmt.Println("Data tidak ditemukan!")
	fmt.Println("---------------------")
}

func (k *DataKonsumen) DeleteData() {

	if len(DataDataKonsumen) <= 0 {
		fmt.Println("Data Masih kosong")
		return
	}

	var Nama string
	fmt.Print("Masukkan Nama Konsumen yang ingin dihapus: ")
	fmt.Scan(&Nama)

	for i, data := range DataDataKonsumen {
		if data.Nama == Nama {
			DataDataKonsumen = append(DataDataKonsumen[:i], DataDataKonsumen[i+1:]...)
			fmt.Println("Data berhasil dihapus!")
			fmt.Println("----------------------")
			return
		}
	}

	fmt.Println("Data tidak ditemukan!")
	fmt.Println("---------------------")
}
