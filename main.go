package main

import (
	"fmt"
	"sync"
)

// map of long url to list of clientIds
var UrlToClientMap map[string][]string

//map of short url to long url
var UrlMap map[string]string

// map of client id to list of longurl and shorturl pair
var ClientDataMap map[string]map[string]string

// Key genration id
var id int64

// Init function
func Init() {
	UrlToClientMap = make(map[string][]string, 0)
	UrlMap = make(map[string]string, 0)
	ClientDataMap = make(map[string]map[string]string, 0)

	id = 0
}

// For each new request generate a new id
// locking mechansim is applied
func GetNewlId() int64 {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	id++
	mutex.Unlock()
	return id
}

// convert an unique id to Base 64
func ConvertToBase64(id int64) string {

	var shortUrl string
	charMap := string("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/")
	for id > 0 {
		shortUrl = string(charMap[id%64]) + shortUrl
		id /= 64
	}
	return shortUrl
}

// Returns a new sjort url
func GetNewUrlId() string {
	return ConvertToBase64(GetNewlId())
}

// Generate a short url and store all relevant information in map
func GetShortUrl(longUrl string, clientId string) string {
	var shorturl string

	// if there is no client generated short url against current longurl , genrate a new one
	clientList, ok := UrlToClientMap[longUrl]
	if !ok {
		shorturl = GetNewUrlId()
		UrlMap[shorturl] = longUrl
		UrlToClientMap[longUrl] = append(UrlToClientMap[longUrl], clientId)
		_, isPresent := ClientDataMap[clientId]
		if !isPresent {
			ClientDataMap[clientId] = make(map[string]string, 0)
		}
		ClientDataMap[clientId][longUrl] = shorturl
	}

	// if there is a client which already generated a short url against current long url return the previous short url
	for _, id := range clientList {
		if string(id) == clientId {
			return ClientDataMap[clientId][longUrl]
		}
	}
	// if control reaches here , it means short url is never generated for any client against the current long url
	shorturl = GetNewUrlId()
	UrlMap[shorturl] = longUrl
	UrlToClientMap[longUrl] = append(UrlToClientMap[longUrl], clientId)
	ClientDataMap[clientId][longUrl] = shorturl
	return shorturl
}

func main() {
	Init()
	// generate short url for same client having different url
	//Test1()

	// generate short url for different client having different url
	//Test2()

	// generate short url for same client having same url
	//Test3()

	// generate short url for different client having same url
	//Test4()
}

func Test1() {
	var shortUrl string
	for i := 0; i < 5; i++ {
		shortUrl = GetShortUrl("bsdnbdsjfbjsdfba"+string(i), "hjhd")
		fmt.Println("short url is :", shortUrl)
		fmt.Println("-----------")
	}
}

func Test2() {
	var shortUrl string
	for i := 0; i < 5; i++ {
		shortUrl = GetShortUrl("bsdnbdsjfbjsdfba"+string(i), "hjhd"+string(i))
		fmt.Println("short url is :", shortUrl)
		fmt.Println("-----------")
	}
}
func Test3() {
	var shortUrl string
	for i := 0; i < 5; i++ {
		shortUrl = GetShortUrl("bsdnbdsjfbjsdfba", "hjhd")
		fmt.Println("short url is :", shortUrl)
		fmt.Println("-----------")
	}
}
func Test4() {
	var shortUrl string
	for i := 0; i < 5; i++ {
		shortUrl = GetShortUrl("bsdnbdsjfbjsdfba"+string(i), "hjhd"+string(i))
		fmt.Println("short url is :", shortUrl)
		fmt.Println("-----------")
	}
}
