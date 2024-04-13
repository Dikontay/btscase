package halyk

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
)

func ParseHalyk() {
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
	if _, err = page.Goto("https://halykbank.kz/promo"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	if err := page.WaitForLoadState(); err != nil {
		log.Fatalf("could not wait for load: %v", err)
	}

	// -------------------------------------------------------------------------

	items, err := page.Locator(".flex.flex-col.flex-3.justify-between.px-6.pb-6").All()
	if err != nil {
		log.Fatalf("could not get all items: %v", err)
	}

	for i, item := range items {
		title, err := item.Locator(".mb-4").TextContent()
		if err != nil {
			log.Fatalf("could not get item's title: %v", err)
		}
		// description, err := item.Locator(".sc-46d3607b-0.bGEJem").TextContent()
		// if err != nil {
		// 	log.Fatalf("could not get item's description: %v", err)
		// }
		//price, err := item.Locator(".sc-de642809-3.hwSkqq").TextContent()
		//if err != nil {
		//	log.Fatalf("could not get item's price: %v", err)
		//}

		ans := fmt.Sprintf("Item %d: %s", i, title)
		fmt.Println(ans)
	}

	// -------------------------------------------------------------------------

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
