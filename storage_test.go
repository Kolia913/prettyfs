package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPAthTransformFunc(key)

	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	expectedFileName := "6804429f74181a63c50c3d81d733a12f14a353ff"

	assert.Equal(t, pathKey.Pathname, expectedPathName)
	assert.Equal(t, pathKey.Filename, expectedFileName)
}

func TestStorageDeleteKey(t *testing.T) {
	s := newStorage()
	defer teardown(t, s)

	key := "momsbestpicture"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStorage(t *testing.T) {
	s := newStorage()
	defer teardown(t, s)

	key := "momsbestpicture"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); !ok {
		t.Errorf("expected to have key %s", key)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	fmt.Println(string(b))

	if string(b) != string(data) {
		t.Errorf("expected %s got %s", data, b)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); ok {
		t.Errorf("expected to not have key %s", key)
	}
}

func newStorage() *Storage {
	opts := StorageOpts{
		Root: "testing_fs_root",
		PathTransformFunc: CASPAthTransformFunc,
	}
	return NewStorage(opts)
}

func teardown(t *testing.T, s *Storage) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}