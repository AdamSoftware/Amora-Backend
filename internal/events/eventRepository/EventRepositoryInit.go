package eventRepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/events"
	"Amora/internal/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)


type EventRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*events.Events]
}

func safeString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func NewEventRepositoryInit(DB *gorm.DB) (*EventRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*events.Events](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &EventRepositoryInit{ Core : coreRepo }, nil
}


func (e *EventRepositoryInit) FindByUserId(userId int) ([]events.Events, error) {
	var eventsList []events.Events
	err := e.Core.DB.Where("user_id = ?", userId).Find(&eventsList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("No events found for user ID %d", userId)
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving events for user ID %d: %v", userId, err)
		return nil, err
	}
	return eventsList, nil
}


func (e *EventRepositoryInit) GetEventsWithGiftsByUserID(userID int) ([]helper.EventWithGiftsDTO, error) {
	var rows []struct {
		EventID       int
		EventName     string
		EventDescription     string
		EventDate     time.Time 
		GiftID        int
		GiftName      string
		GiftDescription      string
		ImageURL      string
		Price         float64
	}

	err := e.Core.DB.Raw(`
		SELECT 
			e.id AS event_id,
			e.name AS event_name,
			e.description AS event_description,
			e.DateOfTheEvent AS event_date,
			g.id AS gift_id,
			g.name AS gift_name,
			g.description AS gift_description,
			g.image_url,
			g.price,
		FROM events e
		JOIN gifts g ON ag.gift_id = g.id
		WHERE e.user_id = ? AND e.is_active = true AND g.is_active = true
	`, userID).Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	eventMap := make(map[int]*helper.EventWithGiftsDTO)
	for _, row := range rows {
		if _, exists := eventMap[row.EventID]; !exists {
			eventMap[row.EventID] = &helper.EventWithGiftsDTO{
				Id:       row.EventID,
				Name:          row.EventName,
				Description:   row.EventDescription,
				DateOfTheEvent: row.EventDate.Format("2006-01-02"),
				Gifts:         []helper.GiftDTO{},
			}
		}

		eventMap[row.EventID].Gifts = append(eventMap[row.EventID].Gifts, helper.GiftDTO{
			Id:          row.GiftID,
			Name:        row.GiftName,
			Description: row.GiftDescription,
			ImageURL:    row.ImageURL,
			Price:       row.Price,
		})
	}

	// Convert map to slice
	var result []helper.EventWithGiftsDTO
	for _, e := range eventMap {
		result = append(result, *e)
	}

	return result, nil
}

func (e *EventRepositoryInit) GetEventsWithReservationsWithUserId(userID int) ([]helper.EventWithReservationsDTO, error) {
	var rows []struct {
		EventID          int        `gorm:"column:event_id"`
		EventName        string     `gorm:"column:event_name"`
		EventDescription string     `gorm:"column:event_description"`
		EventDate        time.Time  `gorm:"column:event_date"`

		ReservationID    int        `gorm:"column:reservation_id"`
		PartnerName      string     `gorm:"column:partner_name"`
		PartnerNickname  *string    `gorm:"column:partner_nickname"`

		// Names for quick reference
		RestaurantName   *string    `gorm:"column:restaurant_name"`
		ExperienceName   *string    `gorm:"column:experience_name"`
		VenueName        *string    `gorm:"column:venue_name"`

		// Full Restaurant info
		RestaurantID     *int       `gorm:"column:restaurant_id"`
		RestaurantDescription *string `gorm:"column:restaurant_description"`
		RestaurantAddress *string   `gorm:"column:restaurant_address"`
		RestaurantCity   *string    `gorm:"column:restaurant_city"`
		RestaurantState  *string    `gorm:"column:restaurant_state"`
		RestaurantZipCode *string   `gorm:"column:restaurant_zip_code"`
		RestaurantImageURL *string  `gorm:"column:restaurant_image_url"`
		RestaurantWebsite  *string  `gorm:"column:restaurant_website"`

		// Full Experience info
		ExperienceID     *int       `gorm:"column:experience_id"`
		ExperienceDescription *string `gorm:"column:experience_description"`
		ExperienceAddress *string   `gorm:"column:experience_address"`
		ExperienceCity   *string    `gorm:"column:experience_city"`
		ExperienceState  *string    `gorm:"column:experience_state"`
		ExperienceZipCode *string   `gorm:"column:experience_zip_code"`
		ExperiencePhone  *string    `gorm:"column:experience_phone"`
		ExperiencePrice  *float64   `gorm:"column:experience_price"`
		ExperienceImageURL *string  `gorm:"column:experience_image_url"`

		// Full Venue info
		VenueID          *int       `gorm:"column:venue_id"`
		VenueAddress     *string    `gorm:"column:venue_address"`
		VenueCity        *string    `gorm:"column:venue_city"`
		VenueState       *string    `gorm:"column:venue_state"`
		VenueZipCode     *string    `gorm:"column:venue_zip_code"`
		VenuePhone       *string    `gorm:"column:venue_phone"`
		VenueEmail       *string    `gorm:"column:venue_email"`
		VenueWebsite     *string    `gorm:"column:venue_website"`
		VenueImageURL    *string    `gorm:"column:venue_image_url"`

		ReservationTime  time.Time  `gorm:"column:reservation_time"`
		PartySize        int        `gorm:"column:party_size"`
		SpecialRequests  *string    `gorm:"column:special_requests"`
	}

	err := e.Core.DB.Raw(`
		SELECT
			e.id AS event_id,
			e.name AS event_name,
			e.description AS event_description,
			e.DateOfTheEvent AS event_date,

			r.id AS reservation_id,

			p.firstname || ' ' || p.lastname AS partner_name,
			p.nickname AS partner_nickname,

			res.id AS restaurant_id,
			res.name AS restaurant_name,
			res.description AS restaurant_description,
			res.address AS restaurant_address,
			res.city AS restaurant_city,
			res.state AS restaurant_state,
			res.zip_code AS restaurant_zip_code,
			res.image_url AS restaurant_image_url,
			res.website AS restaurant_website,

			exp.id AS experience_id,
			exp.name AS experience_name,
			exp.description AS experience_description,
			exp.address AS experience_address,
			exp.city AS experience_city,
			exp.state AS experience_state,
			exp.zip_code AS experience_zip_code,
			exp.phone_number AS experience_phone,
			exp.price AS experience_price,
			exp.image_url AS experience_image_url,

			v.id AS venue_id,
			v.name AS venue_name,
			v.address AS venue_address,
			v.city AS venue_city,
			v.state AS venue_state,
			v.zip_code AS venue_zip_code,
			v.phone_number AS venue_phone,
			v.email AS venue_email,
			v.website AS venue_website,
			v.image_url AS venue_image_url,

			r.reservation_time,
			r.party_size,
			r.special_requests
		FROM events e
		JOIN reservations r ON e.id = r.event_id
		LEFT JOIN partners p ON r.partner_id = p.id
		LEFT JOIN resteruatelist res ON r.restaurant_id = res.id
		LEFT JOIN experienceslist exp ON r.experience_id = exp.id
		LEFT JOIN venueslist v ON r.venue_id = v.id
		WHERE e.user_id = ? AND e.is_active = true
	`, userID).Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	eventMap := make(map[int]*helper.EventWithReservationsDTO)
	for _, row := range rows {
		if _, ok := eventMap[row.EventID]; !ok {
			eventMap[row.EventID] = &helper.EventWithReservationsDTO{
				Id:             row.EventID,
				Name:           row.EventName,
				Description:    row.EventDescription,
				DateOfTheEvent: row.EventDate.Format("2006-01-02"),
				Reservations:   []helper.ReservationDTO{},
			}
		}

		reservation := helper.ReservationDTO{
			ID:              row.ReservationID,
			PartnerName:     row.PartnerName,
			PartnerNickname: safeString(row.PartnerNickname),
			RestaurantName:  row.RestaurantName,
			ExperienceName:  row.ExperienceName,
			VenueName:       row.VenueName,
			ReservationTime: row.ReservationTime.Format("2006-01-02 15:04:05"),
			PartySize:       row.PartySize,
			SpecialRequests: row.SpecialRequests,
		}

		if row.RestaurantID != nil {
			reservation.Restaurant = &helper.RestaurantDTO{
				Id:          *row.RestaurantID,
				RestaurantName: safeString(row.RestaurantName),
				Description: row.RestaurantDescription,
				Address:     safeString(row.RestaurantAddress),
				City:        safeString(row.RestaurantCity),
				State:       safeString(row.RestaurantState),
				ZipCode:     safeString(row.RestaurantZipCode),
				ImageURL:    row.RestaurantImageURL,
				Website:     row.RestaurantWebsite,
			}
		}

		if row.ExperienceID != nil {
			reservation.Experience = &helper.ExperienceDTO{
				Id:           *row.ExperienceID,
				ExperienceName: safeString(row.ExperienceName),
				Description:  row.ExperienceDescription,
				Address:      safeString(row.ExperienceAddress),
				City:         safeString(row.ExperienceCity),
				State:        safeString(row.ExperienceState),
				ZipCode:      safeString(row.ExperienceZipCode),
				Phone:        row.ExperiencePhone,
				Price:        row.ExperiencePrice,
				ImageURL:     row.ExperienceImageURL,
			}
		}

		if row.VenueID != nil {
			reservation.Venue = &helper.VenueDTO{
				Id:        *row.VenueID,
				VenueName: safeString(row.VenueName),
				Address:   safeString(row.VenueAddress),
				City:      safeString(row.VenueCity),
				State:     safeString(row.VenueState),
				ZipCode:   safeString(row.VenueZipCode),
				Phone:     row.VenuePhone,
				Email:     row.VenueEmail,
				Website:   row.VenueWebsite,
				ImageURL:  row.VenueImageURL,
			}
		}

		eventMap[row.EventID].Reservations = append(eventMap[row.EventID].Reservations, reservation)
	}

	var result []helper.EventWithReservationsDTO
	for _, v := range eventMap {
		result = append(result, *v)
	}

	return result, nil
}

