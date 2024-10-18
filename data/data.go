package data

import "fmt"

type Data struct {
	Nama, Kota, Negara string
}

var DataData []Data

func InputData() (Data, []Data) {
	var Nama string
	var Kota string
	var Negara string

	fmt.Print("Masukkan Nama :")
	fmt.Scan(&Nama)
	fmt.Print("Masukkan Kota asal :")
	fmt.Scan(&Kota)
	fmt.Print("Masukkan Negara asal :")
	fmt.Scan(&Negara)

	DataData = append(DataData, Data{Nama, Kota, Negara})
	fmt.Println("Data berhasil ditambah!")
	fmt.Println("-----------------------")
	return Data{Nama, Kota, Negara}, DataData
}

func UpdateData() {
	var Nama string
	fmt.Print("Masukkan Nama yang ingin diperbarui: ")
	fmt.Scan(&Nama)

	for i, data := range DataData {
		if data.Nama == Nama {
			var NamaBaru, KotaBaru, NegaraBaru string
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scan(&NamaBaru)
			fmt.Print("Masukkan Kota baru: ")
			fmt.Scan(&KotaBaru)
			fmt.Print("Masukkan Negara baru: ")
			fmt.Scan(&NegaraBaru)

			DataData[i].Nama = NamaBaru
			DataData[i].Kota = KotaBaru
			DataData[i].Negara = NegaraBaru

			fmt.Println("Data berhasil diperbarui!")
			fmt.Println("-------------------------")
			return
		}
	}

	fmt.Println("Data tidak ditemukan!")
	fmt.Println("---------------------")
}

func DeleteData() {
	var Nama string
	fmt.Print("Masukkan Nama yang ingin dihapus: ")
	fmt.Scan(&Nama)

	for i, data := range DataData {
		if data.Nama == Nama {
			DataData = append(DataData[:i], DataData[i+1:]...)
			fmt.Println("Data berhasil dihapus!")
			fmt.Println("----------------------")
			return
		}
	}

	fmt.Println("Data tidak ditemukan!")
	fmt.Println("---------------------")
}

func TampilkanSemuaData() {
	if len(DataData) == 0 {
		fmt.Println("Tidak ada data yang tersedia.")
		fmt.Println("-----------------------------")
		return
	}

	fmt.Println("=============================================")
	fmt.Println("Daftar Semua Data:")
	fmt.Println("=============================================")
	for _, data := range DataData {
		fmt.Printf("Nama: %s, Kota: %s, Negara: %s\n", data.Nama, data.Kota, data.Negara)
		fmt.Println("----------------------------------------------")
	}
}
