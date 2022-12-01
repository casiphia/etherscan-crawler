package main

var (
	EmailKeyword    = "Email: "
	TwitterKeyword  = "Twitter: "
	LinkedinKeyword = "Linkedin: "
	DiscordKeyword  = "Discord: "
	OpenseaKeyword  = "Opensea: "
)

type ContactDetails struct {
	Email    string
	Twitter  string
	Linkedin string
	Discord  string
	Opensea  string
}
