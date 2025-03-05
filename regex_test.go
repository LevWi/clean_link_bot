package main

import (
	"testing"
)

const examples string = "https://youtube.com/shorts/BlablaFoo?si=ivmftR04ZyaOI1HR\n" +
	"https://www.youtube.com/watch?v=BlablaFoo&feature=featured \n" +
	"https://m.youtube.com/watch?v=BlablaFoo "

func TestRegex(t *testing.T) {
	re := createRegexObj()

	res := re.FindAllStringSubmatch(examples, -1)

	if res == nil {
		t.Fatal("No matches found")
	}

	if len(res) != 3 {
		t.Fatalf("Wrong result len = %v", len(res))
	}

	for _, arr := range res {
		if arr[6] != "BlablaFoo" {
			t.Fatal("Wrong match")
		}
	}
}
