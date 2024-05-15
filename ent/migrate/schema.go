// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "date", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"default", "selebrate"}, Default: "default"},
		{Name: "time", Type: field.TypeString, Nullable: true},
		{Name: "show_time", Type: field.TypeBool, Default: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "timeline_event", Type: field.TypeInt, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_timelines_event",
				Columns:    []*schema.Column{EventsColumns[9]},
				RefColumns: []*schema.Column{TimelinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag", Type: field.TypeString},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// TimelinesColumns holds the columns for the "timelines" table.
	TimelinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_timeline", Type: field.TypeInt, Nullable: true},
	}
	// TimelinesTable holds the schema information for the "timelines" table.
	TimelinesTable = &schema.Table{
		Name:       "timelines",
		Columns:    TimelinesColumns,
		PrimaryKey: []*schema.Column{TimelinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timelines_users_timeline",
				Columns:    []*schema.Column{TimelinesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// TagEventColumns holds the columns for the "tag_event" table.
	TagEventColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeInt},
		{Name: "event_id", Type: field.TypeInt},
	}
	// TagEventTable holds the schema information for the "tag_event" table.
	TagEventTable = &schema.Table{
		Name:       "tag_event",
		Columns:    TagEventColumns,
		PrimaryKey: []*schema.Column{TagEventColumns[0], TagEventColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_event_tag_id",
				Columns:    []*schema.Column{TagEventColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_event_event_id",
				Columns:    []*schema.Column{TagEventColumns[1]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		TagsTable,
		TimelinesTable,
		UsersTable,
		TagEventTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = TimelinesTable
	TimelinesTable.ForeignKeys[0].RefTable = UsersTable
	TagEventTable.ForeignKeys[0].RefTable = TagsTable
	TagEventTable.ForeignKeys[1].RefTable = EventsTable
}
