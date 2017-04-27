package jqdatatables

import (
	"errors"
	"github.com/revel/revel"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	column_regexp *regexp.Regexp
	order_regexp  *regexp.Regexp
)

func init() {
	column_regexp = regexp.MustCompile(`^columns\[(\d+)\]((?:\[(?:\w)+\])+)$`)
	order_regexp = regexp.MustCompile(`^order\[(\d+)\]((?:\[(?:\w)+\])+)$`)

}

func parse_order_property(order *JqTableOrder, property, value string) error {
	switch property {
	case "[column]":
		number, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		order.ColumnIndex = int(number)
		break
	case "[dir]":
		order.Direction = JqTableColumnDirection(value)
		break
	}
	return nil
}

func parse_search_property(search *JqTableSearch, property, value string) error {
	switch property {
	case "[search][value]", "search[value]":
		search.Value = value
		break
	case "[search][regex]", "search[regex]":
		regex, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		search.Regex = regex
		break
	default:
		return errors.New("Failed to parse search property")
	}
	return nil
}

func parse_column_property(column *JqTableColumn, property, value string) error {
	if strings.HasPrefix(property, "[search]") {
		return parse_search_property(&column.Search, property, value)
	}
	switch property {
	case "[name]":
		column.Name = value
		break
	case "[data]":
		column.Data = value
		break
	case "[searchable]":
		searchable, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		column.Searchable = searchable
		break
	case "[orderable]":
		orderable, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		column.Orderable = orderable
		break
	default:
		return errors.New("Unknown property")
	}
	return nil
}

//BindJqDataTableRequest binds jQuery DataTables AJAX request to a JqTableRequest model
func BindJqDataTableRequest(params *revel.Params) JqTableRequest {
	columns_map := make(map[int]*JqTableColumn)
	orders_map := make(map[int]*JqTableOrder)

	request := JqTableRequest{}
	params.Bind(&(request.Draw), "draw")
	params.Bind(&(request.Start), "start")
	params.Bind(&(request.Length), "length")

	for key, value := range params.Values {
		//Bind column properties
		if strings.HasPrefix(key, "column") {
			submatches := column_regexp.FindAllStringSubmatch(key, -1)
			if len(submatches) == 0 {
				continue
			}
			if len(submatches[0]) < 3 {
				continue
			}
			index, err := strconv.ParseInt(submatches[0][1], 10, 32)
			if err != nil {
				revel.ERROR.Printf("Error while parsing index %s into the string: %v", index, err)
				continue
			}
			column, ok := columns_map[int(index)]
			if !ok {
				column = &JqTableColumn{}
				column.Index = int(index)
				columns_map[int(index)] = column
			}
			if err := parse_column_property(column, submatches[0][2], value[0]); err != nil {
				revel.ERROR.Printf("Error while parsing column property %s = %s: %v", submatches[0][2], value[0], err)
			}
			continue
		}
		//Bind order properties
		if strings.HasPrefix(key, "order") {
			submatches := order_regexp.FindAllStringSubmatch(key, -1)
			if len(submatches) == 0 {
				continue
			}
			if len(submatches[0]) < 3 {
				continue
			}
			index, err := strconv.ParseInt(submatches[0][1], 10, 32)
			if err != nil {
				revel.ERROR.Printf("Error while parsing order index %s into the string: %v", index, err)
				continue
			}
			order, ok := orders_map[int(index)]
			if !ok {
				order = &JqTableOrder{}
				order.Index = int(index)
				orders_map[int(index)] = order
			}
			if err := parse_order_property(order, submatches[0][2], value[0]); err != nil {
				revel.ERROR.Printf("Error while parsing column property %s = %s: %v", submatches[0][2], value[0], err)
			}
			continue
		}
		if strings.HasPrefix(key, "search") {
			if err := parse_search_property(&(request.Search), key, value[0]); err != nil {
				revel.ERROR.Printf("Error while request search property %s = %s: %v", key, value[0], err)
			}
			continue
		}
	}

	//Convert orders map to slice
	orders_slice := make([]*JqTableOrder, len(orders_map))
	var i int = 0
	for _, value := range orders_map {
		col, ok := columns_map[value.ColumnIndex]
		if ok {
			//Set column
			value.Column = col
		}
		orders_slice[i] = value
		i++
	}

	columns_slice := make([]*JqTableColumn, len(columns_map))
	i = 0
	for _, value := range columns_map {
		columns_slice[i] = value
		i++
	}

	sort.Sort(OrderByIndex(orders_slice))
	sort.Sort(ColumnByIndex(columns_slice))
	request.Order = orders_slice
	request.Columns = columns_slice

	return request
}
