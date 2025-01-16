package models

// type Extra struct {
// 	ID      int     `json:"id"`
// 	Name    string  `json:"name"`
// 	Price   float64 `json:"price"`
// 	Checked bool    `json:"checked"`
// }

// type Option struct {
// 	ID   int     `json:"id"`
// 	Name string  `json:"name"`
// 	List []Extra `json:"list"`
// }

// type Item struct {
// 	ID      int      `json:"id"`
// 	Name    string   `json:"name"`
// 	Image   string   `json:"image"`
// 	Tag     string   `json:"tag"`
// 	Price   float64  `json:"price"`
// 	Extras  []Extra  `json:"extras"`
// 	Options []Option `json:"options"`
// }

type Store struct {
	ID               int    `gorm:"primary_key"`
	Name             string `json:"name"`
	ActiveTime       string `json:"activeTime"`
	AllowedLocations string `json:"allowedLocations"`
	Image            string `json:"image"`
	Latlng           string `json:"latlng"`
	Tags             string `json:"tags"`
	Type             string `json:"type"`
	Uid              int    `json:"uid"`
	Items            string `json:"items"`
}

func (u *Store) TableName() string {
	return "stores"
}
