package drivers

import (
	"github.com/vattle/sqlboiler/bdb"
	"github.com/vattle/sqlboiler/strmangle"
)

// MockDriver is a mock implementation of the bdb driver Interface
type MockDriver struct{}

// TableNames returns a list of mock table names
func (m *MockDriver) TableNames(exclude []string) ([]string, error) {
	tables := []string{"pilots", "jets", "airports", "licenses", "hangars", "languages", "pilot_languages"}
	return strmangle.SetComplement(tables, exclude), nil
}

// Columns returns a list of mock columns
func (m *MockDriver) Columns(tableName string) ([]bdb.Column, error) {
	return map[string][]bdb.Column{
		"pilots": {
			{Name: "id", Type: "int", DBType: "integer"},
			{Name: "name", Type: "string", DBType: "character"},
		},
		"airports": {
			{Name: "id", Type: "int", DBType: "integer"},
			{Name: "size", Type: "null.Int", DBType: "integer", Nullable: true},
		},
		"jets": {
			{Name: "id", Type: "int", DBType: "integer"},
			{Name: "pilot_id", Type: "int", DBType: "integer", Nullable: true, Unique: true},
			{Name: "airport_id", Type: "int", DBType: "integer"},
			{Name: "name", Type: "string", DBType: "character", Nullable: false},
			{Name: "color", Type: "null.String", DBType: "character", Nullable: true},
			{Name: "uuid", Type: "string", DBType: "uuid", Nullable: true},
			{Name: "identifier", Type: "string", DBType: "uuid", Nullable: false},
			{Name: "cargo", Type: "[]byte", DBType: "bytea", Nullable: false},
			{Name: "manifest", Type: "[]byte", DBType: "bytea", Nullable: true, Unique: true},
		},
		"licenses": {
			{Name: "id", Type: "int", DBType: "integer"},
			{Name: "pilot_id", Type: "int", DBType: "integer"},
		},
		"hangars": {
			{Name: "id", Type: "int", DBType: "integer"},
			{Name: "name", Type: "string", DBType: "character", Nullable: true, Unique: true},
		},
		"languages": {
			{Name: "id", Type: "int", DBType: "integer"},
			{Name: "language", Type: "string", DBType: "character", Nullable: false, Unique: true},
		},
		"pilot_languages": {
			{Name: "pilot_id", Type: "int", DBType: "integer"},
			{Name: "language_id", Type: "int", DBType: "integer"},
		},
	}[tableName], nil
}

// ForeignKeyInfo returns a list of mock foreignkeys
func (m *MockDriver) ForeignKeyInfo(tableName string) ([]bdb.ForeignKey, error) {
	return map[string][]bdb.ForeignKey{
		"jets": {
			{Name: "jets_pilot_id_fk", Column: "pilot_id", ForeignTable: "pilots", ForeignColumn: "id"},
			{Name: "jets_airport_id_fk", Column: "airport_id", ForeignTable: "airports", ForeignColumn: "id"},
		},
		"licenses": {
			{Name: "licenses_pilot_id_fk", Column: "pilot_id", ForeignTable: "pilots", ForeignColumn: "id"},
		},
		"pilot_languages": {
			{Name: "pilot_id_fk", Column: "pilot_id", ForeignTable: "pilots", ForeignColumn: "id"},
			{Name: "jet_id_fk", Column: "language_id", ForeignTable: "languages", ForeignColumn: "id"},
		},
	}[tableName], nil
}

// TranslateColumnType converts a column to its "null." form if it is nullable
func (m *MockDriver) TranslateColumnType(c bdb.Column) bdb.Column {
	p := &PostgresDriver{}
	return p.TranslateColumnType(c)
}

// PrimaryKeyInfo returns mock primary key info for the passed in table name
func (m *MockDriver) PrimaryKeyInfo(tableName string) (*bdb.PrimaryKey, error) {
	return map[string]*bdb.PrimaryKey{
		"pilots": {
			Name:    "pilot_id_pkey",
			Columns: []string{"id"},
		},
		"airports": {
			Name:    "airport_id_pkey",
			Columns: []string{"id"},
		},
		"jets": {
			Name:    "jet_id_pkey",
			Columns: []string{"id"},
		},
		"licenses": {
			Name:    "license_id_pkey",
			Columns: []string{"id"},
		},
		"hangars": {
			Name:    "hangar_id_pkey",
			Columns: []string{"id"},
		},
		"languages": {
			Name:    "language_id_pkey",
			Columns: []string{"id"},
		},
		"pilot_languages": {
			Name:    "pilot_languages_pkey",
			Columns: []string{"pilot_id", "language_id"},
		},
	}[tableName], nil
}

// UseLastInsertID returns a database mock LastInsertID compatability flag
func (m *MockDriver) UseLastInsertID() bool { return false }

// Open mimics a database open call and returns nil for no error
func (m *MockDriver) Open() error { return nil }

// Close mimics a database close call
func (m *MockDriver) Close() {}
