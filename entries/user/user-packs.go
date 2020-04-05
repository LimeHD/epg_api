package user

import "database/sql"

type UserPack struct {
	Id               int
	UserId           sql.NullInt32
	PackId           int
	PackIdentifierId sql.NullInt32
	Begin            int32
	End              int32
	Status           int
	UserDeviceId     sql.NullInt32

	Packs Pack
}

type UserPacks struct {
	PackId int
	UserId int
	Status int
}
