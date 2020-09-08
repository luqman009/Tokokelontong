package main

import (
	"crypto/md5"
	"fmt"
)

type TransaksiReposistory struct {
	transaksiCollection      *[]*Transaksi
	tarnsaksiCollectionInfra *TransaksiRepositoryInfrastructure
}

func NewTransaksiRepository(dataPath string, transaksiCollection []*Transaksi) *TransaksiReposistory {
	transaksiRepoInfrac := NewTransaksiRepoInfra(dataPath)
	return &TransaksiReposistory{transaksiCollection: &transaksiCollection, tarnsaksiCollectionInfra: transaksiRepoInfrac}
}

func (br *TransaksiReposistory) AddNewTransaksi(transaksi *Transaksi) {
	data := []byte(transaksi.Barang)
	transaksi.KodeTransaksi = fmt.Sprintf("%x", md5.Sum(data))
	*br.transaksiCollection = append(*br.transaksiCollection, transaksi)
	br.tarnsaksiCollectionInfra.saveToFile(br.transaksiCollection)
}

func (br *TransaksiReposistory) FindAllTransaksi() []*Transaksi {
	br.tarnsaksiCollectionInfra.readFile(br.transaksiCollection)
	return *br.transaksiCollection
}
