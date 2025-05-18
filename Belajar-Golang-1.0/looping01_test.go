package main

import (
	"bytes"
	"os"
	"testing"
)

// TestNilaiMaksimumID menguji fungsi NilaiMaksimumID
func TestNilaiMaksimumID(t *testing.T) {
	data := []int{1, 5, 3, 9, 2}
	expected := 9
	result := NilaiMaksimumID(data)
	if result != expected {
		t.Errorf("NilaiMaksimumID salah, dapat %d, diharapkan %d", result, expected)
	}

	// Test irisan kosong
	empty := []int{}
	result = NilaiMaksimumID(empty)
	if result != 0 {
		t.Errorf("NilaiMaksimumID untuk irisan kosong harus 0, dapat %d", result)
	}
}

// TestCetakElemenID menguji fungsi CetakElemenID dengan menangkap output
func TestCetakElemenID(t *testing.T) {
	elements := []string{"apel", "jeruk", "pisang"}

	// Simpan stdout asli
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	CetakElemenID(elements)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	for _, e := range elements {
		if !bytes.Contains([]byte(output), []byte(e)) {
			t.Errorf("Output tidak mengandung elemen %s", e)
		}
	}
}

// TestJumlahAngkaGenapID menguji fungsi JumlahAngkaGenapID
func TestJumlahAngkaGenapID(t *testing.T) {
	n := 10
	expected := 0 + 2 + 4 + 6 + 8
	result := JumlahAngkaGenapID(n)
	if result != expected {
		t.Errorf("JumlahAngkaGenapID salah, dapat %d, diharapkan %d", result, expected)
	}
}

// TestContohLoopBersarangID menguji fungsi ContohLoopBersarangID
func TestContohLoopBersarangID(t *testing.T) {
	// Fungsi ini hanya mencetak output, tidak mengembalikan nilai
	// Hanya memeriksa tidak ada panic saat dijalankan
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ContohLoopBersarangID panic: %v", r)
		}
	}()
	ContohLoopBersarangID(3)
}

// TestLoopSepertiWhileID menguji fungsi LoopSepertiWhileID
func TestLoopSepertiWhileID(t *testing.T) {
	// Fungsi ini hanya mencetak output, tidak mengembalikan nilai
	// Hanya memeriksa tidak ada panic saat dijalankan
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("LoopSepertiWhileID panic: %v", r)
		}
	}()
	LoopSepertiWhileID(3)
}
