package DB

type DB struct {
	Store map[string]string
}

func (db *DB) Set(name string, myVal string) {
	db.Store[name] = myVal
}

func (db *DB) Get(name string) (string, bool) {
	dbValue, ok := db.Store[name]

	return dbValue, ok
}

func (db *DB) Delete(name string) {
	delete(db.Store, name)
}

func (db DB) Count(value string) int {
	count := 0
	for _, v := range db.Store {
		if value == v {
			count += 1
		}
	}
	return count
}

func NewDB() *DB {
	return &DB{
		Store: map[string]string{},
	}
}
