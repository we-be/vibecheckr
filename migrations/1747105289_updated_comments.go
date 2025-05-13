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

		collection, err := dao.FindCollectionByNameOrId("d52kf0fi9aovpvk")
		if err != nil {
			return err
		}

		// add
		new_parent := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "kudngd0h",
			"name": "parent",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "d52kf0fi9aovpvk",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_parent); err != nil {
			return err
		}
		collection.Schema.AddField(new_parent)

		// add
		new_reply_to := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "kmftqkc1",
			"name": "reply_to",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "d52kf0fi9aovpvk",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_reply_to); err != nil {
			return err
		}
		collection.Schema.AddField(new_reply_to)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("d52kf0fi9aovpvk")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("kudngd0h")

		// remove
		collection.Schema.RemoveField("kmftqkc1")

		return dao.SaveCollection(collection)
	})
}
