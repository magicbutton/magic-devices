package magicapp

import (
	"log"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/magicbutton/magic-devices/utils"
)

func OpenDatabase() {
	utils.Setup("./.env")

	dsn := viper.GetString("POSTGRES_DB")
	// verbose := viper.GetBool("BUN_VERBOSE")
	verbose := false
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)

	}
	config.PreferSimpleProtocol = true
	sqldb := stdlib.OpenDB(*config)
	utils.Db = bun.NewDB(sqldb, pgdialect.New())
	utils.Db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(verbose)))

}
