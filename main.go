// mysql-args-anonymizer removes arguments from a MySQL query.
// It's useful to normalize the query to generate aggregates by query.
package main

import (
	"flag"
	"fmt"

	"github.com/pingcap/tidb/parser"
)

func main() {
	query := flag.String("q", "", "SQL query")
	flag.Parse()

	// We simply use TiDB's MySQL compatible parser to remove
	// arguments from the query, TiDB uses normalization
	// to report stats per query.
	fmt.Println(parser.Normalize(*query))
}
