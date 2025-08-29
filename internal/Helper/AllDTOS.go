package helper

type ReservationDetailsDTO struct {
	ReservationID   int     `json:"reservation_id"`
	PartnerName     string  `json:"partner_name,omitempty"`
	ReservationTime string  `json:"reservation_time"`
	PartySize       int     `json:"party_size"`
	SpecialRequests *string `json:"special_requests,omitempty"`
	// there will only be one of these at a time each reservation is either a venue, restaurant or experience
	VenueDetails   *VenueDetailsDTO      `json:"venue_details,omitempty"`
	RestaurantDetails *RestaurantDetailsDTO `json:"restaurant_details,omitempty"`
	ExperienceDetails *ExperienceDetailsDTO `json:"experience_details,omitempty"`
}

type VenueDetailsDTO struct {
	Id       int     `json:"id"`
	VenueName string  `json:"venue_name"`
	Address  string  `json:"address"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	ZipCode  string  `json:"zip_code"`
	Phone    *string `json:"phone,omitempty"`
	Email    *string `json:"email,omitempty"`
	Website  *string `json:"website,omitempty"`
	Price 	*float64 `json:"price,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
}

type ExperienceDetailsDTO struct {
	Id             int     `json:"id"`
	ExperienceName string  `json:"experience_name"`
	Description    *string `json:"description,omitempty"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	ZipCode        string  `json:"zip_code"`
	Phone          *string `json:"phone,omitempty"`
	Price          *float64 `json:"price,omitempty"`
	ImageURL       *string `json:"image_url,omitempty"`
}

type RestaurantDetailsDTO struct {
	Id             int     `json:"id"`
	RestaurantName string  `json:"restaurant_name"`
	Description    *string `json:"description,omitempty"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	ZipCode        string  `json:"zip_code"`
	ImageURL       *string `json:"image_url,omitempty"`
	Website        *string `json:"website,omitempty"`
}

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

