package models

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"net/url"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"log"
	"net/http"
)

//Match is object of Match
type JsonFormat struct {
    Api struct {
        Results int `json:"results"`
        Fixtures []Fixture `json:"fixtures`
    }
}

type Fixture struct {
    ID int `json:"fixture_id"`
    LeagueID int `json:"league_id"`
    HomeTeam struct {
        TeamID int `json:"team_id"`
        TeamName string `json:"team_name"`
        Logo string `json:"logo"`
    } `json:"homeTeam"`
    AwayTeam struct {
        TeamID int `json:"team_id"`
        TeamName string `json:"team_name"`
        Logo string `json:"logo"`
    } `json:"awayTeam"`
    HomeGoals int `json:"goalsHomeTeam"`
    AwayGoals int `json:"goalsAwayTeam"`
}

// Matches is array of Match
var Matches []Fixture

// GetByLeagueAndDate is to get matches by league and date
func GetByLeagueAndDate(leagueID int, date string, timezone string) []Fixture {
    // url format https://api-football-v1.p.rapidapi.com/v2/fixtures/league/524/2020-06-26?timezone=Asia/Jakarta
    err := godotenv.Load()
    if err != nil {
        log.Fatal("error loading .env")
    }

    date = "2020-09-10"
    reqURL, _ := url.Parse("https://api-football-v1.p.rapidapi.com/v2/fixtures/league/"+ strconv.Itoa(leagueID) +"/"+ date +"?timezone=" + timezone)
    fmt.Println(reqURL)
    request := &http.Request {
        Method: "GET",
        URL: reqURL,
        Header: map[string][]string {
            "X-RapidAPI-Host": {os.Getenv("RAPID_API_HOST")},
            "X-RapidAPI-KEY": {os.Getenv("RAPID_API_KEY")},
        },
    }

    res, err := http.DefaultClient.Do(request)
    if err != nil {
        log.Fatal(err)
    }

    data, _ := ioutil.ReadAll(res.Body)
    api := JsonFormat{}
    json.Unmarshal([]byte(data), &api)
    res.Body.Close()
    if json.Valid(data) {
        Matches = api.Api.Fixtures
        if(len(Matches) > 0) {
            fmt.Printf("%+v", Matches[0])
            fmt.Println("")
        }
    }
    return Matches
}