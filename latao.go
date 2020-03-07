package main


type LTNResponse struct {
	Itinerary struct {
		Date               string `json:"date"`
		OriginDestinations []struct {
			Duration  int `json:"duration"`
			Departure struct {
				Airport   string `json:"airport"`
				City      string `json:"city"`
				Country   string `json:"country"`
				Timestamp string `json:"timestamp"`
			} `json:"departure"`
			Arrival struct {
				Airport   string `json:"airport"`
				City      string `json:"city"`
				Country   string `json:"country"`
				Timestamp string `json:"timestamp"`
			} `json:"arrival"`
		} `json:"originDestinations"`
		PricePerPassenger struct {
			ADULT struct {
				Base struct {
					Amount   int    `json:"amount"`
					Currency string `json:"currency"`
				} `json:"base"`
				Tax struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"tax"`
				Total struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"total"`
			} `json:"ADULT"`
		} `json:"pricePerPassenger"`
		Links struct {
			BookingDetails struct {
				Href string `json:"href"`
			} `json:"bookingDetails"`
			OrderCreation struct {
				Href string `json:"href"`
			} `json:"orderCreation"`
		} `json:"_links"`
	} `json:"itinerary"`
	BestPrices []struct {
		Date         string `json:"date"`
		Available    bool   `json:"available"`
		Availability string `json:"availability"`
		Price        struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"price"`
		PricePerSlice struct {
			Outbound struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"outbound"`
		} `json:"pricePerSlice"`
	} `json:"bestPrices"`
}
