package meetingfinder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"backend/internal/integrations/meeting-finder/dtos"
	"backend/internal/models"

	"github.com/mitchellh/mapstructure"
)

type UserLocation struct {
	Latitude  string
	Longitude string
}

func FindLocalMeetings(location UserLocation, languageCode string) ([]models.Congregation, error) {
	baseEndpoint := "https://apps.jw.org/api/public/meeting-search/weekly-meetings"
	urlObj, err := url.Parse(baseEndpoint)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	queryParams := urlObj.Query()
	queryParams.Add("includeSuggestions", "true")
	queryParams.Add("latitude", location.Latitude)
	queryParams.Add("longitude", location.Longitude)
	queryParams.Add("searchLanguageCode", languageCode)
	urlObj.RawQuery = queryParams.Encode()

	fmt.Println("URL:", urlObj.String())

	res, err := http.Get(urlObj.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonContent := string(body)
	var meetings dtos.GeoLocationList
	if err := json.Unmarshal([]byte(jsonContent), &meetings); err != nil {
		return nil, err
	}

	// Convert DTOs to Congregation
	var congregations []models.Congregation
	for _, meeting := range meetings.Locations {
		var phones []models.CongregationPhone
		if err := mapstructure.Decode(meeting.Properties.Phones, &phones); err != nil {
			return nil, err
		}

		newCong := models.Congregation{
			Name:    meeting.Properties.OrgName,
			Area:    meeting.Properties.OrgType,
			Address: meeting.Properties.Address,
			Users:   []models.User{}, // Empty until a user is added
		}
		if err := newCong.SetPhones(phones); err != nil {
			return nil, err
		}

		congregations = append(congregations, newCong)
	}

	return congregations, nil
}
