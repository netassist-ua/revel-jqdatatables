package jqdatatables

//Converts columns array to the map
func ColumnsToDataMap(columns []*JqTableColumn) map[string]*JqTableColumn {
	m := make(map[string]*JqTableColumn)
	for _, item := range columns {
		m[item.Data] = item
	}
	return m
}
