package transport

import (
	"errors"
	"net/http"
	"time"

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
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(405, gin.H{"error": err.Error()})
	}
	market := c.Param("market")
	offers, err := t.Service.GetOffersByMarket(c.Request.Context(), market, user.ID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, offers)
}

func (t *Transport) AddCardHandler(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(405, gin.H{"error": err.Error()})
	}
	var input struct {
		Bank  string    `json:"bank"`
		Type  string    `json:"type"`
		Nomer string    `json:"nomer"`
		Due   time.Time `json:"due"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	card := &models.Card{
		Bank:   input.Bank,
		Type:   input.Type,
		Nomer:  input.Nomer,
		Due:    input.Due,
		UserID: user.ID,
	}
	if err := t.Service.AddCard(c.Request.Context(), card); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nil)
}

func (t *Transport) GetUserInfo(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err = t.Service.GetUserInfo(c.Request.Context(), user.ID)
	c.JSON(http.StatusOK, user)

}

func (t *Transport) GetCard(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err = t.Service.GetUserInfo(c.Request.Context(), user.ID)
	c.JSON(http.StatusOK, user)

}

func (t *Transport) GetCardByUserId(c *gin.Context) {
	user, err := GetUserFromGin(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	card, err := t.Service.GetCardByUserID(c.Request.Context(), user.ID)
	c.JSON(http.StatusOK, card)
}
