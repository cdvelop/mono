package godi

import (
	"fmt"
	"strings"
)

const (
	TypeInt    FieldType = "INT"
	TypeString FieldType = "VARCHAR(255)"
)

type FieldType string

func (t entity) CreateTableSQL() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", t.TableName))

	for i, column := range t.Fields {
		sb.WriteString(fmt.Sprintf("    %s %s", column.Name, column.Type))

		if column.Unique {
			sb.WriteString(" UNIQUE")
		}

		if column.PrimaryKey {
			sb.WriteString(" PRIMARY KEY")
			if column.Type == TypeInt {
				// sb.WriteString(" AUTO_INCREMENT")
			}
		}

		if column.NotNull {
			sb.WriteString(" NOT NULL")
		}

		if column.ForeignKey != nil {
			// CONSTRAINT fk_departments FOREIGN KEY (id_department) REFERENCES departments(id_department)
			sb.WriteString(fmt.Sprintf(",\n CONSTRAINT fk_%v FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE CASCADE",
				column.ForeignKey.TableName, column.Name, column.ForeignKey.TableName, column.Name))
		}

		if i < len(t.Fields)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString(");")
	return sb.String()
}
