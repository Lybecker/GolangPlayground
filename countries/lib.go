package countrylist

var countries map[string]string

func init() {
	countries = make(map[string]string)
	countries["dk"] = "Denmark"
	countries["uk"] = "United Kingdom"
}

// Get the contry based on the key
func Get(key string) string {
	return countries[key]
}
func Add(key, value string) {
	countries[key] = value
}
func GetAll() map[string]string {
	return countries
}
