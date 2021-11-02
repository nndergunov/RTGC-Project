//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Usersinroom = newUsersinroomTable("public", "usersinroom", "")

type usersinroomTable struct {
	postgres.Table

	// Columns
	ID       postgres.ColumnInteger
	RoomID   postgres.ColumnInteger
	Userid   postgres.ColumnString
	Username postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UsersinroomTable struct {
	usersinroomTable

	EXCLUDED usersinroomTable
}

// AS creates new UsersinroomTable with assigned alias
func (a UsersinroomTable) AS(alias string) *UsersinroomTable {
	return newUsersinroomTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersinroomTable with assigned schema name
func (a UsersinroomTable) FromSchema(schemaName string) *UsersinroomTable {
	return newUsersinroomTable(schemaName, a.TableName(), a.Alias())
}

func newUsersinroomTable(schemaName, tableName, alias string) *UsersinroomTable {
	return &UsersinroomTable{
		usersinroomTable: newUsersinroomTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newUsersinroomTableImpl("", "excluded", ""),
	}
}

func newUsersinroomTableImpl(schemaName, tableName, alias string) usersinroomTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		RoomIDColumn   = postgres.IntegerColumn("room_id")
		UseridColumn   = postgres.StringColumn("userid")
		UsernameColumn = postgres.StringColumn("username")
		allColumns     = postgres.ColumnList{IDColumn, RoomIDColumn, UseridColumn, UsernameColumn}
		mutableColumns = postgres.ColumnList{IDColumn, RoomIDColumn, UseridColumn, UsernameColumn}
	)

	return usersinroomTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		// Columns
		ID:       IDColumn,
		RoomID:   RoomIDColumn,
		Userid:   UseridColumn,
		Username: UsernameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
