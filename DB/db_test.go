package DB

import "testing"

func TestDBSetGet(t *testing.T) {
	mainDB := NewDB()

	mainDB.Set("a", "foo")
	val, _ := mainDB.Get("a")
	if val != "foo" {
		t.Errorf("Expected %s : Got %s", "foo", val)
	}

	mainDB.Set("b", "bar")
	val, _ = mainDB.Get("b")
	if val != "bar" {
		t.Errorf("Expected %s : Got %s", "bar", val)
	}

	val, _ = mainDB.Get("a")
	if val != "foo" {
		t.Errorf("Expected %s : Got %s", "foo", val)
	}
}

func TestDBCount(t *testing.T) {
	mainDB := NewDB()

	mainDB.Set("a", "foo")
	mainDB.Set("b", "bar")

	val := mainDB.Count("foo")
	if val != 1 {
		t.Errorf("Expected %d: Got %d", 1, val)
	}

	mainDB.Set("b", "foo")
	val = mainDB.Count("foo")
	if val != 2 {
		t.Errorf("Expected %d: Got %d", 2, val)
	}
}

func TestDBDelete(t *testing.T) {
	mainDB := NewDB()

	mainDB.Set("a", "foo")
	mainDB.Set("b", "bar")
	mainDB.Delete("a")
	_, ok := mainDB.Get("a")
	if ok != false {
		t.Errorf("Expected Value to not exist but got value for key: %s", "a")
	}
}
