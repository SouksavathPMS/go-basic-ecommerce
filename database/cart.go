package database

import "errors"

var (
	ErrCanNotFindProduct    = errors.New("can't find the product")
	ErrCanNotDecodeProduct  = errors.New("can't find the product")
	ErrUserIDIsNotValid     = errors.New("this user credential is not valid")
	ErrCanNotUpdateUser     = errors.New("can't add the product to the cart")
	ErrCanNotRemoveItemCart = errors.New("can't remove this item from the cart")
	ErrCanGetItem           = errors.New("unable to get the item from cart")
	ErrCanNotBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuy() {

}
