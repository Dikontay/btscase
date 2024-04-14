package transport

import "github.com/gin-gonic/gin"

func (t *Transport) GetOffersHandler(c *gin.Context) {
	market := c.Param("market")
	offers, err := t.service.GetOffersByMarket(c.Request.Context(), market)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, offers)
}

func (t *Transport) GetAllOffers(c *gin.Context) {

}
