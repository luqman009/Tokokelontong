package main

type Barang struct {
	KodeBarang string
	NamaBarang string
	Kategori   string
	Harga      float64
	Stok       int
	Deskripsi  string
	Ukuran     string
}

func NewBarang(namaBarang string, kategori string, harga float64, stok int, deskripsi string, ukuran string) Barang {
	return Barang{
		NamaBarang: namaBarang,
		Kategori:   kategori,
		Harga:      harga,
		Stok:       stok,
		Deskripsi:  deskripsi,
		Ukuran:     ukuran,
	}
}
