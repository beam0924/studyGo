package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err := fmt.Errorf("data should be a pointer")
		return err
	}

	if t.Elem().Kind() != reflect.Struct {
		err := fmt.Errorf("data should be a struct")
		return err
	}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("%#v\n", lineSlice)

	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return err
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return err
			}
			// fmt.Println(t.Elem().NumField())
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name

					fmt.Printf("找到%s对应的嵌套结构体%s", sectionName, structName)
				}
			}

		} else {

		}

	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./config.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)
	// fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
