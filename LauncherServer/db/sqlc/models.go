// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type PendingTransaction struct {
	BalanceID int32     `json:"balance_id"`
	UserID    int32     `json:"user_id"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	UpdateAt  time.Time `json:"update_at"`
	ID        int32     `json:"id"`
}

type Server struct {
	CategoryID  int32     `json:"category_id"`
	ServerName  string    `json:"server_name"`
	MiniDesc    string    `json:"mini_desc"`
	Description string    `json:"description"`
	ListMods    string    `json:"list_mods"`
	Link        string    `json:"link"`
	Ispvp       bool      `json:"ispvp"`
	Worlds      string    `json:"worlds"`
	CreatedAt   time.Time `json:"created_at"`
	LastWipe    time.Time `json:"last_wipe"`
	ID          int32     `json:"id"`
}

type ServersCategory struct {
	Version string `json:"version"`
	ID      int32  `json:"id"`
}

type ServersDatum struct {
	CategoryID int32 `json:"category_id"`
	ServerIp   int32 `json:"server_ip"`
	ServerPort int32 `json:"server_port"`
	ID         int32 `json:"id"`
}

type ServersImage struct {
	CategoryID int32  `json:"category_id"`
	ImageLink  string `json:"image_link"`
	ImageOne   string `json:"image_one"`
	ImageTwo   string `json:"image_two"`
	ID         int32  `json:"id"`
}

type ServersRcon struct {
	CategoryID   int32  `json:"category_id"`
	RconIp       string `json:"rcon_ip"`
	RconPort     int32  `json:"rcon_port"`
	RconPassword string `json:"rcon_password"`
	ID           int32  `json:"id"`
}

type User struct {
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Isadmin       bool      `json:"isadmin"`
	Ip            string    `json:"ip"`
	Uuid          string    `json:"uuid"`
	Accesstoken   string    `json:"accesstoken"`
	Serverid      string    `json:"serverid"`
	RegData       time.Time `json:"reg_data"`
	LastLoginSite time.Time `json:"last_login_site"`
	ID            int32     `json:"id"`
}

type User2fa struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
	_2fa   bool  `json:"2fa"`
}

type UserEconomy struct {
	UserID   int32 `json:"user_id"`
	Balance  int32 `json:"balance"`
	Isfreeze int32 `json:"isfreeze"`
	ID       int32 `json:"id"`
}

type UserSession struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	Isblocked    int32     `json:"isblocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserVote struct {
	ID            int32     `json:"id"`
	UserID        int32     `json:"user_id"`
	VotesCount    int32     `json:"votes_count"`
	LastDataVotes time.Time `json:"last_data_votes"`
}