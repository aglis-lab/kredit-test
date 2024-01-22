package contract

type LoginCustomerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginCustomerResponse struct {
	Token string `json:"token"`
}

type LimitCustomerResponse struct {
	Risk     string  `json:"risk"`
	Interest float64 `json:"interest"`

	Limits []LimitCustomerItem `json:"limits"`
}

type LimitCustomerItem struct {
	Limit     int `json:"limit"`
	Period    int `json:"period"`
	UsedLimit int `json:"used_limit"`
}
