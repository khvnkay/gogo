package controllers

import (
	"errors"
)

var (
	ErrCantFindProduct   = errors.New("")
	ErrCantDecodeProduct = errors.New("")
	ErrUserIdIsNotValid  = errors.New("")
	ErrCantUpdateUser    = errors.New("")
	ErrCantRemove        = errors.New("")
	ErrCantGetItem       = errors.New("")
	ErrCantBuyCartItem   = errors.New("")
)

func AddProductToCart() {

}

func RemoveItemFromCart() {

}

func BuyItemFromCart() {

}

func InstantBuy() {

}
