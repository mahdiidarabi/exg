package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string
	Email       string `gorm:"unique"`
	Phone       string
	Password    string
	BtcBalance  float64 `gorm:"default:0"`
	EthBalance  float64 `gorm:"default:0"`
	DashBalance float64 `gorm:"default:0"`
	TethBalance float64 `gorm:"default:0"`
	XrpBalance  float64 `gorm:"default:0"`
	BinBalance  float64 `gorm:"default:0"`
	EosBalance  float64 `gorm:"default:0"`
	// Coins [7]float64 `gorm:"default:0"`
}
