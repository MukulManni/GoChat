package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type message struct {
	User  string `json:"user"`
	Msg   string `json:"msg"`
	Color string `json:"color"`
	Time  string `json:"time"`
}

var globalmsgList = []message{{"Mod", "Welcome to the global chat.", "red", "[00:00:00] "}}

func getAllMsgs() []message {
	return globalmsgList
}

func addMsg(data message) {
	globalmsgList = append(globalmsgList, data)
}

func connectDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS messages (
      user VARCHAR(64) NOT NULL,
	  msg VARCHAR(100),
	  color VARCHAR(10),
	  time VARCHAR(10),
      CHECK (CHAR_LENGTH(TRIM(user)) > 0)
    );
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func addMsgtoDB(db *sql.DB, user string, msg string, color string, time string) (*message, error) {
	created := message{}

	row := db.QueryRow(
		`INSERT INTO messages (user,msg,color,time) VALUES ($1,$2,$3,$4) RETURNING user, msg, color, time`,
		user, msg, color, time,
	)

	err := row.Scan(&created.User, &created.Msg, &created.Color, &created.Time)

	if err != nil {
		return nil, err
	}

	return &created, nil
}

func getAllMsgsDB(db *sql.DB) ([]message, error) {

	rows, err := db.Query(
		`SELECT * FROM messages`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	msgs := make([]message, 0, 10)

	for rows.Next() {
		m := message{}

		err = rows.Scan(&m.User, &m.Msg, &m.Color, &m.Time)

		if err != nil {
			return nil, err
		}

		msgs = append(msgs, m)
	}

	return msgs, nil
}
