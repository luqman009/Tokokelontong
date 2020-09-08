package main

type BarangService struct {
	r *BarangReposistory
}

func NewBarangService(repo *BarangReposistory) *BarangService {
	return &BarangService{r: repo}
}

func (bs *BarangService) registerNewBarang(b *Barang) {
	bs.r.AddNewBarang(b)
}

func (bs *BarangService) getAllBarang() []*Barang {
	return bs.r.FindAllBarang()
}
