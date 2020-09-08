package main

import (
	"encoding/json"
	"io/ioutil"
)

type BarangRepositoryInfrastructure struct {
	dataPath string
}

func NewBarangRepoInfra(dataPath string) *BarangRepositoryInfrastructure {
	return &BarangRepositoryInfrastructure{dataPath}
}

func (bri *BarangRepositoryInfrastructure) saveToFile(bookCollection *[]*Barang) {
	file, _ := json.MarshalIndent(bookCollection, "", " ")
	_ = ioutil.WriteFile(bri.dataPath, file, 0644)
}

func (bri *BarangRepositoryInfrastructure) readFile(bookCollection *[]*Barang) {
	file, _ := ioutil.ReadFile(bri.dataPath)
	_ = json.Unmarshal(file, bookCollection)
}
