package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID   `bson:"_id" json:"_id"`
	Username     string               `bson:"username" json:"username"`
	Email        string               `bson:"email" json:"email"`
	ImgUrl       string               `bson:"ImgUrl" json:"ImgUrl"`
	Teamname     string               `bson:"teamname" json:"teamname"`
	Squad        []primitive.ObjectID `bson:"squad" json:"squad"`
	IsPlaying    bool                 `bson:"isPlaying" json:"isPlaying"`
	IsAuctioneer bool                 `bson:"isAuctioneer" json:"isAuctioneer"`
}
