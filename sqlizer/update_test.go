package sqlizer_test

import (
	"github.com/tatskaari/gendb/builder"
	"github.com/tatskaari/gendb/sqlizer"
)

func (s *sqlizerSuite) TestUpdate(){
	sql, args := new(sqlizer.StandardSqlizer).Update(
		new(builder.UpdateBuilder).Update("foo").
			Set("bar_id", 1234).
			Set("name", builder.Bind("new_name")).
			Where(builder.Eq("id", 4321)).
			And(builder.Eq("name", builder.Bind("old_name"))).
			UpdateBuilder,
	)

	s.Equal("UPDATE foo SET bar_id = ?, name = ? WHERE id = ? AND name = ?", sql)

	s.Require().Len(args, 4)
	s.Equal(args[0], 1234)
	s.Equal(args[1], "new_name")
	s.Equal(args[2], 4321)
	s.Equal(args[3], "old_name")
}
