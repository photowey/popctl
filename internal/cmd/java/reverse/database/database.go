package database

type Database struct {
	Name   string `json:"name"`
	Tables []*Table
}

type Table struct {
	Name    string `json:"tableName"`
	Comment string `json:"comment"`
	Fields  []*Field
}

type Field struct {
	TableComment string `json:"tableComment"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Comment      string `json:"comment"`
	TypLen       int    `json:"typLen"`
	TypLength    string `json:"typLength"`
	NotNull      bool   `json:"notNull"`
	PrimaryKey   bool   `json:"primarykey"`
}
