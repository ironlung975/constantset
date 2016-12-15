package subset

type ConstantSubset interface {
	Add(Comparable)
	Get(int) Comparable
	JoinNext(*ConstantSubset)
	JoinPrev(*ConstantSubset)
	Length() int
	Next() *ConstantSubset
	Prev() *ConstantSubset
}

type constantSubsetImpl struct {
	items []Comparable
	keystore map[Comparable]int
	prev *constantSubsetImpl
	next *constantSubsetImpl
}
