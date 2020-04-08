package email

import "os"

var (
	addr = "smtp.gmail.com"
	pw   = os.Getenv("OLMAX_EMAIL")
	gmail = "olmaxmedical@gmail.com"
	url  = "https://medical.olmax.io"
)
