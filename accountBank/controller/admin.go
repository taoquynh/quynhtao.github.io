package controller

import (
	"GoLearn/accountBank/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/xid"
	"log"
	"net/http"
)

// @Tags admin
// @Description Create Account
// @Param user body model.CreateAccount true "Tao tai khoan"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /bank/create-account [post]
func (c *Controller) CreateAccount(ctx *gin.Context) {
	var account model.CreateAccount
	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newAccount model.Account
	copier.Copy(&newAccount, &account)
	newAccount.Id = xid.New().String()
	err = c.DB.Insert(&newAccount)

	if err != nil {
		log.Println((err.Error()))
		model.NewError(ctx, http.StatusBadRequest, errors.New("Khong the tao tai khoan"))
		return
	}

	ctx.String(http.StatusOK, "Tao tai khoan thanh cong")
}

//
//// @Tags admin
//// @Description Get Amount
//// @Success 200 {string} string
//// @Failure 500 {object} model.HTTPError
//// @Router /bank/getAmount [get]
//func (c *Controller) GetAmount(ctx *gin.Context) {
//
//}
//
//
//// @Tags admin
//// @Description AddFund
//// @Success 200 {string} string
//// @Failure 500 {object} model.HTTPError
//// @Router /bank/addFund [post]
//func (c *Controller) AddFund(ctx *gin.Context) {
//
//}
//
//
//// @Tags admin
//// @Description withDrawMoney
//// @Success 200 {string} string
//// @Failure 500 {object} model.HTTPError
//// @Router /bank/withdrawMoney [get]
//func (c *Controller) WithDrawMoney(ctx *gin.Context) {
//
//}