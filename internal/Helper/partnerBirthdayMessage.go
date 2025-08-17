package helper


type PartnerMessage struct {
	PartnerId  int    `json:"partner_id"`
	PartnerName string `json:"partner_name"`
	Nickname string `json:"partner_nickname"`
	PhoneNumber string `json:"phone_number"`
	UserId int `json:"user_id"`	
	Message string `json:"message"`
}
