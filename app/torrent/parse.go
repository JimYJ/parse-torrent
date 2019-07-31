package torrent

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/neoql/btlet/bencode"
)

type Torrent struct {
	Announce     string   `bencode:"announce"`
	AnnounceList []string `bencode:"announce-list"`
	Comment      string   `bencode:"comment"`
	CreationDate uint     `bencode:"creation date"`
	Encoding     string   `bencode:"encoding"`
	Info         Info     `bencode:"info"`
	Publisher    string   `bencode:"publisher"`
	PublisherURL string   `bencode:"publisher-url"`
}
type Files struct {
	Length uint     `bencode:"length"`
	Path   []string `bencode:"path"`
}
type Info struct {
	Length uint    `bencode:"length"`
	Files  []Files `bencode:"files"`
	Name   string  `bencode:"name"`
	//PieceLength int     `bencode:"piece length"`
	//Pieces      string  `bencode:"pieces"`
}

// Parse 解析种子文件
func Parse(b []byte) {
	var val Torrent
	f, err := os.Open("./111.torrent")
	if err != nil {
		log.Println(err)
		return
	}
	b, err = ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}

	// err = bencode.DecodeBytes(b, &torrent)
	err = bencode.Unmarshal(b, &val)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(val.Info.Length)
	// torrent, ok := val.(map[string]interface{})
	// if !ok {
	// 	log.Println(err)
	// 	return
	// }
	// if torrent["info"] == nil {
	// 	log.Println("cannot find info node!")
	// 	return
	// }
	// info, ok := val.(map[string]interface{})
	// if !ok {
	// 	log.Println(err)
	// 	return
	// }
	// if info["length"] != nil {

	// }
	// a, err := json.Marshal(val)
	// log.Println(string(a))
}
