package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rv5xkk58543igib")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ttywekdq")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rv5xkk58543igib")
		if err != nil {
			return err
		}

		// add
		del_created_dt := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ttywekdq",
			"name": "created_dt",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_created_dt); err != nil {
			return err
		}
		collection.Schema.AddField(del_created_dt)

		return dao.SaveCollection(collection)
	})
}
