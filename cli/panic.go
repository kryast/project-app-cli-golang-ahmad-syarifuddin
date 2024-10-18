package cli

import "fmt"

func ForceExit(err error) {
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		panic(err) // Menggunakan panic untuk keluar dari program
	}
}
