package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TokoDelivery struct {
	barangService    *BarangService
	transaksiService *TransaksiService
}

func NewBarangDelevery(c *AppConfigBarang) *TokoDelivery {
	repoBarang := NewBarangRepository(c.getDbPath(), c.getDbBarang())
	repoTransaksi := NewTransaksiRepository(c.getDbPath(), c.getDbTansaksi())
	barangService := NewBarangService(repoBarang)
	transaksiService := NewTransaksiService(repoTransaksi)
	return &TokoDelivery{barangService, transaksiService}
}

func (bd *TokoDelivery) create() {
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
		case "03":
			bd.registerationtransaksiForm()
		case "04":
			bd.viewDataTransaksiForm()
		case "q":
			isExit = true
		default:
			fmt.Println("Unknown menu code")
		}

	}
}

func (bd *TokoDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (bd *TokoDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Add Barang",
		"02": "View Barang",
		"03": "Add Transaksi ",
		"04": "View Transaksi ",
		"q":  " Quit aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "My 1st Book Conllection")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range bd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s %s\n", menuCode, appMenu[menuCode])
	}
}

func (bd *TokoDelivery) registerationBarangForm() {
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
func (bd *TokoDelivery) registerationtransaksiForm() {
	consoleClear()
	var namabarang string
	var jumlah int
	var tanggal int
	var bulan int
	var tahun int
	var confirmation string

	fmt.Println()
	fmt.Printf("%s\n", "Transaksi Registration Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Barang : ")
	sNamabarang, _ := scanner.ReadString('\n')
	namabarang = strings.TrimSpace(sNamabarang)
	fmt.Print("Jumlah  : ")
	sJumlah, _ := scanner.ReadString('\n')
	jumlah, _ = strconv.Atoi(strings.TrimSpace(sJumlah))
	fmt.Print("Masukan Tanggal Bulan Tahun")
	fmt.Print("Tanggal")
	sTanggal, _ := scanner.ReadString('\n')
	tanggal, _ = strconv.Atoi(strings.TrimSpace(sTanggal))
	fmt.Print("Bulan")
	sBulan, _ := scanner.ReadString('\n')
	bulan, _ = strconv.Atoi(strings.TrimSpace(sBulan))
	fmt.Print("Tahun")
	sTahun, _ := scanner.ReadString('\n')
	tahun, _ = strconv.Atoi(strings.TrimSpace(sTahun))

	fmt.Println("Save to collection ? : (Y/N)")
	fmt.Scanln(&confirmation)

	if confirmation == "y" {
		newTransaksi := NewTransaksi(namabarang, jumlah, tanggal, bulan, tahun)
		bd.transaksiService.registerTransaksi(&newTransaksi)
	}
	consoleClear()
	bd.mainMenuForm()
}

func (bd *TokoDelivery) viewDataBarangForm() {
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

func (bd *TokoDelivery) viewDataTransaksiForm() {
	consoleClear()
	transaksi := bd.transaksiService.getAllTransaksi()
	fmt.Println("")
	fmt.Printf("Data Transaksi\n")
	fmt.Printf("%s\n", strings.Repeat("=", 135))
	fmt.Printf("%-40s%-15s%-15s%-15s\n", "Kode transaksi", "Nama Barang", "Jumlah", "Tanggal transaksi")
	fmt.Printf("%s\n", strings.Repeat("=", 135))
	for _, b := range transaksi {
		fmt.Printf("%-40s%-15s%-15v%v-%v-%v\n", b.KodeTransaksi, b.Barang, b.Jumlah, b.Tanggal, b.Bulan, b.Tahun)
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
