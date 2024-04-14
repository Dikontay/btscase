package forte

import (
	"errors"
	"fmt"
	"github.com/playwright-community/playwright-go"

	"github.com/Dikontay/btscase/internal/models"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ForteParser struct {
	Playwright  *playwright.Playwright
	BrowserType *playwright.Browser
	Page        playwright.Page
}

func ForteBankParser() (*ForteParser, error) {

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

	result := &ForteParser{Playwright: pw, BrowserType: &browser, Page: page}

	return result, nil
}

func (parser *ForteParser) ParseHalyk() error {
	url := "https://premier.forte.kz/ru/loyalty-program"

	urlsOfOffers, err := parser.GetDivsOfOffers(url)
	if err != nil {
		return err
	}

	//offers := make([]models.Offer, 0, len(urlsOfOffers)/
	offers, err := parser.GetOffersFromHalyk(urlsOfOffers)
	if err != nil {
		return err
	}
	for i := range offers {
		fmt.Println(offers[i])
	}

	return nil
}

func (parser *ForteParser) GetDivsOfOffers(URL string) ([]string, error) {

	// Navigate to the URL
	if _, err := parser.Page.Goto(URL); err != nil {
		return nil, fmt.Errorf("could not navigate to URL: %v", err)
	}

	// Wait for the page to load
	if err := parser.Page.WaitForLoadState(); err != nil {
		return nil, fmt.Errorf("could not wait for page load: %v", err)
	}

	// Find all <a> elements within the specified class
	divs, err := parser.Page.Locator("MuiGrid-root MuiGrid-item MuiGrid-grid-xs-12 MuiGrid-grid-md-4").All()
	if err != nil {
		return nil, fmt.Errorf("could not find links: %v", err)
	}

	var names []string
	
	for _, div := range divs {
		alt, err := div.GetAttribute("alt")
		if err != nil {

			log.Printf("could not get src attribute: %v", err)
			continue
		}
		if len(alt) == 0 {
			continue
		}
		alt = strings.TrimSpace(alt)

		names = append(names, alt)
	}

	return names, nil
}

func (parser *ForteParser) GetOffersFromHalyk(URLs []string) ([]models.Offer, error) {
	offers := make([]models.Offer, 0, len(URLs))

	for _, URL := range URLs {
		var offer models.Offer
		if _, err := parser.Page.Goto(URL); err != nil {
			log.Printf("could not navigate to URL %s: %v", URL, err)
			continue
		}
		if err := parser.Page.WaitForLoadState(); err != nil {
			log.Printf("could not wait for page load: %v", err)
			continue
		}
		title, err := parser.Page.Locator(`.text-2xl.mb-6.font-semibold`).InnerText()

		if err != nil {
			return nil, err
		}
		bonusAmount, err := getPercent(title)
		if err != nil {
			continue
		}
		offer.Precent = float64(bonusAmount)
		offer.Market = title
		offer.Bank = "Halyk"
		conditions, err := parser.Page.Locator(`body .wrapper .wrapper-in .content._js-bvi-filter .page.page--news .container .flex.flex-wrap.-gw .w-7\/12 .border-b.border-gray-100.pb-6.mb-6.content-inner-editer .mb-10`).AllInnerTexts()

		for j := range conditions {
			if conditions[j] == "" {
				continue
			}
			date, err := extractLastDateFromText(conditions[j])
			date = strings.TrimSpace(date)
			if err != nil {
				continue
			}
			bonusConditions, err := extractBonusConditions(conditions[j])
			if err != nil {
				continue

			}
			dateTime, err := time.Parse("02.01.2006", date)
			if err != nil {
				continue
			}
			limitations, err := getLimitations(conditions[j])

			offer.Due = dateTime
			offer.Condition = bonusConditions
			offer.Limitation = limitations

		}

		if err != nil {
			log.Printf("could not wait for page load: %v", err)
			continue
		}

		offers = append(offers, offer)

	}

	if err := (*parser.BrowserType).Close(); err != nil {
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

func extractBonusConditions(text string) (string, error) {
	// Compile the regular expression to match the bonus conditions section
	re, err := regexp.Compile(`(?m)Условия начисления бонуса:\s*\n(.*?)(?:\n\S|$)`)
	if err != nil {
		return "", err
	}

	// Find the match for the bonus conditions section
	match := re.FindStringSubmatch(text)
	if match == nil || len(match) < 2 {
		return "", fmt.Errorf("bonus conditions section not found")
	}

	// The first submatch contains the conditions text
	return match[1], nil
}

//	func findMarketName(text string) (string, error) {
//		re, err := regexp.Compile(`в магазин(е|ах) (\p{L}+)`)
//		if err != nil {
//			return "", err
//		}
//
//		matches := re.FindStringSubmatch(text)
//		if matches == nil || len(matches) < 2 {
//			return "", fmt.Errorf("store name not found")
//		}
//
//		return matches[1], nil
//	}
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

func getLimitations(text string) (string, error) {
	re, err := regexp.Compile(`(?m)Исключения по начислению бонусов:\s*\n(.*?)(?:\n\S|$)`)
	if err != nil {
		return "", err
	}

	// Find the match for the bonus conditions section
	match := re.FindStringSubmatch(text)
	if match == nil || len(match) < 2 {
		return "", fmt.Errorf("bonus conditions section not found")
	}

	// The first submatch contains the conditions text
	return match[1], nil
}
