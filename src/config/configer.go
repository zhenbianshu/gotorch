package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var instance *conf

type conf struct {
	data map[string]string
}

func newConf() *conf {
	rootPath, _ := os.Getwd()
	file, err := os.Open(rootPath + "/src/config/gotask.ini")
	if err != nil {
		os.Exit(222)
	}
	defer file.Close()

	configData := make(map[string]string)
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		parseLine(line, configData)
	}

	conf := &conf{configData}

	return conf
}

func parseLine(line []byte, configData map[string]string) {
	if len(line) <= 0 {
		return
	}
	if line[0] == '#' {
		return
	}

	var key, value []byte
	equalSign := false

	for _, v := range line {
		if v == '#' {
			break
		}
		if v == '=' {
			equalSign = true
			continue
		}

		if equalSign {
			value = append(value, v)
		} else {
			key = append(key, v)
		}
	}

	if equalSign {
		k := strings.Trim(string(key), " ")
		v := strings.Trim(string(value), " ")
		configData[k] = v
	}

}

func GetConfig(field string) string {
	if instance == nil {
		instance = newConf()
	}

	return instance.data[field]
}
