package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	Id                   primitive.ObjectID `bson:"_id" json:"_id"`
	PlayerNumber         int                `bson:"playerNumber" json:"playerNumber"`
	PlayerName           string             `bson:"playerName" json:"playerName"`
	Country              string             `bson:"country" json:"country"`
	PlayerType           string             `bson:"playerType" json:"playerType"`
	IplTeam              string             `bson:"iplTeam" json:"iplTeam"`
	PrevTeam             string             `bson:"prevTeam" json:"prevTeam"`
	CurrentTeam          string             `bson:"currentTeam" json:"currentTeam"`
	BasePrice            float64            `bson:"basePrice" json:"basePrice"`
	PrevFantasyPoints    int                `bson:"prevFantasyPoints" json:"prevFantasyPoints"`
	CurrentFantasyPoints int                `bson:"currentFantasyPoints" json:"currentFantasyPoints"`
	SellingPrice         float64            `bson:"sellingPrice" json:"sellingPrice"`
	Match                primitive.ObjectID `bson:"match" json:"match"`
	Sold                 bool               `bson:"sold" json:"sold"`
	Unsold               bool               `bson:"unsold" json:"unsold"`
}
