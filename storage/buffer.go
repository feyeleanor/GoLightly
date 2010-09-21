package storage

type Buffer interface {
	Copy(i, j int)
	Swap(i, j int)
	Clear(i, n int)
}
