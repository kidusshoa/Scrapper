package main

import (
	"github.com/gocolly/colly"
)

func main(){
	c := colly.NewCollector(colly.AllowedDomains("www.mydomain.com"))


	c.Visit("www.mydomain.com/some/fany/url")
}