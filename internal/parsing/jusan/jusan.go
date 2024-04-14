package jusan

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Dikontay/btscase/internal/models"
	"github.com/Dikontay/btscase/internal/parsing/parser"
)

// ParseJusan orchestrates the fetching and parsing of offer data.
func ParseJusan() error {
	pars, err := parser.NewParser()
	if err != nil {
		return err
	}
	defer pars.Close()

	urls, err := getUrlsOfOffers("https://jmart.kz/store", pars)
	if err != nil {
		return err
	}

	offers, err := fetchOffers(urls, pars)
	if err != nil {
		return err
	}
	bonuses, err := getBonuses("https://jmart.kz/store", pars)
	for i, offer := range offers {
		offers[i].Precent = bonuses[i]
		fmt.Println(offer)
	}
	return nil
}

// getUrlsOfOffers fetches URLs from the provided base URL using the parser.
func getUrlsOfOffers(url string, parser *parser.Parser) ([]string, error) {
	if _, err := parser.Page.Goto(url); err != nil {
		return nil, fmt.Errorf("could not navigate to URL: %v", err)
	}

	links, err := parser.Page.Locator(".StyledBox--1sqduml.cLycsw a").All()
	if err != nil {
		return nil, fmt.Errorf("could not find links: %v", err)
	}
	URLs := make([]string, 0, len(links))
	for _, link := range links {
		href, err := link.GetAttribute("href")
		if err != nil {
			continue
		}
		href = strings.TrimSpace(href)
		if href != "" {
			URLs = append(URLs, href)
		}
	}
	return URLs, nil
}

func getBonuses(url string, p *parser.Parser) ([]float64, error) {
	if _, err := p.Page.Goto(url); err != nil {
		return nil, fmt.Errorf("could not navigate to URL: %v", err)
	}

	bonuses, err := p.Page.Locator(".StyledText--v46vlv.fkdwMt.StyledProductLabel-sc-ig5wgt-1.iOnNIZ").AllInnerTexts()

	result := make([]float64, 0, len(bonuses))

	for i := range bonuses {
		result[i], err = strconv.ParseFloat(bonuses[i][0:len(bonuses)-1], 64)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// fetchOffers retrieves offers from the list of URLs.
func fetchOffers(URLs []string, parser *parser.Parser) ([]models.Offer, error) {
	var (
		offers []models.Offer
		wg     sync.WaitGroup
		mu     sync.Mutex
		errors []error
	)

	for _, url := range URLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			offer, err := fetchOffer(url, parser)
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
				return
			}
			mu.Lock()
			offers = append(offers, offer)
			mu.Unlock()
		}("https://jmart.kz" + url)
	}

	wg.Wait()
	if len(errors) > 0 {
		return nil, fmt.Errorf("errors encountered: %v", errors)
	}
	return offers, nil
}

// fetchOffer fetches a single offer from a URL.
func fetchOffer(url string, parser *parser.Parser) (models.Offer, error) {
	var maxRetries = 3
	var retryDelay time.Duration = 1 * time.Second // Start with 1 second

	for attempt := 0; attempt < maxRetries; attempt++ {
		_, err := parser.Page.Goto(url, playwright.PageGotoOptions{
			Timeout: playwright.Float(60000), // 60 seconds timeout
		})
		if err == nil {
			if err := parser.Page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
				State: playwright.LoadStateLoad,
			}); err == nil {
				name, err := parser.Page.Locator(`.StyledText--v46vlv eTaBJM`).InnerText()
				condition, err := parser.Page.Locator(`.StyledText--v46vlv.cAwPMV`).InnerText()
				if err == nil {
					return models.Offer{Market: name, Condition: condition}, nil
				}
				return models.Offer{}, err
			}
		}
		if attempt < maxRetries-1 {
			time.Sleep(retryDelay)
			retryDelay *= 2 // Double the delay for the next retry
		}
		fmt.Printf("Retry %d for URL %s due to error: %v\n", attempt+1, url, err)
	}
	return models.Offer{}, fmt.Errorf("could not navigate to URL after %d attempts", maxRetries)
}
