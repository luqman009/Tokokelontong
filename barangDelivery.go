package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BarangDelivery struct {
	barangService *BarangService
}

func NewBarangDelevery(c *AppConfigBarang) *BarangDelivery {
	repo := NewBarangRepository(c.getDbPath(), c.getDb())
	barangService := NewBarangService(repo)
	return &BarangDelivery{barangService}
}

type TransaksigDelivery struct {
	transaksiService *TransaksiService
}

func NewtransaksiDelevery(c *AppConfigTransaksi) *TransaksigDelivery {
	repo := NewTransaksiRepository(c.getDbPath(), c.getDb())
	transaksiService := NewTransaksiService(repo)
	return &TransaksigDelivery{transaksiService}
}

func (bd *BarangDelivery) create() {
	var isExit = false
	var userChoice string

	bd.mainMenuForm()
	for isExit == false {
		fmt.Printf("\n%s", "Your choice : ")
		fmt.Scanln(&userChoice)
		switch userChoice {
		case "01":
			bd.registerationBarangForm()
		case "02":
			bd.viewDataBarangForm()
		case "q":
			isExit = true
		default:
			fmt.Println("Unknown menu code")
		}
	}
}

func (bd *BarangDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (bd *BarangDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Add book to collection",
		"02": "View book to collection",
		"q":  " Quit aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "My 1st Book Conllection")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range bd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s %s\n", menuCode, appMenu[menuCode])
	}
}

func (bd *BarangDelivery) registerationBarangForm() {
	consoleClear()
	var namabarang string
	var kategori string
	var harga float64
	var stok int
	var deskripsi string
	var ukuran string
	var confirmation string

	fmt.Println()
	fmt.Printf("%s\n", "Barang Registration Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Barang : ")
	sNamabarang, _ := scanner.ReadString('\n')
	namabarang = strings.TrimSpace(sNamabarang)
	fmt.Print("Kategori : ")
	sKategori, _ := scanner.ReadString('\n')
	kategori = strings.TrimSpace(sKategori)
	fmt.Print("Harga : ")
	sHarga, _ := scanner.ReadString('\n')
	harga, _ = strconv.ParseFloat(strings.TrimSpace(sHarga), 64)
	fmt.Print("Stok : ")
	sStock, _ := scanner.ReadString('\n')
	stok, _ = strconv.Atoi(strings.TrimSpace(sStock))
	fmt.Print("Deskripsi : ")
	sDeskripsi, _ := scanner.ReadString('\n')
	deskripsi = strings.TrimSpace(sDeskripsi)
	fmt.Print("Ukuran : ")
	sUkuran, _ := scanner.ReadString('\n')
	ukuran = strings.TrimSpace(sUkuran)

	fmt.Println("Save to collection ? : (Y/N)")
	fmt.Scanln(&confirmation)

	if confirmation == "y" {
		newBarang := NewBarang(namabarang, kategori, harga, stok, deskripsi, ukuran)
		bd.barangService.registerNewBarang(&newBarang)
	}
	consoleClear()
	bd.mainMenuForm()
}

func (bd *BarangDelivery) viewDataBarangForm() {
	consoleClear()
	barang := bd.barangService.getAllBarang()
	fmt.Println("")
	fmt.Printf("Data Barang\n")
	fmt.Printf("%s\n", strings.Repeat("=", 135))
	fmt.Printf("%-40s%-15s%-15s%-15s%-10s%-25s%-10s\n", "Kode", "Nama", "Kategori", "Harga", "Stok", "Deskripsi", "Ukuran")
	fmt.Printf("%s\n", strings.Repeat("=", 135))
	for _, b := range barang {
		fmt.Printf("%-40s%-15s%-15s%-15.2f%-10v%-25s%-15s\n", b.KodeBarang, b.NamaBarang, b.Kategori, b.Harga, b.Stok, b.Deskripsi, b.Ukuran)
	}

	for {
		var kembali string
		fmt.Print("Back To menu ? : (Y)")
		fmt.Scanln(&kembali)
		if kembali == "y" {
			consoleClear()
			bd.mainMenuForm()
			break

		} else {
			fmt.Println("inputan salah")
		}

	}
}
