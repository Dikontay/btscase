package transport

import (
	"errors"

	"github.com/Dikontay/btscase/internal/models"
	"github.com/gin-gonic/gin"
)

func GetUserFromGin(c *gin.Context) (*models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, errors.New("not Authorized")
	}

	userModel, ok := user.(*models.User)
	if !ok {
		return nil, errors.New("Forbiden")
	}

	return userModel, nil
}

func (t *Transport) GetOffersHandler(c *gin.Context) {
	market := c.Param("market")
	offers, err := t.Service.GetOffersByMarket(c.Request.Context(), market)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, offers)
}

func (t *Transport) GetAllOffers(c *gin.Context) {

}
