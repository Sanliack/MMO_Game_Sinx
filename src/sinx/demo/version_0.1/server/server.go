package main

import "sinx/simodel"

func main() {
	serber := simodel.NewSinxServer("version_0.1", 33366)
	serber.Server()
}
