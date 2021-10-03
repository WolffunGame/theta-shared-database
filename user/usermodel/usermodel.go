package usermodel

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (User) CollectionName() string {
	return "Users"
}

type User struct {
	mongodb.DefaultModel  `bson:",inline"`
	mongodb.DateFields    `bson:",inline"`
	Version               int             `json:"version" bson:"version"`
	Email                 string          `json:"email" bson:"email,omitempty"`
	UserName              string          `json:"username" bson:"username,omitempty"`
	NumChangeName         int             `json:"numChangeName" bson:"numChangeName"`
	Status                UserStatus      `json:"status" bson:"status,omitempty"`
	Address               string          `json:"address" bson:"address,omitempty"`
	Avatar                string          `json:"avatar" bson:"avatar,omitempty"`
	Bio                   string          `json:"bio" bson:"bio,omitempty"`
	Cover                 string          `json:"cover" bson:"cover,omitempty"`
	EliteUser             bool            `json:"eliteUser" bson:"eliteUser,omitempty"`
	TotalItems            int32           `json:"totalItems" bson:"totalItems,omitempty"`
	TotalLiveAuctionItems int32           `json:"totalLiveAuctionItems" bson:"totalLiveAuctionItems,omitempty"`
	TotalOnSaleItems      int32           `json:"totalOnSaleItems" bson:"totalOnSaleItems,omitempty"`
	TotalSoldItems        int32           `json:"totalSoldItems" bson:"totalSoldItems,omitempty"`
	Nonce                 int             `json:"nonce" bson:"nonce"`
	IsClaimedFreeHero     bool            `json:"canClaimFreeHero" bson:"canClaimFreeHero"`
	UserProfile           UserProfile     `json:"userProfile" bson:"userProfile"`
	PlayerStatistic       PlayerStatistic `json:"playerStatistic" bson:"playerStatistic"`
	Suspicious            int             `json:"-" bson:"suspicious"`
	SuspiciousWrongData   int             `json:"-" bson:"suspiciousWrongData"`
}

type PlayerStatistic struct {
	Battle    int32 `json:"battle" bson:"battle"`
	Victory   int32 `json:"victory" bson:"victory"`
	Streak    int32 `json:"streak" bson:"streak"`
	CurStreak int32 `json:"-" bson:"curStreak"`
	Triple    int32 `json:"triple" bson:"triple"`
	Mega      int32 `json:"mega" bson:"mega"`
	Mvp       int32 `json:"mvp" bson:"mvp"`
	Hero      int32 `json:"hero" bson:"hero"`
}

func (user *User) GetUserId() string {
	return user.ID.(primitive.ObjectID).Hex()
}

type UserStatus int

const (
	ACTIVE UserStatus = 1
	BANNED UserStatus = -1
)

type UserProfile struct {
	Level      int `bson:"level" json:"level"`
	XP         int `bson:"xp" json:"xp"`
	LevelUpGPP int `bson:"levelUpGPP" json:"levelUpGPP"`
}