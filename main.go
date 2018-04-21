package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fatih/color"
)

var (
	phone string
	email string
	user  string
)

const banner = `
_________          _______  _______  _       
\__   __/|\     /|(  ___  )(  ____ )( (    /|
   ) (   | )   ( || (   ) || (    )||  \  ( |
   | |   | (___) || |   | || (____)||   \ | |
   | |   |  ___  || |   | ||     __)| (\ \) |
   | |   | (   ) || |   | || (\ (   | | \   |
   | |   | )   ( || (___) || ) \ \__| )  \  |
   )_(   |/     \|(_______)|/   \__/|/    )_)
`

func init() {
	flag.StringVar(&phone, "p", "", "手机号码")
	flag.StringVar(&email, "e", "", "邮箱地址")
	flag.StringVar(&user, "u", "", "用户名")
	flag.Parse()

	if phone == "" && email == "" && user == "" {
		flag.Usage()
		os.Exit(-1)
	}
}

func main() {
	var (
		checkType string
		keyword   string
	)

	wg := new(sync.WaitGroup)

	pluginsPath, err := filepath.Glob("./plugins/*.json")
	if err != nil {
		log.Fatalln(err)
	}

	color.Green(banner)
	println("\n[*] App: Search Registration")
	println("[*] Version: V0.1(2018-04-20)\n")

	switch {
	case phone != "":
		checkType = "phone"
		keyword = phone
		color.Yellow("[+] Phone Checking: %s", phone)
	case user != "":
		checkType = "user"
		keyword = user
		color.Yellow("[+] Username Checking: %s", user)
	case email != "":
		checkType = "email"
		keyword = email
		color.Yellow("[+] Email Checking: %s", email)
	}

	plugins := readConfig(pluginsPath)
	for _, v := range plugins {
		wg.Add(1)

		go func(data *plugin) {
			defer wg.Done()
			target := data.getURL(checkType)
			if target == "" {
				return
			}

			if data.Request.Method == "POST" {
				post(target, keyword, data)
			} else {
				get(target, keyword, data)
			}
		}(v)
	}
	wg.Wait()
}
