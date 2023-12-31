package vector

import (
	"github.com/stretchr/testify/require"
	"testing"
	"tiny_planner/pkg/c_sql/c_exec_engine/a_coldata/a_types"
)

func Test1(t *testing.T) {
	vec := NewVec(types.T_int32.ToType())
	err := vec.Append(int32(1), false)
	require.NoError(t, err)

	err = vec.Append(int32(2), false)
	require.NoError(t, err)

	err = vec.Append(int32(0), true)
	require.NoError(t, err)

	v, null := Get[int32](vec, 0)
	require.Equal(t, int32(1), v)

	v, null = Get[int32](vec, 1)
	require.Equal(t, int32(2), v)

	v, null = Get[int32](vec, 2)
	require.Equal(t, true, null)
}
