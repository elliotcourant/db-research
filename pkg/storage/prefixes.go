package storage

type MetaPrefix = byte

const (
	MetaPrefix_Database       MetaPrefix = 'd'
	MetaPrefix_Schema                    = 's'
	MetaPrefix_Table                     = 't'
	MetaPrefix_Column                    = 'c'
	MetaPrefix_MyId                      = '@'
	MetaPrefix_Log                       = 'l'
	MetaPrefix_StableStore               = 's'
	MetaPrefix_Datum                     = '>'
	MetaPrefix_ObjectSequence            = '+'
)
