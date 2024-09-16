package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/models"
	"github.com/rs/zerolog/log"
)

type EventSourcingHttpClient struct {
	httpClient *http.Client
	url        string
}

func NewEventSourcingHttpClient(urlStr string) (*EventSourcingHttpClient, error) {

	path, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	baseUrl := fmt.Sprintf("%s://%s", path.Scheme, path.Host)

	httpClient := http.Client{}
	return &EventSourcingHttpClient{
		httpClient: &httpClient,
		url:        baseUrl,
	}, nil
}

func (client *EventSourcingHttpClient) AddEvents(aggregateId string, events []models.Event) error {
	bodyBytes, err := json.Marshal(events)
	if err != nil {
		log.Info().Err(err).Msg("Could not marshal events")
		return err
	}
	buf := bytes.NewBuffer(bodyBytes)
	addEventsUrl, err := url.JoinPath(client.url, fmt.Sprintf("/%s/events", aggregateId))
	if err != nil {
		log.Info().Err(err).Msg("Could not use url")
		return err
	}

	resp, err := client.httpClient.Post(addEventsUrl, "application/json", buf)

	if err != nil {
		log.Info().Err(err).Msg("Error during the request")
		return err
	}
	if resp.StatusCode != 200 {
		log.Info().Err(err).Msg("Got non 200 header")
		return fmt.Errorf("UNSUCCESSFUL REQUEST")
	}
	return nil
}
