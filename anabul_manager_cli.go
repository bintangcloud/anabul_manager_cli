package main

import (
	"bufio"
	"fmt"
	"os"
)

type Kucing struct {
	Nama   string
	Umur   string
	Status bool
}

var databaseKucing []Kucing

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n === APLIKASI ANABUL MANAGER ===")
		fmt.Println("1. Tambah Kucing")
		fmt.Println("2. Lihat Semua Kucing")
		fmt.Println("3. Cek Kucing Sakit")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih Menu 1-4:")

		scanner.Scan()
		pilihan := scanner.Text()

		if pilihan == "1" {
			tambahKucing(scanner)
		} else if pilihan == "2" {
			lihatKucing()
		} else if pilihan == "3" {
			cekKucingStatus()
		} else if pilihan == "4" {
			fmt.Println("Dadah Meong!")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahKucing(scanner *bufio.Scanner) {
	fmt.Print("Masukkan nama kucing:")
	scanner.Scan()
	InputNama := scanner.Text()

	fmt.Print("Masukkan umur kucing:")
	scanner.Scan()
	inputUmur := scanner.Text()

	fmt.Print("Masukkan status kucing (Sehat/Sakit):")
	scanner.Scan()
	inputStatus := scanner.Text()

	var jikaSakit bool
	if inputStatus == "Sakit" {
		jikaSakit = true
	} else if inputStatus == "Sehat" {
		jikaSakit = false
	} else {
		fmt.Println("Masukkan status yang valid! Dianggap sehat dulu ya")
		jikaSakit = false
	}

	KucingBaru := Kucing{
		Nama:   InputNama,
		Umur:   inputUmur,
		Status: jikaSakit,
	}

	databaseKucing = append(databaseKucing, KucingBaru)

	fmt.Printf("Berhasil menyimpan si %s !\n", InputNama)

}

func lihatKucing() {
	fmt.Println("--Data Semua Kucingmu--")
	for i, data := range databaseKucing {
		var teksStatus string
		if data.Status == true {
			teksStatus = "Sakit"
		} else {
			teksStatus = "Sehat"
		}

		fmt.Printf("%d. %s %s tahun, %s \n", i+1, data.Nama, data.Umur, teksStatus)
	}
}

func cekKucingStatus() {
	fmt.Println("Segera periksakan. Ini data anabulmu yang sakit:")
	for i, data := range databaseKucing {
		if data.Status == true {
			fmt.Printf("%d. %s %s tahun\n", i+1, data.Nama, data.Umur)
		}
	}
}
