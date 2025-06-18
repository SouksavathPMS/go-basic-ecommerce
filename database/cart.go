package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	ErrCanNotFindProduct    = errors.New("can't find the product")
	ErrCanNotDecodeProduct  = errors.New("can't find the product")
	ErrUserIDIsNotValid     = errors.New("this user credential is not valid")
	ErrCanNotUpdateUser     = errors.New("can't add the product to the cart")
	ErrCanNotRemoveItemCart = errors.New("can't remove this item from the cart")
	ErrCanGetItem           = errors.New("unable to get the item from cart")
	ErrCanNotBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart(ctx context.Context,
	productCollection,
	userCollection *mongo.Collection,
	productID primitive.ObjectID,
	userID string,
) error {
	return errors.New("")
}

func RemoveCartItem(ctx context.Context,
	productCollection,
	userCollection *mongo.Collection,
	productID primitive.ObjectID,
	userID string,
) error {
	return errors.New("")
}

func BuyItemFromCart(ctx context.Context,
	productCollection,
	userCollection *mongo.Collection,
	userID string,
) error {
	return errors.New("")
}

func InstantBuy(ctx context.Context,
	productCollection,
	userCollection *mongo.Collection,
	productID primitive.ObjectID,
	userID string,
) error {
	return errors.New("")
}
