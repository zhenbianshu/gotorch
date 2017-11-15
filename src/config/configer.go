package config

import (
	"bufio"
	"io"
	"os"
	"strings"
	"logger"
)

var instance *conf

const confFile = "/etc/gotorch.conf"
const confDefault = `# 配置文件
tasks = /tmp/gotorch/task.json

# 日志目录
log_dir = /tmp/gotorch/

# 使用的shell环境
bash = /bin/bash

# 出问题时的邮件通知人
mail_to = zhenbianshu@foxmail.com

# 进程PID文件
pid_file = /tmp/gotorch.pid

# 每轮任务的间隔时间 (ms)
interval = 100`

type conf struct {
	data map[string]string
}

func newConf() *conf {
	file, err := os.Open(confFile)
	if err != nil {
		// 如果没有此文件，则将默认配置写入其中
		logger.Warning("bootstrap", "open conf file error:"+err.Error())
		defaultFile, _ := os.OpenFile(confFile, os.O_WRONLY|os.O_CREATE, 0644)
		defer defaultFile.Close()
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
