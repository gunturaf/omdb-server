package omdbservice

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository"
)

var (
	ErrNoData = errors.New("no data")
)

type OMDBServiceImpl struct {
	client  repository.HTTPClient
	baseURL string
	apiKey  string
}

func NewOMDBService(client repository.HTTPClient, baseURL string, apiKey string) OMDBServiceImpl {
	return OMDBServiceImpl{
		client:  client,
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (s OMDBServiceImpl) httpGetByte(ctx context.Context, endpoint string) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(r)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (s OMDBServiceImpl) Search(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
	qs := url.Values{
		"apikey": []string{s.apiKey},
		"s":      []string{text},
		"page":   []string{strconv.Itoa(int(page))},
	}

	endpoint := s.baseURL + "?" + qs.Encode()

	body, err := s.httpGetByte(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var out entity.OMDBSearchResult

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}

	if out.Response == "False" {
		return nil, ErrNoData
	}

	return &out, nil
}

func (s OMDBServiceImpl) GetByID(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
	qs := url.Values{
		"apikey": []string{s.apiKey},
		"i":      []string{id},
	}

	endpoint := s.baseURL + "?" + qs.Encode()

	body, err := s.httpGetByte(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var out entity.OMDBResultSingle

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}

	if out.Response == "False" {
		return nil, ErrNoData
	}

	return &out, nil
}
