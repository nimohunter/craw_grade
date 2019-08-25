package facerank

import (
	"fmt"
	"testing"
)

func TestGetFaceRank(t *testing.T) {
	rank := GetFaceRank("https://photo.zastatic.com/images/photo/419564/1678252745/3165804982886731.jpg")
	fmt.Println(rank)
}
