package main

import (
	ini "2023Project/inianalysis"
	"fmt"
	"io/ioutil"
)

func UnmarshalIni(filename string, dbconf interface{}) {
	data, err := ioutil.ReadFile(filename)
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
