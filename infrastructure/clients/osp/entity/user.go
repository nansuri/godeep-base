package entity

import "time"

type QueryUserInfoResponse struct {
	PageNum   int          `json:"pageNum"`
	PageSize  int          `json:"pageSize"`
	TotalUser int          `json:"totalUser"`
	UserData  []PublicUser `json:"userData"`
}

type GetUserInfoResponse struct {
	TraceId  string     `json:"traceId"`
	UserData PublicUser `json:"userData"`
}

type PublicUser struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	FullName  string    `gorm:"size:100;not null;" json:"fullName"`
	UserName  string    `gorm:"size:100;not null;" json:"userName"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Domain    string    `gorm:"size:100;not null;" json:"domain"`
	Role      string    `gorm:"size:100;not null;" json:"role"`
	AuthRole  string    `gorm:"size:100;not null;" json:"authRole"`
	IsDeleted bool      `gorm:"size:100;not null;" json:"isDeleted"`
	LastLogin time.Time `gorm:"size:100;not null;" json:"lastLogin"`
}
