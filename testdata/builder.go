package main

import (
	"fmt"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func main() {

	err := playwright.Install()
	if err != nil {
		log.Fatalf("Unable to install playwright")
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	})
	if err != nil {
		log.Fatalf("Context: %v", err)
	}

	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	log.Printf("Navigate to %v", os.Args[1])
	if _, err = page.Goto(os.Args[1], playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	log.Printf("Navigating")
	herotitle := page.Locator("body")
	text, err := herotitle.TextContent()
	if err != nil {
		log.Fatalf("Bad: %v", err)
	}
	fmt.Print("%v\n", text)

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
