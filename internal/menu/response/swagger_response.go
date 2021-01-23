package response

type Base struct {
	Header DetailHeader
	Data   string
}

type DetailHeader struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SwaggerMenuList struct {
	Header DetailHeader
	Data   []Menulist
}

type Menulist struct {
	Id        int    `json:"id"`
	MenuName  string `json:"menu_name"`
	MenuPrice int    `json:"menu_price"`
}

type SwaggerMenuDetail struct {
	Header DetailHeader
	Data   []MenuDetail
}

type MenuDetail struct {
	Id         int    `json:"id"`
	MenuName   string `json:"menu_name"`
	MenuDetail string `json:"menu_detail"`
	MenuPrice  int    `json:"menu_price"`
}
