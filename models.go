package jqdatatables

//jQuery DataTable request model
type JqTableRequest struct {
	//Number of current draw action (serial)
	Draw int

	//Ammount of records to skip
	Start int

	//Ammount of records to take
	Length int

	//Columns information from the script
	Columns []*JqTableColumn

	//Columns ordering information from the script
	Order []*JqTableOrder

	//Current search parameters
	Search JqTableSearch
}

//jQuery DataTable table column model
type JqTableColumn struct {
	Index      int
	Data       string
	Name       string
	Searchable bool
	Orderable  bool
	Search     JqTableSearch
}

//jQuery DataTable order model
type JqTableOrder struct {
	//Order
	Index int

	//Column index
	ColumnIndex int

	//JqDataTable column instance
	Column *JqTableColumn

	//Sorting direction
	Direction JqTableColumnDirection
}

//jQuery DataTable table search model
type JqTableSearch struct {
	//Value to search
	Value string

	//Set if we should use regular expression
	Regex bool
}
