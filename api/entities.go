package api

type DetailedUser struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	BirthDate int64  `json:"birth_date"` // a Unix time that always has 0 hours, minutes and seconds
}

type Address struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

// type Preferences struct {
// 	Language string
//   DateFormat dateFormatPref
// }
//
// type dateFormatPref string
// const (
// 	MonthFirstFormat dateFormatPref = "dayfirst"
// 	DayFirstFormat   dateFormatPref = "monthfirst"
// )
