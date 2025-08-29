package reservationrepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/reservations"
	"errors"
	"gorm.io/gorm"
	"Amora/internal/helper"
)

type ReservationRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*reservations.Reservations]
}

func NewreservationRepositoryInit(DB *gorm.DB) (*ReservationRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*reservations.Reservations](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &ReservationRepositoryInit{ Core : coreRepo }, nil
}

func (r *ReservationRepositoryInit) GetByUserId(userId int) ([]reservations.Reservations, error) {
	var reservation []reservations.Reservations
	err := r.Core.DB.Where("user_id = ?", userId).Find(&reservation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("No reservations found for user ID %d", userId)
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving reservations for user ID %d: %v", userId, err)
		return nil, err
	}
	return reservation, nil
}

// getbyEventId this will drop down into the details of the event like restaurant, venue, experience the name image website and location city state and address 
func (r *ReservationRepositoryInit) GetByEventId(eventId int) ([]helper.ReservationDetailsDTO, error) {
	query := `
		SELECT
			r.id AS reservation_id,
			r.reservation_time,
			r.party_size,
			r.special_requests,
			p.firstname || ' ' || p.lastname AS partner_name,

			-- Venue
			v.id AS venue_id,
			v.name AS venue_name,
			v.address AS venue_addr,
			v.city AS venue_city,
			v.state AS venue_state,
			v.zip_code AS venue_zip,
			v.phone_number AS venue_phone,
			v.email AS venue_email,
			v.website AS venue_site,
			v.price AS venue_price,
			v.image_url AS venue_img,

			-- Restaurant
			rest.id AS restaurant_id,
			rest.name AS restaurant_name,
			rest.description AS restaurant_desc,
			rest.address AS restaurant_addr,
			rest.city AS restaurant_city,
			rest.state AS restaurant_state,
			rest.zip_code AS restaurant_zip,
			rest.image_url AS restaurant_img,
			rest.website AS restaurant_site,

			-- Experience
			e.id AS experience_id,
			e.name AS experience_name,
			e.description AS experience_desc,
			e.address AS experience_addr,
			e.city AS experience_city,
			e.state AS experience_state,
			e.zip_code AS experience_zip,
			e.phone_number AS experience_phone,
			e.price AS experience_price,
			e.image_url AS experience_img
		FROM reservations r
		LEFT JOIN venueslist v ON r.venue_id = v.id
		LEFT JOIN resteruatelist rest ON r.restaurant_id = rest.id
		LEFT JOIN experienceslist e ON r.experience_id = e.id
		LEFT JOIN Partners p ON r.partner_id = p.id
		WHERE r.event_id = ?
	`

	// Define an intermediate flat struct for scanning
	var rows []struct {
		ReservationID   int
		ReservationTime string
		PartySize       int
		SpecialRequests *string
		PartnerName     string

		VenueID    *int
		VenueName  *string
		VenueAddr  *string
		VenueCity  *string
		VenueState *string
		VenueZip   *string
		VenuePhone *string
		VenueEmail *string
		VenuePrice *float64
		VenueSite  *string
		VenueImg   *string

		RestaurantID    *int
		RestaurantName  *string
		RestaurantDesc  *string
		RestaurantAddr  *string
		RestaurantCity  *string
		RestaurantState *string
		RestaurantZip   *string
		RestaurantImg   *string
		RestaurantSite  *string

		ExperienceID    *int
		ExperienceName  *string
		ExperienceDesc  *string
		ExperienceAddr  *string
		ExperienceCity  *string
		ExperienceState *string
		ExperienceZip   *string
		ExperiencePhone *string
		ExperiencePrice *float64
		ExperienceImg   *string
	}

	if err := r.Core.DB.Raw(query, eventId).Scan(&rows).Error; err != nil {
		return nil, err
	}

	var results []helper.ReservationDetailsDTO
	for _, row := range rows {
		dto := helper.ReservationDetailsDTO{
			ReservationID:   row.ReservationID,
			ReservationTime: row.ReservationTime,
			PartySize:       row.PartySize,
			SpecialRequests: row.SpecialRequests,
			PartnerName:     row.PartnerName,
		}

		switch {
		case row.VenueID != nil:
			dto.VenueDetails = &helper.VenueDetailsDTO{
				Id:        *row.VenueID,
				VenueName: helper.Deref(row.VenueName),
				Address:   helper.Deref(row.VenueAddr),
				City:      helper.Deref(row.VenueCity),
				State:     helper.Deref(row.VenueState),
				ZipCode:   helper.Deref(row.VenueZip),
				Phone:     row.VenuePhone,
				Email:     row.VenueEmail,
				Website:   row.VenueSite,
				ImageURL:  row.VenueImg,
				Price:     row.VenuePrice,
			}
		case row.RestaurantID != nil:
			dto.RestaurantDetails = &helper.RestaurantDetailsDTO{
				Id:             *row.RestaurantID,
				RestaurantName: helper.Deref(row.RestaurantName),
				Description:    row.RestaurantDesc,
				Address:        helper.Deref(row.RestaurantAddr),
				City:           helper.Deref(row.RestaurantCity),
				State:          helper.Deref(row.RestaurantState),
				ZipCode:        helper.Deref(row.RestaurantZip),
				ImageURL:       row.RestaurantImg,
				Website:        row.RestaurantSite,
			}
		case row.ExperienceID != nil:
			dto.ExperienceDetails = &helper.ExperienceDetailsDTO{
				Id:             *row.ExperienceID,
				ExperienceName: helper.Deref(row.ExperienceName),
				Description:    row.ExperienceDesc,
				Address:        helper.Deref(row.ExperienceAddr),
				City:           helper.Deref(row.ExperienceCity),
				State:          helper.Deref(row.ExperienceState),
				ZipCode:        helper.Deref(row.ExperienceZip),
				Phone:          row.ExperiencePhone,
				Price:          row.ExperiencePrice,
				ImageURL:       row.ExperienceImg,
			}
		default:
			Logging.Logs.Warnf("Reservation ID %d has no associated venue, restaurant, or experience", row.ReservationID)
			continue
		}

		results = append(results, dto)
	}

	return results, nil
}

