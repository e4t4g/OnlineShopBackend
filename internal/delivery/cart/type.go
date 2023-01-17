package cart

type Cart struct {
	Id     string     `json:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	UserId string     `json:"userId,omitempty" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	Items  []CartItem `json:"items" binding:"min=0" minimum:"0"`
}

type ShortCart struct {
	CartId string `json:"cartId" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	ItemId string `json:"itemId" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
}

type CartId struct {
	Value string `json:"id" uri:"cartID" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
}

type CartItem struct {
	Item `json:"item"`
	Quantity `json:"quantity"`
}

type Item struct {
	Id       string `json:"itemId" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	Title    string `json:"title" binding:"required" example:"Пылесос"`
	Price    int32  `json:"price" example:"1990" default:"10" binding:"required" minimum:"0"`
	Image    string `json:"image,omitempty"`
	Category `json:"category"`
}

type Quantity struct {
	Quantity int `json:"quantity" example:"3" default:"1" binding:"required" minimum:"1"`
}

type Category struct {
	Id          string `json:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	Name        string `json:"name" binding:"required" example:"Электротехника"`
	Description string `json:"description" binding:"required" example:"Электротехнические товары для дома"`
	Image       string `json:"image,omitempty"`
}
