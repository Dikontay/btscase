package halyk

import (
	"errors"
	"fmt"
	"github.com/Dikontay/btscase/internal/models"
	"github.com/playwright-community/playwright-go"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type HalykParser struct {
	Playwright   *playwright.Playwright
	BrowtherType *playwright.Browser
	Page         playwright.Page
}

func NewHalykParser() (*HalykParser, error) {

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

	result := &HalykParser{Playwright: pw, BrowtherType: &browser, Page: page}

	return result, nil
}

func (parser *HalykParser) ParseHalyk() error {
	url := "https://halykbank.kz/promo"

	urlsOfOffers, err := parser.GetURLsOfOffers(url)
	if err != nil {
		return err
	}

	//offers := make([]models.Offer, 0, len(urlsOfOffers)/
	_, err = parser.GetOffersFromHalyk(urlsOfOffers)
	if err != nil {
		return err
	}

	return nil
}

func (parser *HalykParser) GetURLsOfOffers(URL string) ([]string, error) {

	// Navigate to the URL
	if _, err := parser.Page.Goto(URL); err != nil {
		return nil, fmt.Errorf("could not navigate to URL: %v", err)
	}

	// Wait for the page to load
	if err := parser.Page.WaitForLoadState(); err != nil {
		return nil, fmt.Errorf("could not wait for page load: %v", err)
	}

	// Find all <a> elements within the specified class
	links, err := parser.Page.Locator(".h-full.flex.flex-col.rounded-xl.border.border-gray-100.bg-gray-50").All()
	if err != nil {
		return nil, fmt.Errorf("could not find links: %v", err)
	}

	var URLs []string

	for _, link := range links {
		href, err := link.GetAttribute("href")
		if err != nil {

			log.Printf("could not get href attribute: %v", err)
			continue
		}
		if len(href) == 0 {
			continue
		}
		href = strings.TrimSpace(href)

		URLs = append(URLs, href)
	}

	return URLs, nil
}

func (parser *HalykParser) GetOffersFromHalyk(URLs []string) ([]models.Offer, error) {
	offers := make([]models.Offer, 0, len(URLs))

	for i, URL := range URLs {
		if _, err := parser.Page.Goto(URL); err != nil {
			log.Printf("could not navigate to URL %s: %v", URL, err)
			continue
		}
		if err := parser.Page.WaitForLoadState(); err != nil {
			log.Printf("could not wait for page load: %v", err)
			continue
		}

		conditions, err := parser.Page.Locator(`body .wrapper .wrapper-in .content._js-bvi-filter .page.page--news .container .flex.flex-wrap.-gw .w-7\/12 .border-b.border-gray-100.pb-6.mb-6.content-inner-editer .mb-10`).AllTextContents()
		for i := range conditions {
			_, err := extractLastDateFromText(conditions[i])
			if err != nil {
				return nil, err
			}
			//fmt.Println(date)

		}

		if err != nil {
			log.Printf("could not wait for page load: %v", err)
			continue
		}
		if i == 0 {
			break
		}

	}

	if err := (*parser.BrowtherType).Close(); err != nil {
		return nil, fmt.Errorf("could not close browser: %v", err)
	}
	if err := parser.Playwright.Stop(); err != nil {
		return nil, fmt.Errorf("could not stop Playwright: %v", err)
	}
	return offers, nil

}

func getPercent(text string) (int, error) {
	re := regexp.MustCompile(`(\d+)%\s*бонусов`)
	match := re.FindStringSubmatch(text)
	if match != nil && len(match) > 1 {
		percentage, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		return percentage, nil
	}

	return 0, errors.New("cannot get percentage of offer")
}

func extractLastDateFromText(text string) (string, error) {
	// First, extract the date range text from the HTML using regex
	re, err := regexp.Compile(`\b(\d{2}\.\d{2}\.\d{4})\b`)
	if err != nil {
		return "", err
	}

	dates := re.FindAllString(text, -1)
	if len(dates) == 0 {
		return "", fmt.Errorf("no dates found in the string")
	}

	// Return the last date from the matches.
	lastDate := dates[len(dates)-1]
	return lastDate, nil
}
