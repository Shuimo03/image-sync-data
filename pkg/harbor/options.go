package harbor

type Options struct {
	Username        string `json:"principal"`
	Password        string `json:"password"`
	HarborAPIServer string `json:"HarborAPIServer"`
}
