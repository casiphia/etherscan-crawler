package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func GetContactByContractAddress(c *colly.Collector, ContractAddress string) *ContactDetails {
	contactDetails := new(ContactDetails)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[data-original-title]", func(e *colly.HTMLElement) {
		htmlSrc := e.Attr("data-original-title")
		if strings.HasPrefix(htmlSrc, EmailKeyword) {
			contactDetails.Email = strings.Trim(htmlSrc, EmailKeyword)
		}
		if strings.HasPrefix(htmlSrc, TwitterKeyword) {
			contactDetails.Twitter = strings.Trim(htmlSrc, TwitterKeyword)
		}
		if strings.HasPrefix(htmlSrc, LinkedinKeyword) {
			contactDetails.Linkedin = strings.Trim(htmlSrc, LinkedinKeyword)
		}
		if strings.HasPrefix(htmlSrc, DiscordKeyword) {
			contactDetails.Discord = strings.Trim(htmlSrc, DiscordKeyword)
		}
		if strings.HasPrefix(htmlSrc, OpenseaKeyword) {
			contactDetails.Opensea = strings.Trim(htmlSrc, OpenseaKeyword)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Finished", r.Request.URL)
	})

	c.Visit(EtherscanBaseUrl + ContractAddress)
	c.Wait()
	return contactDetails
}
