package main

import (
	"log"
	"parse-torrent/app/torrent"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	torrent.Parse(nil)
}
