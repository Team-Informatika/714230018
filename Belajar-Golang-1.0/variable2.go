// Added import of variable.go content to fix undefined errors
package main

import (
	"errors"
	"fmt"
)

// Re-declare Item struct and daftarBarang from variable.go to fix undefined errors
type Item struct {
	KodeBarang  string
	NamaBarang  string
	Kategori    string
	Jumlah      int
	HargaSatuan float64
}

var daftarBarang = []Item{
	{
		KodeBarang:  "BRG001",
		NamaBarang:  "Laptop Lenovo ThinkPad",
		Kategori:    "Elektronik",
		Jumlah:      15,
		HargaSatuan: 8500000,
	},
	{
		KodeBarang:  "BRG002",
		NamaBarang:  "Mouse Wireless Logitech",
		Kategori:    "Aksesoris Komputer",
		Jumlah:      50,
		HargaSatuan: 150000,
	},
	{
		KodeBarang:  "BRG003",
		NamaBarang:  "Meja Kerja Kayu",
		Kategori:    "Furniture",
		Jumlah:      10,
		HargaSatuan: 750000,
	},
}

// Transaction represents a distribution transaction of goods
type Transaction struct {
	KodeBarang string  // Item code
	Jumlah     int     // Quantity distributed
	TotalHarga float64 // Total price of the distributed goods
}

// daftarTransaksi stores all distribution transactions
var daftarTransaksi []Transaction

// distribusiBarang performs a distribution transaction for a given item code and quantity
func distribusiBarang(kode string, jumlahDistribusi int) error {
	// Find the item in daftarBarang
	for i, item := range daftarBarang {
		if item.KodeBarang == kode {
			if item.Jumlah < jumlahDistribusi {
				return errors.New("stok barang tidak cukup untuk distribusi")
			}
			// Update stock quantity
			daftarBarang[i].Jumlah -= jumlahDistribusi

			// Calculate total price
			totalHarga := float64(jumlahDistribusi) * item.HargaSatuan

			// Create a new transaction record
			transaksi := Transaction{
				KodeBarang: kode,
				Jumlah:     jumlahDistribusi,
				TotalHarga: totalHarga,
			}
			daftarTransaksi = append(daftarTransaksi, transaksi)
			return nil
		}
	}
	return errors.New("barang dengan kode tersebut tidak ditemukan")
}

// cetakLaporanSiManBar prints the SiManBar report listing all items with details
func cetakLaporanSiManBar(barang []Item) {
	fmt.Println("Laporan Sistem Informasi Manajemen Barang (SiManBar)")
	fmt.Println("------------------------------------------------------")
	fmt.Printf("%-10s | %-25s | %-20s | %-7s | %-12s\n", "Kode", "Nama Barang", "Kategori", "Jumlah", "Harga Satuan")
	fmt.Println("----------------------------------------------------------------------------------------------")

	var totalNilai float64 = 0
	for _, b := range barang {
		fmt.Printf("%-10s | %-25s | %-20s | %-7d | Rp %10.2f\n", b.KodeBarang, b.NamaBarang, b.Kategori, b.Jumlah, b.HargaSatuan)
		totalNilai += float64(b.Jumlah) * b.HargaSatuan
	}
	fmt.Println("----------------------------------------------------------------------------------------------")
	fmt.Printf("Total Nilai Barang: Rp %.2f\n", totalNilai)
}

// cetakLaporanTransaksi prints all distribution transactions
func cetakLaporanTransaksi() {
	fmt.Println("Laporan Transaksi Distribusi Barang")
	fmt.Println("-----------------------------------")
	fmt.Printf("%-10s | %-10s | %-15s\n", "Kode", "Jumlah", "Total Harga")
	fmt.Println("-----------------------------------")

	var totalPendapatan float64 = 0
	for _, t := range daftarTransaksi {
		fmt.Printf("%-10s | %-10d | Rp %12.2f\n", t.KodeBarang, t.Jumlah, t.TotalHarga)
		totalPendapatan += t.TotalHarga
	}
	fmt.Println("-----------------------------------")
	fmt.Printf("Total Pendapatan: Rp %.2f\n", totalPendapatan)
}

// contohPenggunaan demonstrates usage of distribusiBarang and printing reports
func contohPenggunaan() {
	fmt.Println("Sebelum distribusi:")
	cetakLaporanSiManBar(daftarBarang)

	err := distribusiBarang("BRG001", 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Distribusi berhasil.")
	}

	fmt.Println("\nSetelah distribusi:")
	cetakLaporanSiManBar(daftarBarang)

	fmt.Println()
	cetakLaporanTransaksi()
}
