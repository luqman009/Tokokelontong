package main

type TransaksiService struct {
	r *TransaksiReposistory
}

func NewTransaksiService(repo *TransaksiReposistory) *TransaksiService {
	return &TransaksiService{r: repo}
}

func (bs *TransaksiService) registerTransaksi(t *Transaksi) {
	bs.r.AddNewTransaksi(t)
}

func (bs *TransaksiService) getAllTransaksi() []*Transaksi {
	return bs.r.FindAllTransaksi()
}
