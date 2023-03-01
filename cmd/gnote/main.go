// cmd/gnote/main.go
package main

import (
	app_gnote "github.com/LiaoSirui/GNote/internal/app/gnote"
)

func main() {
	// app entry
	app_gnote.StartServer()
}
