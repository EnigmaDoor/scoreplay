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
}
func (r Competitor) GetId() string { return r.Id }
func (r Competitor) GetName() string { return r.Name }

type Player struct {
	Id string
	Name string
	Competitor *Competitor
}
func (r Player) GetId() string { return r.Id }
func (r Player) GetName() string { return r.Name }

type ScoreplayType interface {
	Competition | Season | Competitor | Player | Category
	GetId() string
	GetName() string
}

type ScoreplayResponseType interface {
	CompetitionResponse | SeasonResponse
}
