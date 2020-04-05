package sqlizer_test

import (
	"github.com/tatskaari/gendb/builder"
	"github.com/tatskaari/gendb/sqlizer"
)

func (s *sqlizerSuite) TestInsert() {
	s.Run("basic insert", func() {
		sql, args := sqlizer.Insert(builder.Insert("foo").Values(map[string]interface{}{"id": 123}))

		s.Equal("INSERT INTO foo (id) VALUES (?)", sql)
		s.Len(args, 1)
		s.Equal(123, args[0])
	})

	s.Run("multi row insert", func() {
		sql, args := sqlizer.Insert(
			builder.Insert("foo").
				Values(map[string]interface{}{"id": 123}).
				Values(map[string]interface{}{"id": 456}),
		)

		s.Equal("INSERT INTO foo (id) VALUES (?), (?)", sql)
		s.Len(args, 2)
		s.Equal(123, args[0])
		s.Equal(456, args[1])
	})

	s.Run("multi row insert with multi args", func() {

		ib := builder.Insert("foo")

		// set the column order so the columns always come out in the same order
		ib.ColumnsOrder = map[string]int{
			"id":     0,
			"name":   1,
			"bar_id": 2,
		}

		sql, args := sqlizer.Insert(
			ib.
				Values(map[string]interface{}{"id": 123, "name": builder.Bind("first"), "bar_id": 234}).
				Values(map[string]interface{}{"id": 321, "name": builder.Bind("second"), "bar_id": 432}),
		)

		s.Equal("INSERT INTO foo (id, name, bar_id) VALUES (?, ?, ?), (?, ?, ?)", sql)
		s.Len(args, 6)

		s.Equal(123, args[0])
		s.Equal("first", args[1])
		s.Equal(234, args[2])

		s.Equal(321, args[3])
		s.Equal("second", args[4])
		s.Equal(432, args[5])

	})
}
