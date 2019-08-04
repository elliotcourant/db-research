package storage

import (
	"github.com/elliotcourant/arkdb/pkg/buffers"
)

var (
	columnMinimumSize = 8
)

type Column struct {
	ColumnID   uint8
	DatabaseID uint8
	SchemaID   uint8
	TableID    uint8
	ColumnName string
	ColumnType uint16
}

func (i Column) Size() int {
	return columnMinimumSize + len(i.ColumnName)
}

func (i Column) Path() []byte {
	l := len(i.ColumnName)
	buf := buffers.NewAllocatedBytesBuffer(columnMinimumSize + l)
	buf.Append(MetaPrefix_Column, i.DatabaseID, i.SchemaID, i.TableID, uint8(l))
	buf.AppendString(i.ColumnName)
	buf.Append(i.ColumnID)
	buf.AppendUint16(i.ColumnType)
	return buf.Bytes()
}
