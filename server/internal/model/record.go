package model

import (
	"database/sql"
	"time"
)

type Record struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Score    int       `json:"score"`
	Comment  string    `json:"comment"`
	Finished time.Time `json:"finished"`
}

type RecordModel struct {
	DB *sql.DB
}

func NewRecordModel(db *sql.DB) *RecordModel { return &RecordModel{DB: db} }

func (m *RecordModel) Init() error {
	_, err := m.DB.Exec(`CREATE TABLE IF NOT EXISTS records (
		id BIGINT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(128) NOT NULL,
		score INT NOT NULL,
		comment TEXT,
		finished VARCHAR(64) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)
	return err
}

func (m *RecordModel) Insert(name string, score int, comment string) error {
	// store finished as RFC3339 string for consistent ordering and parsing
	_, err := m.DB.Exec(`INSERT INTO records (name, score, comment, finished) VALUES (?, ?, ?, ?)`, name, score, comment, time.Now().UTC().Format(time.RFC3339Nano))
	return err
}

func (m *RecordModel) Top(limit int) ([]Record, error) {
	rows, err := m.DB.Query(`SELECT id, name, score, comment, finished FROM records ORDER BY score DESC, finished ASC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Record
	for rows.Next() {
		var r Record
		var ts string
		if err := rows.Scan(&r.ID, &r.Name, &r.Score, &r.Comment, &ts); err != nil {
			return nil, err
		}
		pt, _ := time.Parse(time.RFC3339Nano, ts)
		r.Finished = pt
		list = append(list, r)
	}
	return list, nil
}

func (m *RecordModel) Recent(limit int) ([]Record, error) {
	rows, err := m.DB.Query(`SELECT id, name, score, comment, finished FROM records ORDER BY finished DESC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Record
	for rows.Next() {
		var r Record
		var ts string
		if err := rows.Scan(&r.ID, &r.Name, &r.Score, &r.Comment, &ts); err != nil {
			return nil, err
		}
		pt, _ := time.Parse(time.RFC3339Nano, ts)
		r.Finished = pt
		list = append(list, r)
	}
	return list, nil
}
