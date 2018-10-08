package main

import (
	gin "github.com/guizot/go-gin-mongodb/config/server"
)

func main() {
	var s gin.Server
	s.StartGin()
}

/*
	Karena variable s memiliki tipe gin.Server milik server.go
	maka s dapat menjalan kan fungsi StartGin
*/
