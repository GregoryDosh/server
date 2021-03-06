// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package native

import (
	"reflect"
	"testing"

	"github.com/go-vela/server/database/sqlite"
	"github.com/go-vela/types/library"
)

func TestNative_Update(t *testing.T) {
	// setup types
	want := new(library.Secret)
	want.SetID(1)
	want.SetOrg("foo")
	want.SetRepo("bar")
	want.SetTeam("")
	want.SetName("baz")
	want.SetValue("foob")
	want.SetType("repo")
	want.SetImages([]string{"foo", "bar"})
	want.SetEvents([]string{"foo", "bar"})
	want.SetAllowCommand(false)

	// setup database
	db, _ := sqlite.NewTest()

	defer func() {
		db.Sqlite.Exec("delete from secrets;")
		_sql, _ := db.Sqlite.DB()
		_sql.Close()
	}()

	_ = db.CreateSecret(want)

	// run test
	s, err := New(
		WithDatabase(db),
	)
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	err = s.Update("repo", "foo", "bar", want)
	if err != nil {
		t.Errorf("Update returned err: %v", err)
	}

	got, _ := s.Get("repo", "foo", "bar", "baz")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Update is %v, want %v", got, want)
	}
}

func TestNative_Update_Invalid(t *testing.T) {
	// setup types
	sec := new(library.Secret)
	sec.SetName("baz")
	sec.SetValue("foob")

	// setup database
	db, _ := sqlite.NewTest()
	defer func() { _sql, _ := db.Sqlite.DB(); _sql.Close() }()

	// run test
	s, err := New(
		WithDatabase(db),
	)
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	err = s.Update("repo", "foo", "bar", sec)
	if err == nil {
		t.Errorf("Update should have returned err")
	}
}
