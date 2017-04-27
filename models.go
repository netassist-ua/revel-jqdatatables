package jqdatatables

type JqTableColumnDirection string

const (
	ORDER_ASC  JqTableColumnDirection = "asc"
	ORDER_DESC JqTableColumnDirection = "desc"
)

func (d JqTableColumnDirection) Valid() bool {
	return d == ORDER_ASC || d == ORDER_DESC
}

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

//jQuery DataTables table search model
type JqTableSearch struct {
	//Value to search
	Value string

	//Set if we should use regular expression
	Regex bool
}

//JQuery DataTables object source response
type JqTableResponse struct {
	Draw            int `json:"draw"`
	RecordsTotal    int `json:"recordsTotal"`
	RecordsFiltered int `json:"recordsFiltered"`

	Data interface{} `json:"data"`
}
