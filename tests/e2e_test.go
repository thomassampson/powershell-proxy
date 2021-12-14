package main

import (
	"log"
	"testing"

	"github.com/mxschmitt/playwright-go"
)

func Test_End2End(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://one.acuityads.cloud/activate"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	log.Print(page.Content())

	//assert.Nil(t, page.Fill("username", "1234"))
	//assert.Nil(t, page.Type("input.user-code", "123456"))

}
