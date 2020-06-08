package log

import (
	"fmt"
	"log"
	"os"
)

var f *os.File
var err error

//Init init func
func Init() {
	f, err = os.OpenFile("./file.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//Info info log
func Info(s string) {
	q := fmt.Sprintf("INFO: %s\n", s)
	_, err1 := f.WriteString(q)
	if err1 != nil {
		panic(err1)
	}
	f.Sync()
}

//Fatal fatal log
func Fatal(s string) {
	q := fmt.Sprintf("FATA: %s\n", s)
	_, err1 := f.WriteString(q)
	if err1 != nil {
		panic(err1)
	}
	f.Sync()
	panic("Fatal")
}

//Warn warn log
func Warn(s string) {
	q := fmt.Sprintf("WARN: %s\n", s)
	_, err1 := f.WriteString(q)
	if err1 != nil {
		panic(err1)
	}
	f.Sync()
}
