package main

import (
	"crypto/tls"
	_"mm-wiki/app"
	_"github.com/astaxie/beego/session/memcache"
	_"github.com/astaxie/beego/session/redis"
	_"github.com/astaxie/beego/session/redis_cluster"
	"github.com/astaxie/beego"
	//_"github.com/astaxie/BeeApp/Server"
)

func main() {
	beego.BeeApp.Server.TLSConfig = &tls.Config{ MinVersion: 0x0302}
	beego.Run()
}
