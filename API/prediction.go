package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	database "github.com/Boo-Geonhyeok/football-alarm/Database"
	models "github.com/Boo-Geonhyeok/football-alarm/Models"
)

func GetPredictions(fixtureID int) {
	ID := fmt.Sprint(fixtureID)
	url := "https://v3.football.api-sports.io/predictions?fixture=" + ID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal("Can't make a Request: ", err)
		return
	}
	req.Header.Add("x-apisports-key", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Can't Get a Respone: ", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Can't Read a Response:", err)
		return
	}

	predictionAPI := &models.PredictionAPI{}
	err = json.Unmarshal(body, predictionAPI)
	if err != nil {
		log.Fatal("Can't Unmarshal a body: ", err)
		return
	}

	for _, p := range predictionAPI.Response {
		type team struct {
			ID   int
			Name string
		}
		prediction := models.Prediction{
			Percent: struct {
				Home string
				Draw string
				Away string
			}(p.Predictions.Percent),
			Home:       team{ID: p.Teams.Home.ID, Name: p.Teams.Home.Name},
			Away:       team{ID: p.Teams.Away.ID, Name: p.Teams.Away.Name},
			Comparison: p.Comparison,
		}

		database.DB.Create(&prediction)
	}
}
