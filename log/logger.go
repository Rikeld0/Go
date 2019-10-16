package log

import (
	"fmt"
	"log"
	"os"
)

/*var LogG Logger = &err{}

type err struct {
	Logger
}*/
var f *os.File
var err error

func Init() {
	//create your file with desired read/write permissions
	f, err = os.OpenFile("./file.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//defer to close when you're done with it, not because you think it's idiomatic!
}

func Info(s string) {
	q := fmt.Sprintf("INFO: %s\n", s)
	_, err1 := f.WriteString(q)
	if err1 != nil {
		panic(err1)
	}
	f.Sync()
}

func Fatal(s string) {
	q := fmt.Sprintf("FATA: %s\n", s)
	_, err1 := f.WriteString(q)
	if err1 != nil {
		panic(err1)
	}
	f.Sync()
	panic("Fatal")
}

func Warn(s string) {
	q := fmt.Sprintf("WARN: %s\n", s)
	_, err1 := f.WriteString(q)
	if err1 != nil {
		panic(err1)
	}
	f.Sync()
}
