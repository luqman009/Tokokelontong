package main

type Transaksi struct {
	KodeTransaksi string
	Barang        string
	Jumlah        int
	Tanggal       int
	Bulan         int
	Tahun         int
}

func NewTransaksi(namaBarang string, jumlah int, tanggal int, bulan int, tahun int) Transaksi {
	return Transaksi{
		Barang:  namaBarang,
		Jumlah:  jumlah,
		Tanggal: tanggal,
		Bulan:   bulan,
		Tahun:   tahun,
	}
}
