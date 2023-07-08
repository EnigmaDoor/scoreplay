package scoreplay

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

type CompetitionResponse struct {
	GeneratedAt string
	Competitions []Competition
}

type Season struct {
	Id string
	Name string
	StartDate string
	EndDate string
	Year int
	CompetitionId string
}
func (r Season) GetId() string { return r.Id }
func (r Season) GetName() string { return r.Name }

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

type CompetitorResponse struct {
	GeneratedAt string
	Competitors []Competitor `json:"season_competitors"`
}

type Player struct {
	Id string
	Name string
	Type string
	DateOfBirth string
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

type PlayerResponse struct {
	GeneratedAt string
	Competitors []Competitor `json:"season_competitors"`
}

type ScoreplayType interface {
	Competition | Season | Competitor | Player | Category
	GetId() string
	GetName() string
}

type ScoreplayResponseType interface {
	CompetitionResponse | SeasonResponse | CompetitorResponse
}
