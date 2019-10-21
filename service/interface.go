package service

type Idb interface {
	Put(string, string) bool
	Get(string) string
	GetAll() []string
}

type Idata interface {
	Data()
}
