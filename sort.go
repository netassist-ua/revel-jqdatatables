package jqdatatables

type OrderByIndex []*JqTableOrder

func (a OrderByIndex) Len() int           { return len(a) }
func (a OrderByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrderByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }

type ColumnByIndex []*JqTableColumn

func (a ColumnByIndex) Len() int           { return len(a) }
func (a ColumnByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ColumnByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }
