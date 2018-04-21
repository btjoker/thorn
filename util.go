package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/fatih/color"
)

// readConfig 读取插件配置
func readConfig(pluginsPath []string) []*plugin {
	plugins := make([]*plugin, 0, 30)
	for _, path := range pluginsPath {
		var p plugin
		doc, err := ioutil.ReadFile(path)
		if err != nil {
			continue
		}
		if err := json.Unmarshal(doc, &p); err != nil {
			continue
		}
		plugins = append(plugins, &p)
	}
	return plugins
}

// check 检查是否匹配配置文件里的前缀
func check(resp *http.Response, p *plugin) {
	var lock sync.Mutex
	doc, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	lock.Lock()
	defer lock.Unlock()
	print(p.name(), "\t")
	if bytes.HasPrefix(doc, p.judgeYesKeyword()) {
		color.Green("存在")
	} else {
		color.Red("不存在")
	}
}
