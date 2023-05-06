// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DetailsColumns holds the columns for the "details" table.
	DetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "auto_id", Type: field.TypeUint32, Unique: true},
		{Name: "sample_col", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// DetailsTable holds the schema information for the "details" table.
	DetailsTable = &schema.Table{
		Name:       "details",
		Columns:    DetailsColumns,
		PrimaryKey: []*schema.Column{DetailsColumns[0]},
	}
	// PubsubMessagesColumns holds the columns for the "pubsub_messages" table.
	PubsubMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "auto_id", Type: field.TypeUint32, Unique: true},
		{Name: "message_id", Type: field.TypeString, Nullable: true, Default: "DefaultMsgID"},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultMsgState"},
		{Name: "resp_to_id", Type: field.TypeUUID, Nullable: true},
		{Name: "undo_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// PubsubMessagesTable holds the schema information for the "pubsub_messages" table.
	PubsubMessagesTable = &schema.Table{
		Name:       "pubsub_messages",
		Columns:    PubsubMessagesColumns,
		PrimaryKey: []*schema.Column{PubsubMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pubsubmessage_state_resp_to_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[7]},
			},
			{
				Name:    "pubsubmessage_state_undo_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[8]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DetailsTable,
		PubsubMessagesTable,
	}
)

func init() {
}
