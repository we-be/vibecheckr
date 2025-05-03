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

		collection, err := dao.FindCollectionByNameOrId("6srnjmqmi4hml6z")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("3ak4azcw")

		// add
		new_locations := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "42awmtix",
			"name": "locations",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "wkgf7uttc4lj0mt",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_locations); err != nil {
			return err
		}
		collection.Schema.AddField(new_locations)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6srnjmqmi4hml6z")
		if err != nil {
			return err
		}

		// add
		del_location := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "3ak4azcw",
			"name": "location",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_location); err != nil {
			return err
		}
		collection.Schema.AddField(del_location)

		// remove
		collection.Schema.RemoveField("42awmtix")

		return dao.SaveCollection(collection)
	})
}
