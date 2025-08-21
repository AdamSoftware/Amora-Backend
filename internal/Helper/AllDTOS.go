package helper

type GiftDTO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       float64 `json:"price"`
}

type EventWithGiftsDTO struct {
	Id					int        `json:"id"`
	Name				string     `json:"name"`
	Description		string     `json:"description"`
	DateOfTheEvent	string     `json:"date_of_the_event"`
	Gifts				[]GiftDTO `json:"gifts"`
}

type AttachedGiftsWithGiftsDTO struct {
	Id          int        `json:"id"`
	EventId     int        `json:"event_id"`
	GiftId      int        `json:"gift_id"`
	Quantity    int        `json:"quantity"`
	IsActive    bool       `json:"is_active"`
	GiftDetails []GiftDTO `json:"gift_details"`
}

type EventWithReservationsDTO struct {
	Id             int
	Name           string
	Description    string
	DateOfTheEvent string 
	Reservations   []ReservationDTO
}

type ReservationDTO struct {
	ID              int       `json:"id"`
	PartnerName     string    `json:"partner_name,omitempty"`
	PartnerNickname string    `json:"partner_nickname,omitempty"`
	RestaurantName  *string   `json:"restaurant_name,omitempty"`
	ExperienceName  *string   `json:"experience_name,omitempty"`
	VenueName       *string   `json:"venue_name,omitempty"`
	ReservationTime string `json:"reservation_time"`
	PartySize       int       `json:"party_size"`
	SpecialRequests *string   `json:"special_requests,omitempty"`

  Restaurant     *RestaurantDTO  `json:"restaurant,omitempty"`
  Experience     *ExperienceDTO  `json:"experience,omitempty"`
  Venue          *VenueDTO       `json:"venue,omitempty"`
}

type PartnerDTO struct {
	Id         int
	PartnerName  string
	Nickname   *string
	Birthdate  string 
	Phone      string
	Anniversary *string
}

type RestaurantDTO struct {
	Id          int
	RestaurantName        string
	Description *string
	Address     string
	City        string
	State       string
	ZipCode     string
	ImageURL    *string
	Website     *string
}

type ExperienceDTO struct {
	Id          int
	ExperienceName string
	Description *string
	Address     string
	City        string
	State       string
	ZipCode     string
	Phone       *string
	Price       *float64
	ImageURL    *string
}

type VenueDTO struct {
	Id          int
	VenueName        string
	Address     string
	City        string
	State       string
	ZipCode     string
	Phone       *string
	Email       *string
	Website     *string
	ImageURL    *string
}

