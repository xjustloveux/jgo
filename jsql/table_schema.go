package jsql

type TableSchema struct {
	ColumnName    string `json:"COLUMN_NAME"`
	DataType      string `json:"DATA_TYPE"`
	IsNullable    string `json:"IS_NULLABLE"`
	DataDefault   string `json:"DATA_DEFAULT"`
	PrimaryKey    *int8  `json:"PRIMARY_KEY"`
	ColumnComment string `json:"COLUMN_COMMENT"`
	TableComment  string `json:"TABLE_COMMENT"`
}
