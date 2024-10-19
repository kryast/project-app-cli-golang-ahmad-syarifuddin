package data

import "fmt"

func (d *DataPegawai) UpdateData() {

	if len(DataDataPegawai) <= 0 {
		fmt.Println("Data Masih kosong")
		return
	}

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

	if len(DataDataKonsumen) <= 0 {
		fmt.Println("Data Masih kosong")
		return
	}

	var Nama string
	fmt.Print("Masukkan Nama Konsumen yang ingin diperbarui: ")
	fmt.Scan(&Nama)

	for i, data := range DataDataKonsumen {
		if data.Nama == Nama {
			var NamaBaru, AlamatBaru string
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scan(&NamaBaru)
			if NamaBaru == "" {
				fmt.Println("Error: Nama baru tidak boleh kosong.")
				return
			}

			fmt.Print("Masukkan Alamat baru: ")
			fmt.Scan(&AlamatBaru)
			if AlamatBaru == "" {
				fmt.Println("Error: Alamat tidak boleh kosong.")
				return
			}

			var SaldoBaru int
			fmt.Println("----------------------------------------------------------")
			fmt.Println("(Ketik berupa Angka, Jika tidak maka akan kembali Log out)")
			fmt.Print("Masukkan Saldo baru: ")
			fmt.Scan(&SaldoBaru)

			TotalKas := data.Saldo.Kas + SaldoBaru

			if TotalKas < 0 {
				fmt.Println("Error: Saldo tidak mencukupi.")
				return
			}

			DataDataKonsumen[i] = DataKonsumen{NamaBaru, AlamatBaru, Saldo{Kas: TotalKas}}
			fmt.Println("Data berhasil diperbarui!")
			fmt.Println("-------------------------")
			return
		}
	}

	fmt.Println("Data tidak ditemukan!")
	fmt.Println("---------------------")
}
