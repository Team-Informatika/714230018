// package main

// import (
// 	"fmt"
// )

// // Item represents a single inventory item in SiManBar
// type Item struct {
// 	KodeBarang  string  // Item code
// 	NamaBarang  string  // Item name
// 	Kategori    string  // Item category
// 	Jumlah      int     // Quantity in stock
// 	HargaSatuan float64 // Price per unit
// }

// // Sample data for Sistem Informasi Manajemen Barang (SiManBar)
// var daftarBarang = []Item{
// 	{
// 		KodeBarang:  "BRG001",
// 		NamaBarang:  "Laptop Lenovo ThinkPad",
// 		Kategori:    "Elektronik",
// 		Jumlah:      15,
// 		HargaSatuan: 8500000,
// 	},
// 	{
// 		KodeBarang:  "BRG002",
// 		NamaBarang:  "Mouse Wireless Logitech",
// 		Kategori:    "Aksesoris Komputer",
// 		Jumlah:      50,
// 		HargaSatuan: 150000,
// 	},
// 	{
// 		KodeBarang:  "BRG003",
// 		NamaBarang:  "Meja Kerja Kayu",
// 		Kategori:    "Furniture",
// 		Jumlah:      10,
// 		HargaSatuan: 750000,
// 	},

// 	{
// 		KodeBarang:  "BRG003",
// 		NamaBarang:  "Meja Kerja Kayu",
// 		Kategori:    "Furniture",
// 		Jumlah:      10,
// 		HargaSatuan: 750000,
// 	},
// }

// // cetakLaporanSiManBar prints the SiManBar report listing all items with details
// func cetakLaporanSiManBar(barang []Item) {
// 	fmt.Println("Laporan Sistem Informasi Manajemen Barang (SiManBar)")
// 	fmt.Println("------------------------------------------------------")
// 	fmt.Printf("%-10s | %-25s | %-20s | %-7s | %-12s\n", "Kode", "Nama Barang", "Kategori", "Jumlah", "Harga Satuan")
// 	fmt.Println("----------------------------------------------------------------------------------------------")

// 	var totalNilai float64 = 0
// 	for _, b := range barang {
// 		fmt.Printf("%-10s | %-25s | %-20s | %-7d | Rp %10.2f\n", b.KodeBarang, b.NamaBarang, b.Kategori, b.Jumlah, b.HargaSatuan)
// 		totalNilai += float64(b.Jumlah) * b.HargaSatuan
// 	}
// 	fmt.Println("----------------------------------------------------------------------------------------------")
// 	fmt.Printf("Total Nilai Barang: Rp %.2f\n", totalNilai)
// }

// func variable() {
// 	cetakLaporanSiManBar(daftarBarang)
// }
