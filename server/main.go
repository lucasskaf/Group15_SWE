package main

//imports commented out to avoid generating errors for unused

import (
	"fmt"
	"gorm.io/gorm"
	/*
			"gorm.io/driver/sqlite"
			"net/http"
		    "github.com/gin-gonic/gin"
	*/)

type movie struct {
	title   string
	runtime float32
}

type user struct {
	gorm.Model
	username  string
	password  string
	watchlist []movie
	genres    []string
	rating    float32
	providers []string
}

func main() {
	fmt.Println("test")
}
