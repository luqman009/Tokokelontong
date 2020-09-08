package main

import (
	"crypto/md5"
	"fmt"
)

type BarangReposistory struct {
	barangCollection      *[]*Barang
	barangCollectionInfra *BarangRepositoryInfrastructure
}

func NewBarangRepository(dataPath string, barangCollection []*Barang) *BarangReposistory {
	barangRepoInfrac := NewBarangRepoInfra(dataPath)
	return &BarangReposistory{barangCollection: &barangCollection, barangCollectionInfra: barangRepoInfrac}
}

func (br *BarangReposistory) AddNewBarang(barang *Barang) {
	data := []byte(barang.NamaBarang)
	barang.KodeBarang = fmt.Sprintf("%x", md5.Sum(data))
	*br.barangCollection = append(*br.barangCollection, barang)
	br.barangCollectionInfra.saveToFile(br.barangCollection)
}

func (br *BarangReposistory) FindAllBarang() []*Barang {
	br.barangCollectionInfra.readFile(br.barangCollection)
	return *br.barangCollection
}
