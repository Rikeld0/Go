package service

type Idb interface {
	Put(string, string) bool
	Get(string) string
}

/*type Data interface {
	Ss()
	Dd()
}*/
