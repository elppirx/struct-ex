package bins

type BinList struct {
	bins []Bin
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{
		bins: bins,
	}
}
