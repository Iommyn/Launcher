package main

import "LauncherServer/cmd/launcherserver"

func main() {
	webserver := new(launcherserver.WebServer)
	if err := webserver.StartLauncherService(); err != nil {
		panic(err)
	}
}
