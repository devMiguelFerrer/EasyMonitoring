package main

import (
	"github.com/devMiguelFerrer/EasyMonitoring/pkg/proxy"
	"github.com/devMiguelFerrer/EasyMonitoring/pkg/tracing"
)

const remoteURL = "http://localhost:7000"
const proxyPort = 7001

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
