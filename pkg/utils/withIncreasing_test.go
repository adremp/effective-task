package utils_test

import (
	"effective-task/pkg/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)


type WInc struct {
	idx   int
	field string
}

func TestWithIncreasing(t *testing.T) {
	formatField := func(f WInc) string {
		return fmt.Sprintf("%v ILIKE $%v", f.field, f.idx)
	}

	t.Run("WithIncreasing", func(t *testing.T) {
		idx, format := utils.WithIncreasing(formatField)
		res1 := format(WInc{*idx, "name"})
		res2 := format(WInc{*idx, "surname"})
		res3 := format(WInc{*idx, "patronymic"})
		require.Equal(t, res1, "name ILIKE $1")
		require.Equal(t, res2, "surname ILIKE $2")
		require.Equal(t, res3, "patronymic ILIKE $3")
	})
}
