package main

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/vattle/sqlboiler/bdb"
	"github.com/vattle/sqlboiler/bdb/drivers"
)

func TestTxtsFromOne(t *testing.T) {
	t.Parallel()

	tables, err := bdb.Tables(&drivers.MockDriver{}, "public", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	jets := bdb.GetTable(tables, "jets")
	texts := txtsFromFKey(tables, jets, jets.FKeys[0])
	expect := TxtToOne{}

	expect.ForeignKey = jets.FKeys[0]

	expect.LocalTable.NameGo = "Jet"
	expect.LocalTable.ColumnNameGo = "PilotID"

	expect.ForeignTable.NameGo = "Pilot"
	expect.ForeignTable.NamePluralGo = "Pilots"
	expect.ForeignTable.ColumnName = "id"
	expect.ForeignTable.ColumnNameGo = "ID"

	expect.Function.Name = "Pilot"
	expect.Function.ForeignName = "Jet"
	expect.Function.Varname = "pilot"

	expect.Function.LocalAssignment = "PilotID.Int"
	expect.Function.ForeignAssignment = "ID"

	if !reflect.DeepEqual(expect, texts) {
		t.Errorf("Want:\n%s\nGot:\n%s\n", spew.Sdump(expect), spew.Sdump(texts))
	}

	texts = txtsFromFKey(tables, jets, jets.FKeys[1])
	expect = TxtToOne{}
	expect.ForeignKey = jets.FKeys[1]

	expect.LocalTable.NameGo = "Jet"
	expect.LocalTable.ColumnNameGo = "AirportID"

	expect.ForeignTable.NameGo = "Airport"
	expect.ForeignTable.NamePluralGo = "Airports"
	expect.ForeignTable.ColumnName = "id"
	expect.ForeignTable.ColumnNameGo = "ID"

	expect.Function.Name = "Airport"
	expect.Function.ForeignName = "Jets"
	expect.Function.Varname = "airport"

	expect.Function.LocalAssignment = "AirportID"
	expect.Function.ForeignAssignment = "ID"

	if !reflect.DeepEqual(expect, texts) {
		t.Errorf("Want:\n%s\nGot:\n%s\n", spew.Sdump(expect), spew.Sdump(texts))
	}

	if !reflect.DeepEqual(expect, texts) {
		t.Errorf("Want:\n%s\nGot:\n%s\n", spew.Sdump(expect), spew.Sdump(texts))
	}
}

func TestTxtsFromOneToOne(t *testing.T) {
	t.Parallel()

	tables, err := bdb.Tables(&drivers.MockDriver{}, "public", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	pilots := bdb.GetTable(tables, "pilots")
	texts := txtsFromOneToOne(tables, pilots, pilots.ToOneRelationships[0])
	expect := TxtToOne{}

	expect.ForeignKey = bdb.ForeignKey{
		Name: "none",

		Table:    "jets",
		Column:   "pilot_id",
		Nullable: true,
		Unique:   true,

		ForeignTable:          "pilots",
		ForeignColumn:         "id",
		ForeignColumnNullable: false,
		ForeignColumnUnique:   false,
	}

	expect.LocalTable.NameGo = "Pilot"
	expect.LocalTable.ColumnNameGo = "ID"

	expect.ForeignTable.NameGo = "Jet"
	expect.ForeignTable.NamePluralGo = "Jets"
	expect.ForeignTable.ColumnName = "pilot_id"
	expect.ForeignTable.ColumnNameGo = "PilotID"

	expect.Function.Name = "Jet"
	expect.Function.ForeignName = "Pilot"
	expect.Function.Varname = "jet"

	expect.Function.LocalAssignment = "ID"
	expect.Function.ForeignAssignment = "PilotID.Int"

	if !reflect.DeepEqual(expect, texts) {
		t.Errorf("Want:\n%s\nGot:\n%s\n", spew.Sdump(expect), spew.Sdump(texts))
	}
}

func TestTxtsFromMany(t *testing.T) {
	t.Parallel()

	tables, err := bdb.Tables(&drivers.MockDriver{}, "public", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	pilots := bdb.GetTable(tables, "pilots")
	texts := txtsFromToMany(tables, pilots, pilots.ToManyRelationships[0])
	expect := TxtToMany{}
	expect.LocalTable.NameGo = "Pilot"
	expect.LocalTable.NameSingular = "pilot"
	expect.LocalTable.ColumnNameGo = "ID"

	expect.ForeignTable.NameGo = "License"
	expect.ForeignTable.NameSingular = "license"
	expect.ForeignTable.NamePluralGo = "Licenses"
	expect.ForeignTable.NameHumanReadable = "licenses"
	expect.ForeignTable.ColumnNameGo = "PilotID"
	expect.ForeignTable.Slice = "LicenseSlice"

	expect.Function.Name = "Licenses"
	expect.Function.ForeignName = "Pilot"
	expect.Function.LocalAssignment = "ID"
	expect.Function.ForeignAssignment = "PilotID"

	if !reflect.DeepEqual(expect, texts) {
		t.Errorf("Want:\n%s\nGot:\n%s\n", spew.Sdump(expect), spew.Sdump(texts))
	}

	texts = txtsFromToMany(tables, pilots, pilots.ToManyRelationships[1])
	expect = TxtToMany{}
	expect.LocalTable.NameGo = "Pilot"
	expect.LocalTable.NameSingular = "pilot"
	expect.LocalTable.ColumnNameGo = "ID"

	expect.ForeignTable.NameGo = "Language"
	expect.ForeignTable.NameSingular = "language"
	expect.ForeignTable.NamePluralGo = "Languages"
	expect.ForeignTable.NameHumanReadable = "languages"
	expect.ForeignTable.ColumnNameGo = "ID"
	expect.ForeignTable.Slice = "LanguageSlice"

	expect.Function.Name = "Languages"
	expect.Function.ForeignName = "Pilots"
	expect.Function.LocalAssignment = "ID"
	expect.Function.ForeignAssignment = "ID"

	if !reflect.DeepEqual(expect, texts) {
		t.Errorf("Want:\n%s\nGot:\n%s\n", spew.Sdump(expect), spew.Sdump(texts))
	}
}

func TestTrimSuffixes(t *testing.T) {
	t.Parallel()

	for _, s := range identifierSuffixes {
		a := "hello" + s

		if z := trimSuffixes(a); z != "hello" {
			t.Errorf("got %s", z)
		}
	}
}
