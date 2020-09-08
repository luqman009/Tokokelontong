package main

import (
	"encoding/json"
	"io/ioutil"
)

type TransaksiRepositoryInfrastructure struct {
	dataPath string
}

func NewTransaksiRepoInfra(dataPath string) *TransaksiRepositoryInfrastructure {
	return &TransaksiRepositoryInfrastructure{dataPath}
}

func (bri *TransaksiRepositoryInfrastructure) saveToFile(transaksiCollection *[]*Transaksi) {
	file, _ := json.MarshalIndent(transaksiCollection, "", " ")
	_ = ioutil.WriteFile(bri.dataPath, file, 0644)
}

func (bri *TransaksiRepositoryInfrastructure) readFile(TransaksiCollection *[]*Transaksi) {
	file, _ := ioutil.ReadFile(bri.dataPath)
	_ = json.Unmarshal(file, TransaksiCollection)
}
