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

func GetStanding(leagueID int) {
	ID := fmt.Sprint(leagueID)
	url := "https://v3.football.api-sports.io/standings?league=" + ID + "&season=2022"
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

	standingAPI := &models.StandingAPI{}
	err = json.Unmarshal(body, standingAPI)
	if err != nil {
		log.Fatal("Can't Unmarshal a body: ", err)
		return
	}

	for _, s := range standingAPI.Response[0].League.Standings[0] {
		standing := models.Standing{
			Rank:      s.Rank,
			TeamID:    s.Team.ID,
			Team:      s.Team.Name,
			Points:    s.Points,
			GoalsDiff: s.GoalsDiff,
			Played:    s.All.Played,
			Win:       s.All.Win,
			Draw:      s.All.Draw,
			Lose:      s.All.Played,
		}

		database.DB.Create(&standing)
	}
}
