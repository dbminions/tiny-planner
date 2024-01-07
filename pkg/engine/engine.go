package engine

import (
	"fmt"
	"github.com/apache/arrow/go/v12/arrow"
)

type LiteralValueVector struct {
	arrowType arrow.DataType
	value     any
	size      int
}

func (v LiteralValueVector) DataType() arrow.DataType {
	return v.arrowType
}

func (v LiteralValueVector) GetValue(i int) any {
	if i < 0 || i >= v.size {
		panic(fmt.Sprintf("index out of bounds %d vecsize: %d", i, v.size))
	}
	return v.value
}

func (v LiteralValueVector) Len() int {
	return v.size
}

func (s Schema) Select(projection []string) Schema {
	fields := make([]arrow.Field, 0)
	for _, columnName := range projection {
		field, ok := s.FieldsByName(columnName)
		if ok {
			fields = append(fields, field...)
		}
	}
	new := arrow.NewSchema(fields, nil)
	return Schema{new}
}
