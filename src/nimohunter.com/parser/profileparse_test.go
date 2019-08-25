package parser

import (
	"fmt"
	"nimohunter.com/fetcher"
	"testing"
)

func TestPraseProfile(t *testing.T) {
	url := "https://album.zhenai.com/u/1678252745"
	bytes, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("occur error:", err)
		return
	}
	result := ProfileParser(bytes)
	fmt.Printf("%v", result)
	t.Error("ddd")
}
