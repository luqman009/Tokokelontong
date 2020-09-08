package main

import "path"

type AppConfigBarang struct {
	dataPath    string
	dbBarang    []*Barang
	dbTransaksi []*Transaksi
}

func NewConfig() *AppConfigBarang {
	const DATA_PATH = "data"
	const DATA_FILE_NAME = "data_barang.json"
	const DATA_FILE_NAME1 = "data_transaksi.json"

	var dataBarang = make([]*Barang, 0)
	return &AppConfigBarang{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		dbBarang: dataBarang,
	}

	var dataTransaksi = make([]*Transaksi, 0)
	return &AppConfigBarang{
		dataPath:    path.Join(DATA_PATH, DATA_FILE_NAME1),
		dbTransaksi: dataTransaksi,
	}
}

func (c *AppConfigBarang) getDbBarang() []*Barang {
	return c.dbBarang
}

func (c *AppConfigBarang) getDbPath() string {
	return c.dataPath
}

func (c *AppConfigBarang) getDbTansaksi() []*Transaksi {
	return c.dbTransaksi
}
