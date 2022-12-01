package main

var EtherscanBaseUrl = "https://etherscan.io/token/"
var OpenseaBaseUrl = "https://opensea.io/zh-CN/rankings?sortBy=total_volume"

var (
	Async             = true
	MaxDepth          = 1
	Limit             = 3
	DisableKeepAlives = true
	UserAgent         = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
)

var (
	Database = "opensea"
	Ip       = "127.0.0.1"
	Port     = 27017
)
