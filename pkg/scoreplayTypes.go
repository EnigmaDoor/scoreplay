package scoreplay

import (
	"fmt"
	"time"
	"math"
	"strings"
)

// ScoreplayTime used to (un)marshalJSON times in Scoreplay API format
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
func (t ScoreplayTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02"))), nil
}

// Resource Category from Scoreplay API
type Category struct {
	Id string
	Name string
	CountryCode string
}
func (r Category) GetId() string { return r.Id }
func (r Category) GetName() string { return r.Name }

// Resource Competition from Scoreplay API
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

// Competition Response format from Scoreplay API
type CompetitionResponse struct {
	GeneratedAt string
	Competitions []Competition
}

// Resource Season from Scoreplay API
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

// Season Response format from Scoreplay API
type SeasonResponse struct {
	GeneratedAt string
	Seasons []Season
}

// Resource Competitor from Scoreplay API
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

// Competitor Response format from Scoreplay API
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
		int(math.Floor(r.DateOfBirth.Sub(now).Hours() / -24 / 365)), // todo cannot use custom type for sub. Instead, workaround with *-1
		// int(math.Floor(now.Sub(r.DateOfBirth).Hours() / 24 / 365)),
		r.Type,
	)
	return
}

// Interface for Scoreplay resources
type ScoreplayType interface {
	Competition | Season | Competitor | Player | Category
	GetId() string
	GetName() string
}

// Interface for Scoreplay responses
type ScoreplayResponseType interface {
	CompetitionResponse | SeasonResponse | CompetitorResponse
}
