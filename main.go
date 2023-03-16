package main

import (
	ini "UnmarshalIni/inianalysis"
	"fmt"
	"os"
)

func UnmarshalIni(filename string, dbconf interface{}) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	err = ini.Unmarshal(data, dbconf)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var dbcf ini.Config
	UnmarshalIni("./config.ini", &dbcf)
	fmt.Printf("%#v", dbcf)
}
