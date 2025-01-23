package mono

import (
	"fmt"
)

type sourceData interface {
	DataSource() any
}

type dataSource struct {
	// 	DataSource() []map[string]string
	data sourceData
}

func (d dataSource) DataSource() {

	fmt.Println("DATA SOURCE FROM")

}
