package parser

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
)

type Parser struct {
	Playwright   *playwright.Playwright
	BrowtherType *playwright.Browser
	Page         playwright.Page
}

func NewParser() (*Parser, error) {

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		return nil, err
	}
	page, err := browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create new page: %v", err)
	}

	result := &Parser{Playwright: pw, BrowtherType: &browser, Page: page}

	return result, nil
}

func (p *Parser) Close() error {
	if err := (*p.BrowtherType).Close(); err != nil {
		return fmt.Errorf("could not close browser: %v", err)
	}
	if err := p.Playwright.Stop(); err != nil {
		return fmt.Errorf("could not stop Playwright: %v", err)
	}
	return nil
}
