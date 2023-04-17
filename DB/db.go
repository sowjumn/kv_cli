package DB

type DB struct {
	Store    map[string]string
	CountMap map[string]int
}

func (db *DB) Set(name string, myVal string) {
	db.Store[name] = myVal
	db.CountMap[myVal] += 1
}

func (db *DB) Get(name string) (string, bool) {
	dbValue, ok := db.Store[name]

	return dbValue, ok
}

func (db *DB) Delete(name string) {
	dbValue, ok := db.Store[name]

	if ok {
		db.CountMap[dbValue] -= 1
		delete(db.Store, name)
	}
}

func (db DB) Count(value string) int {
	return db.CountMap[value]
}

func NewDB() *DB {
	return &DB{
		Store:    map[string]string{},
		CountMap: map[string]int{},
	}
}
