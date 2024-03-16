//go:build gorm

package database

// import (
// 	"lx-source/src/database/modules"
// 	"lx-source/src/env"
// 	"lx-source/src/sources"

// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func InitDB(dsn string) (err error) {
// 	loger := env.Loger.NewGroup(`InitDB`)
// 	defer loger.Free()
// 	DB, err = gorm.Open(modules.Sqlite(dsn), &gorm.Config{})
// 	if err == nil {
// 		for _, typ := range []struct {
// 			Name string
// 			Type interface{}
// 		}{
// 			{Name: T_music, Type: &XMusicItem{}},
// 			{Name: T_lyric, Type: &XLyricItem{}},
// 		} {
// 			for _, src := range sources.S_al {
// 				err = DB.Table(src + `_` + typ.Name).AutoMigrate(typ.Type)
// 				if err != nil {
// 					return
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// type Driver struct{}
