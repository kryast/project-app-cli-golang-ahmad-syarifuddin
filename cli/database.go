package cli

import "fmt"

func Database() {
	defer func() {
		if r := recover(); r != nil {
			ClearScreen()
			fmt.Println("Terjadi kesalahan yang tidak terduga")
			Home()
		}
	}()

	var input int
	fmt.Println("========================")
	fmt.Println("Selamat Datang Di UdinDB")
	fmt.Println("========================")
	fmt.Println("1. Input Data Pegawai")
	fmt.Println("2. Input Data Konsumen")
	fmt.Println("3. Update Data")
	fmt.Println("4. Delete Data")
	fmt.Println("5. Tampilkan Semua Data")
	fmt.Println("99. Exit")
	fmt.Print("Masukkan nomor menu: ")
	fmt.Scan(&input)
}
