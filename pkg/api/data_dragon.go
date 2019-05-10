package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
)

type DataDragonClient struct {
	DataDragonVersion  string
	DataDragonLanguage string
	client             *http.Client
}

func (c DataDragonClient) GetChampions() (map[string]model.ChampionData, error) {
	var champions map[string]model.ChampionData
	if err := c.getInto(dataDragonDataURLFormat, "/champion.json", &champions); err != nil {
		return nil, err
	}
	return champions, nil
}

func (c DataDragonClient) GetChampion(name string) (*model.ChampionDataExtended, error) {
	var data map[string]*model.ChampionDataExtended
	if err := c.getInto(dataDragonDataURLFormat, fmt.Sprintf("/champion/%s.json", name), &data); err != nil {
		return nil, err
	}
	if data, ok := data[name]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("response does not contain requested champion data")
}

func (c DataDragonClient) GetProfileIcons() (map[string]model.ProfileIcon, error) {
	var icons map[string]model.ProfileIcon
	if err := c.getInto(dataDragonDataURLFormat, "/profileicon.json", &icons); err != nil {
		return nil, err
	}
	return icons, nil
}

func (c DataDragonClient) GetItems() (map[string]model.Item, error) {
	var icons map[string]model.Item
	if err := c.getInto(dataDragonDataURLFormat, "/item.json", &icons); err != nil {
		return nil, err
	}
	return icons, nil
}

func (c DataDragonClient) GetMasteries() (map[string]model.Mastery, error) {
	var icons map[string]model.Mastery
	if err := c.getInto(dataDragonDataURLFormat, "/mastery.json", &icons); err != nil {
		return nil, err
	}
	return icons, nil
}

func (c DataDragonClient) GetSummonerSpells() (map[string]model.SummonerSpell, error) {
	var icons map[string]model.SummonerSpell
	if err := c.getInto(dataDragonDataURLFormat, "/summoner.json", &icons); err != nil {
		return nil, err
	}
	return icons, nil
}

func (c DataDragonClient) getInto(format dataDragonURL, endpoint string, target interface{}) error {
	response, err := c.doRequest(format, endpoint)
	if err != nil {
		return err
	}
	var ddResponse model.DataDragonResponse
	if err := json.NewDecoder(response.Body).Decode(&ddResponse); err != nil {
		return err
	}
	data, err := json.Marshal(ddResponse.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &target)
}

func (c DataDragonClient) doRequest(format dataDragonURL, endpoint string) (*http.Response, error) {
	request, err := c.newRequest(format, endpoint)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("error response: %v", response.Status)
	}
	return response, nil
}

func (c DataDragonClient) newRequest(format dataDragonURL, endpoint string) (*http.Request, error) {
	var url string
	switch format {
	case dataDragonDataURLFormat:
		url = fmt.Sprintf(string(format), c.DataDragonVersion, c.DataDragonLanguage)
	case dataDragonImageURLFormat:
		url = fmt.Sprintf(string(format), c.DataDragonVersion)
	default:
		url = string(format)
	}
	url = "https://" + url + endpoint
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}
