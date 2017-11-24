package config

import (
	"bufio"
	"io"
	"os"
	"strings"
	"fmt"
	"runtime"
)

var instance *conf

const confFileMac = "/tmp/gotorch.conf"
const confFileLinux = "/etc/gotorch.conf"
const confDefault = `# task list file
tasks = /tmp/gotorch/task.json

# log directory
log_dir = /tmp/gotorch/

# shell evn
bash = /bin/bash

# mail receiver when problem
mail_to = zhenbianshu@foxmail.com

# running service pid file
pid_file = /tmp/gotorch.pid

# interval time check if there is task can be exec (ms)
interval = 100`

type conf struct {
	data map[string]string
}

func newConf() *conf {
	var confFile string
	if runtime.GOOS == "darwin" {
		confFile = confFileMac
	} else {
		confFile = confFileLinux
	}
	file, err := os.Open(confFile)
	if err != nil {
		file.Close()
		// if config file not exists, then add one
		defaultFile, err := os.OpenFile(confFile, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(222)
		}
		defaultFile.Write([]byte(confDefault))

		file, _ = os.Open(confFile)
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

// parse config line and save configs
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
