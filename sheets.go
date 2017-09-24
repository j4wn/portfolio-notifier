package main

import (
	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var (
	value        float64
	renderOption = "UNFORMATTED_VALUE"
	sheetsScope  = "https://www.googleapis.com/auth/spreadsheets.readonly"
)

func portfolioValue(spreadsheetID string, readRange string) float64 {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()

	mailLog := log.WithFields(log.Fields{
		"file": "sheets.go",
	})

	client, err := google.DefaultClient(ctx, sheetsScope)
	if err != nil {
		mailLog.Fatalf("Could not build client: %v", err)
	}

	mailLog.Infof("Sheets client: %s", sheetsScope)
	service, err := sheets.New(client)
	if err != nil {
		mailLog.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	// Get the unformatted value of the range
	mailLog.Infof("Spreadsheets.Values.Get: %s, %s, %s", renderOption, spreadsheetID, readRange)
	resp, err := service.Spreadsheets.Values.Get(spreadsheetID, readRange).ValueRenderOption(renderOption).Do()
	if err != nil {
		mailLog.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	if len(resp.Values) > 0 {
		// return the contents of the first (only) cell
		// I have no idea why I had to do this
		if v, ok := resp.Values[0][0].(float64); ok {
			value = v
		} else {
			mailLog.Fatalf("Could not return cell contents")
		}
	}
	return value
}
