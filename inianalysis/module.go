package ini

import (
	"reflect"
	"strconv"
	"strings"
)

type Server struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type Mysql struct {
	Username            string  `ini:"username"`
	Passwd              string  `ini:"passwd"`
	Database            string  `ini:"database"`
	Host                string  `ini:"host"`
	Port                int     `ini:"port"`
	Timeout             float64 `ini:"timeout"`
	DefaultCharacterSet string  `ini:"default-character-set"`
}

type Config struct {
	SvcCfg Server `ini:"server"`
	DbCfg  Mysql  `ini:"mysql"`
}

// 数据分割
func datahandling(data []byte) (dataline []string, err error) {
	//数据按行分割
	datalinetmp := strings.Split(string(data), "\n")

	//忽略注释，去除空格
	for _, d := range datalinetmp {
		if len(d) == 0 {
			continue
		}
		if string(d[0]) == "#" || string(d[0]) == ";" {
			continue
		}
		dataline = append(dataline, strings.Replace(d, " ", "", -1))
	}
	return
}

// 大标签处理
func bigLabel(line string, iniTF reflect.Type) (bigField string, err error) {
	line = line[1 : len(line)-1]

	//拿到大标签的值
	for i := 0; i < iniTF.Elem().NumField(); i++ {
		if line == iniTF.Elem().Field(i).Tag.Get("ini") {
			bigField = iniTF.Elem().Field(i).Name
			break
		}
	}
	return
}

// 普通数据处理
func normalData(FieldName string, line string, iniVF reflect.Value) (err error) {
	normalK := line[:strings.Index(line, "=")]
	normalV := line[strings.Index(line, "=")+1:]

	//通过大标签的值获取对应的结构体的结构体
	instantlyCarrier := iniVF.Elem().FieldByName(FieldName)
	//遍历结构体对比key和结构体tag获取结构体内字段名称
	var keyName string
	for i := 0; i < instantlyCarrier.Type().NumField(); i++ {
		if normalK == instantlyCarrier.Type().Field(i).Tag.Get("ini") {
			keyName = instantlyCarrier.Type().Field(i).Name
		}
	}

	normalkValue := instantlyCarrier.FieldByName(keyName)
	switch normalkValue.Type().Kind() {
	case reflect.String:
		normalkValue.SetString(normalV)
	case reflect.Int:
		I, err := strconv.ParseInt(normalV, 10, 64)
		if err != nil {
			return err
		}
		normalkValue.SetInt(I)
	case reflect.Float64:
		I, err := strconv.ParseFloat(normalV, 64)
		if err != nil {
			return err
		}
		normalkValue.SetFloat(I)
	}

	return
}

func Unmarshal(data []byte, ini interface{}) (err error) {
	iniTF := reflect.TypeOf(ini)
	iniVF := reflect.ValueOf(ini)
	//判断传入对象是否为指针
	if iniTF.Kind() != reflect.Ptr {
		panic("请传入指针类型")
	}

	//判断传入对象是否为指针中是否存的结构体
	if iniVF.Elem().Kind() != reflect.Struct {
		panic("请传入指针类型的结构体")
	}

	//处理读入的数据
	dataslice, err := datahandling(data)
	//此时拿到数据：[[server] ip=10.238.2.2 port=8080 [mysql] username=root passwd=admin database=test host=192.168.10.10 port=8000 timeout=1.2]
	if err != nil {
		return
	}

	//将处理的数据分类 区分大标签[]、和小标签k=v
	var FieldName string
	for _, line := range dataslice {
		//处理大标签
		if line[0] == '[' {
			FieldName, err = bigLabel(line, iniTF)
			if err != nil {
				return err
			}
		} else {
			err = normalData(FieldName, line, iniVF)
			if err != nil {
				return
			}
		}
	}
	return
}
