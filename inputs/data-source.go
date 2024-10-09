package inputs

import "fmt"

// type dataSource interface {
// 	DataSource() any
// }

type dataSource struct {
	// 	DataSource() []map[string]string
}

func (d dataSource) DataSource() {

	fmt.Println("DATA SOURCE FROM")

}
