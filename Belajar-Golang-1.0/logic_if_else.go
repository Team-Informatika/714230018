package main

import "fmt"

func LogicIfElseExample(nilai int) string {
	switch {
	case nilai >= 80:
		return "Nilai Anda A"
	case nilai >= 70:
		return "Nilai Anda B"
	case nilai >= 60:
		return "Nilai Anda C"
	case nilai >= 50:
		return "Nilai Anda D"
	default:
		return "Nilai Anda E"
	}
}

// MenuItem represents a food menu item
type MenuItem struct {
	Name  string
	Price int
}

func ShowMenu(menu []MenuItem) {
	fmt.Println("Menu Makanan:")
	for i, item := range menu {
		fmt.Printf("%d. %s - Rp%d\n", i+1, item.Name, item.Price)
	}
	fmt.Println("0. Selesai")
}

func KasirMenu() {
	menu := []MenuItem{
		{"Nasi Goreng", 15000},
		{"Mie Goreng", 12000},
		{"Sate Ayam", 20000},
		{"Bakso", 18000},
		{"Es Teh", 5000},
		{"Es Jeruk", 7000},
	}

	order := make(map[int]int)
	var choice, qty int

	for {
		ShowMenu(menu)
		fmt.Print("Pilih nomor menu (0 untuk selesai): ")
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		if choice < 1 || choice > len(menu) {
			fmt.Println("Pilihan tidak valid, coba lagi.")
			continue
		}

		fmt.Print("Masukkan jumlah pesanan: ")
		fmt.Scan(&qty)
		if qty <= 0 {
			fmt.Println("Jumlah harus lebih dari 0, coba lagi.")
			continue
		}

		order[choice-1] += qty
		fmt.Println("Pesanan ditambahkan.")
	}

	if len(order) == 0 {
		fmt.Println("Tidak ada pesanan.")
		return
	}

	fmt.Println("\nRingkasan Pesanan:")
	total := 0
	for idx, qty := range order {
		item := menu[idx]
		subtotal := item.Price * qty
		fmt.Printf("%s x%d = Rp%d\n", item.Name, qty, subtotal)
		total += subtotal
	}
	fmt.Printf("Total Bayar: Rp%d\n", total)
}

// Example usage function to demonstrate the logic
func Example() {
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	grade := LogicIfElseExample(nilai)
	fmt.Println(grade)

	fmt.Println("\n--- Kasir Menu ---")
	KasirMenu()
}
