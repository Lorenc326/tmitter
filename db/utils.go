package db

type Iterate func() interface{}
type CloseIterator func()
