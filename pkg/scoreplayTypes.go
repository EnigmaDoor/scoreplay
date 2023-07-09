package scoreplay

import (
	"fmt"
	"time"
	"math"
	"strings"
)

type ScoreplayTime struct {
	time.Time
}
func (t *ScoreplayTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type Category struct {
	Id string
	Name string
	CountryCode string
}
func (r Category) GetId() string { return r.Id }
func (r Category) GetName() string { return r.Name }

type Competition struct {
	Id string
	Name string
	Gender string
	Category Category
}
func (r Competition) GetId() string { return r.Id }
func (r Competition) GetName() string { return r.Name }
func (r Competition) Display() (str string) {
	str += fmt.Sprintf("Competition %s (%s)", r.Name, r.Gender)
	return
}


type CompetitionResponse struct {
	GeneratedAt string
	Competitions []Competition
}

type Season struct {
	Id string
	Name string
	StartDate string
	EndDate string
	Year string
	CompetitionId string
}
func (r Season) GetId() string { return r.Id }
func (r Season) GetName() string { return r.Name }
func (r Season) Display() (str string) {
	str += fmt.Sprintf("Season %s (%s)", r.Name, r.Year)
	return
}

type SeasonResponse struct {
	GeneratedAt string
	Seasons []Season
}

type Competitor struct {
	Id string
	Name string
	ShortName string
	Abbreviation string
	Players []Player
}
func (r Competitor) GetId() string { return r.Id }
func (r Competitor) GetName() string { return r.Name }
func (r Competitor) Display() (str string) {
	str += fmt.Sprintf("Competitor %s (%s) has the following players:\n", r.Name, r.Id)
	for i := range r.Players {
		str += r.Players[i].Display()
	}
	return
}

type CompetitorResponse struct {
	GeneratedAt string
	Competitors []Competitor `json:"season_competitor_players"`
}

type Player struct {
	Id string
	Name string
	Type string
	DateOfBirth ScoreplayTime `json:"date_of_birth"`
	Nationality string
	CountryCode string
	Height int
	Weight int
	JerseyNumber int
	PreferredFoot string
	PlaceOfBirth string
}
func (r Player) GetId() string { return r.Id }
func (r Player) GetName() string { return r.Name }
func (r Player) Display() (str string) {
	now := time.Now()
	str += fmt.Sprintf(
		"Player %s (age %d) occupies role %s\n",
		r.Name,
		int(math.Floor(r.DateOfBirth.Sub(now).Hours() / -24 / 365)), // todo cannot use custom type for sub
		// int(math.Floor(now.Sub(r.DateOfBirth).Hours() / 24 / 365)),
		r.Type,
	)
	return
}

type ScoreplayType interface {
	Competition | Season | Competitor | Player | Category
	GetId() string
	GetName() string
}

type ScoreplayResponseType interface {
	CompetitionResponse | SeasonResponse | CompetitorResponse
}
