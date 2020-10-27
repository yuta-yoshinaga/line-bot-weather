package entities

// Weather more info Weather condition codes)
type Weather struct {
	ID          int    `json:"id"`          // Weather condition id
	Main        string `json:"main"`        // Group of weather parameters (Rain, Snow, Extreme etc.)
	Description string `json:"description"` // Weather condition within the group. You can get the output in your language. Learn more
	Icon        string `json:"icon"`        // Weather icon id
}
