package main

import "testing"

func TestPut(t *testing.T) {
	Put("hello", "world")
	t.Log("insert a record")
	value := Get("hello")
	if value == "world" {
		t.Log("get a record")
	} else {
		t.Error("insert error")
	}

	t.Log("save")
	defer Save()
	t.Log("save")
}

func TestGet(t *testing.T) {
	value := Get("hello")
	if value == "world" {
		t.Log("get a record")
	} else {
		t.Error("insert error")
	}
	defer Save()
}

func TestDelete(t *testing.T) {
	Delete("hello")
	value := Get("hello")
	if value == "world" {
		t.Error("get a record")
	} else {
		t.Log("insert error")
	}
	defer Save()
}