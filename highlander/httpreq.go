package main

import (
    "time"
    "github.com/andelf/go-curl"
)

func NewReq(){
	rv := new(Req)
	return rv
}

func req() {
    
    ch1 := curl.EasyInit()
	ch2 := curl.EasyInit()

	ch1.Setopt(curl.OPT_URL, "http://www.shutterstock.com")
	ch1.Setopt(curl.OPT_HEADER, 0)
	ch1.Setopt(curl.OPT_VERBOSE, true)
	ch2.Setopt(curl.OPT_URL, "http://www.cnn.com")
	ch2.Setopt(curl.OPT_HEADER, 0)
	ch2.Setopt(curl.OPT_VERBOSE, true)

	mh := curl.MultiInit()

	mh.AddHandle(ch1)
	mh.AddHandle(ch2)

	for {
		nRunning, _ := mh.Perform()
		// println("n =", nRunning)
		if nRunning == 0 {
			println("ok")
			break
		}
		time.Sleep(1000)
	}

}