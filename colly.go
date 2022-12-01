package main

import (
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/storage"
)

func NewColler(s storage.Storage) *colly.Collector {
	c := colly.NewCollector(
		colly.Async(Async),
		colly.MaxDepth(MaxDepth),
		colly.UserAgent(UserAgent),
	)

	c.Limit(&colly.LimitRule{Parallelism: Limit})
	c.WithTransport(&http.Transport{DisableKeepAlives: DisableKeepAlives})
	c.SetStorage(s)
	return c
}
