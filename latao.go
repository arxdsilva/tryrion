package main


type LATAOResponse struct {
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

// https://goplay.space/#-fbQ6tq3AQC
func lataoFlightInfos() (lf LataoResponse, err error) {
	req, err := http.NewRequest("GET", "https://bff.latam.com/ws/proxy/booking-webapp-bff/v1/public/revenue/bestprices/oneway?departure=2020-03-14^&origin=RIO^&destination=SAO^&cabin=Y^&country=BR^&language=PT^&home=pt_br^&adult=1^&promoCode=", nil)
	if err != nil {
		return
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-Track-Id", "1218459210")
	req.Header.Set("Origin", "https://www.latam.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")
	req.Header.Set("X-Home", "pt_br")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-Flow-Id", "aa743252-5e75-4b7b-a537-db1fd030b441")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://www.latam.com/pt_br/apps/personas/booking?fecha1_dia=14^&fecha1_anomes=2020-03^&auAvailability=1^&ida_vuelta=ida^&vuelos_origen=Rio^%^20de^%^20Janeiro^&from_city1=RIO^&vuelos_destino=S^%^C3^%^A3o^%^20Paulo^&to_city1=SAO^&flex=1^&vuelos_fecha_salida_ddmmaaaa=14/03/2020^&cabina=Y^&nadults=1^&nchildren=0^&ninfants=0^&cod_promo=^&stopover_outbound_days=0^&stopover_inbound_days=0")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(req.Body).Decode(&lf)
	return
}
