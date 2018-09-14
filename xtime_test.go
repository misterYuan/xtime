package xtime

import (
	"log"

	"testing"
)

func TestOption(t *testing.T) {
	p := NewStampParser([]string{"20060102", "2006"})
	s, err := p.Parse("201805071")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}
