package jsql

// TableSchema table schema
type TableSchema struct {
	// ColumnName column name
	ColumnName string `json:"COLUMN_NAME"`
	// DataType column data type
	DataType string `json:"DATA_TYPE"`
	// IsNullable is null(YES or NO)
	IsNullable string `json:"IS_NULLABLE"`
	// DataDefault column default data
	DataDefault string `json:"DATA_DEFAULT"`
	// PrimaryKey primary key
	PrimaryKey interface{} `json:"PRIMARY_KEY"`
	// IsIdentity is identity(YES OR NO)
	IsIdentity string `json:"IS_IDENTITY"`
	// ColumnComment column comment
	ColumnComment string `json:"COLUMN_COMMENT"`
	// TableComment table comment
	TableComment string `json:"TABLE_COMMENT"`
}
