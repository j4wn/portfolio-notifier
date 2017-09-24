package main

import (
	"flag"

	"github.com/leekchan/accounting"
)

func main() {

	// gmail flags
	sender := flag.String("from", "me@example.com", "From address")
	recipient := flag.String("to", "you@example.com", "Recipient of the email")

	// sheets flags
	// https://docs.google.com/spreadsheets/d/1lTQNpSixfQwS3y6XrTc-GH9mJQKpf82toB3yjuUNCQQ/edit
	sheet := flag.String("sheet", "1lTQNpSixfQwS3y6XrTc-GH9mJQKpf82toB3yjuUNCQQ", "Google Sheet to search")
	cell := flag.String("cell", "Main Portfolio!P8", "Sheet Name!Cell containing the value")
	flag.Parse()

	// fetch the float64(value)
	rawValue := portfolioValue(*sheet, *cell)

	// format the raw value
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	value := ac.FormatMoney(rawValue)

	// send some mail
	sendMail(*sender, *recipient, value)
}
