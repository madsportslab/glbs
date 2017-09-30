package glbs

import (
	"os"
	"testing"
)

var key string

func TestPut(t *testing.T) {

	file, err := os.Open("blob.go")

	defer file.Close()

	if err != nil {
		t.Error(err)
	}

  SetNamespace("vmw")

  key, err = Put(file)

	if err != nil {
		t.Error(err)
	}

	if len(key) == 0 {
		t.Error("Blob not stored.")
	}

} // TestPut

func TestGet(t *testing.T) {

	SetNamespace("vmw")

  buf, err := Get(key)

	if err != nil {
		t.Error(err)
	}

	if len(buf) == 0 {
		t.Error("Blob exists, but not found.")
	}

} // TestGet

func TestGetNonExistent(t *testing.T) {

	SetNamespace("vmw")

  _, err := Get("abc")

  if err == nil {
		t.Error("blob should not exist")
	}

} // TestGetNonExistent

func TestExists(t *testing.T) {

	SetNamespace("vmw")

  if !Exists(key) {
		t.Error("Blob not found, but should exist.")
	}

} // TestExists

func TestExistsBadKey(t *testing.T) {

	SetNamespace("vmw")

  if Exists("abc") {
		t.Error("Blob found, but should not exist.")
	}

} // TestExistsBadKey
