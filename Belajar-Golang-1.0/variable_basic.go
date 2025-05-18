package main

import "fmt"

// Biodata holds personal information
var Biodata = struct {
	Name    string
	Age     int
	Address string
	Email   string
}{
	Name:    "Efendi Sugiantoro",
	Age:     25,
	Address: "Jl. Veteran No. 21, Kediri, Jawa Timur, Indonesia",
	Email:   "efendi.sugiantoro@example.com",
}

// GetBiodataInfo returns formatted biodata information
func GetBiodataInfo() string {
	return fmt.Sprintf("Name: %s\nAge: %d\nAddress: %s\nEmail: %s", Biodata.Name, Biodata.Age, Biodata.Address, Biodata.Email)
}
