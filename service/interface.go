package service

type Logger interface {
	Error(string)
}

type Idb interface {
	Put(string, string) bool
	Get(string) string
}
