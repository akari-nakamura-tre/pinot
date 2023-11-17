package models

type OutputData struct {
	Area                   string `json:"area"`
	AreaCode               int32  `json:"area_code"`
	BusinessDepartment     string `json:"business_department"`
	BusinessDepartmentCode int32  `json:"business_department_code"`
	CustomerCount          int32  `json:"customer_count"`
	Date                   string `json:"date"`
	Department             string `json:"department"`
	DepartmentCode         int32  `json:"department_code"`
	Division               string `json:"division"`
	DivisionCode           int32  `json:"division_code"`
	Jan                    string `json:"jan"`
	Line                   string `json:"line"`
	LineCode               int32  `json:"line_code"`
	ProductName            string `json:"product_name"`
	SalesQuantity          int32  `json:"sales_quantity"`
	Store                  string `json:"store"`
	StoreCode              int32  `json:"store_code"`
	Team                   string `json:"team"`
	TeamCode               int32  `json:"team_code"`
	TotalPrice             int32  `json:"total_price"`
	Zone                   string `json:"zone"`
	ZoneCode               int32  `json:"zone_code"`
}
