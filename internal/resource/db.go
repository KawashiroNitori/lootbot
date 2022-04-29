package resource

import (
	"fmt"

	"github.com/KawashiroNitori/lootbot/ent"
	_ "github.com/mattn/go-sqlite3"
)

var DBClient *ent.Client

func init() {
	var err error
	DBClient, err = ent.Open("sqlite3", "file:./userdata/lootbot.db?cache=shared&mode=rwc&_fk=1")
	if err != nil {
		panic(fmt.Errorf("failed to open connection to sqlite: %w", err))
	}
}
