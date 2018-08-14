package main

import (
	"fmt"
	"github.com/novalagung/gubrak"
	"io/ioutil"
)

type Sample struct {
	EbookName      string
	DailyDownloads int
}

func hashArray(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	return content, err
}

func main() {
	c, err := hashArray("../mongo/m.go")
	if err != nil {
		return
	}
	data := []Sample{
		{EbookName: "clean code", DailyDownloads: 10000},
		{EbookName: "rework", DailyDownloads: 12000},
		{EbookName: "detective comics", DailyDownloads: 11500},
	}

	result, err := gubrak.Filter(data, func(each Sample) bool {
		return each.DailyDownloads > 11000
	})
	fmt.Printf("%s %s", c, result)
	data2 := []string{"abc", "efg", "hij"}
	err2 := gubrak.ForEach(data2, func(each string) {
		fmt.Println(each)
	})
	if err != nil {
		fmt.Printf("%s", err2)
	}

}
