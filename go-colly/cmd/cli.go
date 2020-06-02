package main

import (
	"github.com/rikisan1993/go-web-scrapper/internal/pkg/coronastats"
	"github.com/rikisan1993/go-web-scrapper/internal/pkg/firebaseclient"
)

func main() {
	app := firebaseclient.CreateApp()
	coronastats.Run()
}