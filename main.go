package main

import "io/ioutil"

func main() {
	data, err := ioutil.ReadFile("fixtures/medivh.json")
	check(err)
	auction.Parse(data)
}
