// Versi bahasa Indonesia dengan nama fungsi berbeda untuk menghindari duplikasi

package main

import "fmt"

// NilaiMaksimumID mengembalikan nilai maksimum dari irisan angka
func NilaiMaksimumID(angka []int) int {
	if len(angka) == 0 {
		return 0 // atau nilai error
	}

	maks := angka[0]
	for _, n := range angka {
		if n > maks {
			maks = n
		}
	}
	return maks
}

// CetakElemenID mencetak semua elemen dari irisan string
func CetakElemenID(elemen []string) {
	for i := 0; i < len(elemen); i++ {
		fmt.Println(elemen[i])
	}
}

// JumlahAngkaGenapID menjumlahkan semua angka genap dari 0 sampai n (eksklusif)
func JumlahAngkaGenapID(n int) int {
	jumlah := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			jumlah += i
		}
	}
	return jumlah
}

// ContohLoopBersarangID menunjukkan loop bersarang dengan mencetak tabel perkalian
func ContohLoopBersarangID(ukuran int) {
	for i := 1; i <= ukuran; i++ {
		for j := 1; j <= ukuran; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

// LoopSepertiWhileID menunjukkan loop seperti while menggunakan for
func LoopSepertiWhileID(batas int) {
	i := 0
	for i < batas {
		fmt.Println(i)
		i++
	}
}
