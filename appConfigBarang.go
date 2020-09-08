package main

import "path"

type AppConfigBarang struct {
	dataPath string
	db       []*Barang
}

func NewConfig() *AppConfigBarang {
	const DATA_PATH = "data"
	const DATA_FILE_NAME = "data_barang.json"

	var dataBarang = make([]*Barang, 0)
	return &AppConfigBarang{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		db:       dataBarang,
	}
}

func (c *AppConfigBarang) getDb() []*Barang {
	return c.db
}

func (c *AppConfigBarang) getDbPath() string {
	return c.dataPath
}
