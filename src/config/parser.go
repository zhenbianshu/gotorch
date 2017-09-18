package config

import (
	"bufio"
	"fmt"
	"io"
	"logger"
	"os"
	"strings"
)

var instance *conf

type conf struct {
	data map[string]string
}

func newConf() *conf {
	root_path, _ := os.Getwd()
	file, err := os.Open(root_path + "/src/config/gotask.ini")
	fmt.Println(root_path + "/config/gotask.ini")
	if err != nil {
		log_info := map[string]string{"error": err.Error(), "desc": "error get config file!"}
		logger.Error(log_info)
		os.Exit(222)
	}
	defer file.Close()

	config_data := make(map[string]string)
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		parseLine(line, config_data)
	}

	conf := &conf{config_data}
	fmt.Println("new Configger")

	return conf
}

func parseLine(line []byte, config_data map[string]string) {
	if line[0] == '#' {
		return
	}

	var key, value []byte
	equal_sign := false

	for _, v := range line {
		if v == '#' {
			break
		}
		if v == '=' {
			equal_sign = true
			continue
		}

		if equal_sign {
			value = append(value, v)
		} else {
			key = append(key, v)
		}
	}

	if equal_sign {
		k := strings.Trim(string(key), " ")
		v := strings.Trim(string(value), " ")
		config_data[k] = v
	}

}

func GetConfig(field string) string {
	if instance == nil {
		instance = newConf()
	}

	return instance.data[field]
}
