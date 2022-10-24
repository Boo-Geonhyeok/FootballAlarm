package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	database "github.com/Boo-Geonhyeok/football-alarm/Database"
	models "github.com/Boo-Geonhyeok/football-alarm/Models"
)

func getFIxturesByLeague(leagueID int) {
	url := "https://v3.football.api-sports.io/fixtures?league=39&timezone=Asia%2FSeoul&season=2022"
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

	fixtureAPI := &models.FixtureAPI{}
	err = json.Unmarshal(body, fixtureAPI)
	if err != nil {
		log.Fatal("Can't Unmarshal a body: ", err)
		return
	}

	for _, f := range fixtureAPI.Response {
		type team struct {
			ID   int
			Name string
		}
		fixture := models.Fixture{
			ID:     f.Fixture.ID,
			Date:   f.Fixture.Date,
			League: f.League.Name,
			Home: team{
				ID:   f.Teams.Home.ID,
				Name: f.Teams.Home.Name,
			},
			Away: team{
				ID:   f.Teams.Away.ID,
				Name: f.Teams.Away.Name,
			},
			HomeGoal: f.Goals.Home,
			AwayGoal: f.Goals.Away,
		}

		database.DB.Create(&fixture)
	}
}

func GetAllFixtures() {
	leagueIDs := []int{1, 39, 140, 78, 135, 61, 2}
	for _, leagueID := range leagueIDs {
		getFIxturesByLeague(leagueID)
	}
}
