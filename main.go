package main

import (
	"github.com/devMiguelFerrer/EasyMonitoring/pkg/proxy"
	"github.com/devMiguelFerrer/EasyMonitoring/pkg/tracing"
)

const remoteURL = "http://localhost:9559"
const proxyPort = 7000

func main() {
	track := &tracing.Tracing{
		Host:           "localhost",
		Port:           27017,
		DBName:         "DBName",
		CollectionName: "CLName",
	}

	track.Connect()

	proxy.Create(remoteURL, proxyPort, track)
}
