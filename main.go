package main

import (
	"fmt"

	"github.com/kaz/sql-mask/mask"
)

func main() {
	sqls := []string{`
		SELECT
			a,
			b,
			c,
			NULL,
			TRUE
		FROM (
			SELECT
				d,
				e,
				f
			FROM hogehoge
			WHERE g = 'hoge'
			AND h IN (1, 2, 3)
		) AS fugagufa
		ORDER BY DUMMY_FUNCTION(popopo, 5, "hello, \"world\" !") DESC
		LIMIT 20, 500
	`,
		"REPLACE INTO hoge VALUES(1, 2)",
		"REPLACE INTO hoge VALUES(1, 2, (",
	}

	for _, sql := range sqls {
		result, err := mask.Mask(sql)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
}
