package main

import "github.com/devMiguelFerrer/EasyMonitoring/pkg/proxy"

const remoteURL = "http://localhost:8081"
const proxyPort = 7000

func main() {
	proxy.Create(remoteURL, proxyPort)
}
