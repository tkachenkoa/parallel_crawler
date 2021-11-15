package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSuffix (t *testing.T) {
	link := "https://passport.yandex.ru"
	main_url := "yandex.ru"
	if strings.HasSuffix(link, main_url) {
		fmt.Printf("Link %s has suffix %s\n", link, main_url)
	} else {
		t.Errorf("Link %s does not have suffix %s\n", link, main_url)
	}
}

func TestSubdomain (t *testing.T) {
	link := "https://passport.yandex.ru/subdir"
	main_url := "yandex.ru"
	if strings.Contains(link, main_url) {
		fmt.Printf("Link %s is subdomain of %s\n", link, main_url)
	} else {
		t.Errorf("Link %s is not a subdomain of %s\n", link, main_url)
	}
}
