package bins

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{
		Bins: bins,
	}
}
