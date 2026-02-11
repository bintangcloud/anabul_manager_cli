package main

import (
	"bufio"
	"fmt"
	"os"
)

type Kucing struct {
	Nama   string
	Umur   float
	Status bool
}

var databaseKucing []Kucing

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Anabul Manager")
		fmt.Println("1. Tambah Kucing")
		fmt.Println("2. Lihat Semua Kucing")
		fmt.Println("3. Cek Kucing Sakit")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih Menu 1-4:")

		scanner.Scan()
		pilihan := scanner.Text()

		if pilihan == 1 {
			tambahKucing(scanner)
		} else if pilihan == 2 {
			lihatKucing()
		} else if pilihan == 3 {
			cekKucingStatus()
		} else if pilihan == 4 {
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
	if status == "Sehat" {
		jikaSakit = true
	} else if status == "Sakit" {
		jikaSakit = false
	} else {
		fmt.Println("Masukkan status yang valid!(Sehat/Sakit)")
		jikaSakit = true
	}

	KucingBaru := Kucing{
		Nama:      InputNama,
		Umur:      inputUmur,
		jikaSakit: jikaSakit,
	}

	databaseKucing = append(databaseKucing, KucingBaru())

}
