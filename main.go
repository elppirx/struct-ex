package main

import "time"

type BinList struct {
	bins []Bin
}

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{
		bins: bins,
	}
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}

func main() {}
