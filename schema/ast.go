package schema

type DDL interface {
	Statement() string
}

type CreateTable struct {
	statement string
	table     Table
}

type CreateIndex struct {
	statement string
	tableName string
	index     Index
}

type AddIndex struct {
	statement string
	tableName string
	index     Index
}

type AddPrimaryKey struct {
	statement string
	tableName string
	index     Index
}

type AddForeignKey struct {
	statement  string
	tableName  string
	foreignKey ForeignKey
}

type Table struct {
	name        string
	columns     []Column
	indexes     []Index
	foreignKeys []ForeignKey
	// XXX: have options and alter on its change?
}

type Column struct {
	name          string
	position      int
	typeName      string
	unsigned      bool
	notNull       bool
	autoIncrement bool
	defaultVal    *Value
	length        *Value
	scale         *Value
	keyOption     ColumnKeyOption
	onUpdate      *Value
	// TODO: keyopt
	// XXX: charset, collate, zerofill?
}

type Index struct {
	name      string
	indexType string // Parsed only in "create table" but not parsed in "add index". Only used inside `generateDDLsForCreateTable`.
	columns   []IndexColumn
	primary   bool
	unique    bool
}

type IndexColumn struct {
	column string
	length *Value // Parsed in "create table" but not parsed in "add index". So actually not used yet.
}

type ForeignKey struct {
	constraintName   string
	indexName        string
	indexColumns     []string
	referenceName    string
	referenceColumns []string
	onDelete         string
	onUpdate         string
}

type Value struct {
	valueType ValueType
	raw       []byte

	// ValueType-specific. Should be union?
	strVal   string  // ValueTypeStr
	intVal   int     // ValueTypeInt
	floatVal float64 // ValueTypeFloat
	bitVal   bool    // ValueTypeBit
}

type ValueType int

const (
	ValueTypeStr = ValueType(iota)
	ValueTypeInt
	ValueTypeFloat
	ValueTypeHexNum
	ValueTypeHex
	ValueTypeValArg
	ValueTypeBit
)

type ColumnKeyOption int

const (
	ColumnKeyNone = ColumnKeyOption(iota)
	ColumnKeyPrimary
	ColumnKeySpatialKey
	ColumnKeyUnique
	ColumnKeyUniqueKey
	ColumnKey
)

func (c *CreateTable) Statement() string {
	return c.statement
}

func (c *CreateIndex) Statement() string {
	return c.statement
}

func (a *AddIndex) Statement() string {
	return a.statement
}

func (a *AddPrimaryKey) Statement() string {
	return a.statement
}

func (a *AddForeignKey) Statement() string {
	return a.statement
}

func (t *Table) PrimaryKey() *Index {
	for _, index := range t.indexes {
		if index.primary {
			return &index
		}
	}

	primaryColumns := []IndexColumn{}
	for _, column := range t.columns {
		if column.keyOption == ColumnKeyPrimary {
			primaryColumns = append(primaryColumns, IndexColumn{
				column: column.name,
				length: column.length,
			})
		}
	}

	if len(primaryColumns) == 0 {
		return nil
	}

	return &Index{
		name:      "PRIMARY",
		indexType: "primary key",
		columns:   primaryColumns,
		primary:   true,
		unique:    true,
	}
}
