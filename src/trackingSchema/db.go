package trackingSchema

import (
	"database/sql"

	"log"

	"fmt"

	_ "github.com/lib/pq" //postgres driver
)

var (
	dbTracking          *sql.DB
	stmtSchemaByEventID *sql.Stmt
	stmtParamsByEventID *sql.Stmt
)

func initDB() (err error) {
	trackingDocumentConnStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		dbCfg.Postgres["trackingDocument"].User,
		dbCfg.Postgres["trackingDocument"].Password,
		dbCfg.Postgres["trackingDocument"].Host,
		dbCfg.Postgres["trackingDocument"].Database,
		dbCfg.Postgres["trackingDocument"].SSLMode,
	)

	dbTracking, err = sql.Open("postgres", trackingDocumentConnStr)
	if err != nil {
		log.Println("couldn't open tracking db")
		return
	}
	return
}

func prepareStatement() (err error) {
	getSchemaBySQL := `
        SELECT event_id, alias, description
        FROM "eventSchema"
        WHERE %s
    `

	getParamBySQL := `
        SELECT name, data_type, mandatory, ep.description, format
		FROM "eventParams" ep
		LEFT JOIN "eventSchema" es ON es.event_id = $1
		WHERE id = es.id
    `

	stmtSchemaByEventID, err = dbTracking.Prepare(fmt.Sprintf(getSchemaBySQL, "event_id = $1"))
	if err != nil {
		log.Println("couldn't prepare statement")
	}
	stmtParamsByEventID, err = dbTracking.Prepare(getParamBySQL)
	if err != nil {
		log.Println("couldn't prepare statement")
	}
	return
}
