package trackingSchema

import (
	"log"
)

//GetSchema gets tracking schema
func GetSchema(eventID string) (schema *Schema, err error) {
	schema = &Schema{}

	err = stmtSchemaByEventID.QueryRow(eventID).Scan(&schema.EventID, &schema.Alias, &schema.Desc)
	if err != nil {
		log.Println("coudn't get schema")
		return
	}
	rows, err := stmtParamsByEventID.Query(eventID)
	if err != nil {
		log.Println("coudn't get schema")
		return
	}

	for rows.Next() {
		params := &Parameters{}
		err = rows.Scan(&params.Name, &params.DataType, &params.Mandatory, &params.Description, &params.Format)
		if err != nil {
			log.Println("coudn't get schema")
			return
		}
		schema.Params = append(schema.Params, params)
	}

	return
}
