package scoreplay

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"regexp"
	"golang.org/x/exp/slices"
)

type SrData struct {
	CompetitionId string
	Competition *Competition
	CompetitionDataset []Competition

	SeasonId string
	Season *Season
	SeasonDataset []Season

	CompetitorId string
	Competitor *Competitor
	CompetitorDataset []Competitor
}

func Scoreplay(opts *Options) {

	// Get input: interactive competition > season > competitor > player
	// Or from CLI
	// Or read from file
	if (len(opts.Input) > 0) {

	} else {
		_, err := InteractiveFetchData(opts); if err != nil {
			log.Fatal("[Scoreplay] Fatal failure")
		}
	}

	// Write output if requested
}

func InteractiveFetchData(opts *Options) (*SrData, error) {
	var err error
	var idRegex *regexp.Regexp
	var data SrData
	var route string
	baseRoute := opts.ApiRoute + "/" + opts.ApiEnv + "/" + opts.ApiVer + "/" + opts.ApiLoc + "/"

	// Competition
	idRegex, err = buildRegex("competition"); if err != nil {
		log.Println("[InteractiveFetchData] buildRegex Competition Failure", err)
		return &data, err
	}
	if (len(opts.Competition) > 0 && len(idRegex.FindString(opts.Competition)) > 0) {
		data.CompetitionId = opts.Competition
		fmt.Println("Automatically selected competition " + data.CompetitionId)
	} else {
		var payload *CompetitionResponse
		route = baseRoute + "competitions"
		payload, err = ApiCall[CompetitionResponse](route, opts.ApiKey); if err != nil {
			log.Println("[InteractiveFetchData] ApiCall Competition Failure", err)
			return &data, err
		}
		data.Competition, err = InteractiveSelectData[Competition](payload.Competitions)
		data.CompetitionId = data.Competition.Id
		data.CompetitionDataset = payload.Competitions
		fmt.Println(data.Competition.Display())
	}

	// Season
	idRegex, err = buildRegex("season"); if err != nil {
		log.Println("[InteractiveFetchData] buildRegex Season Failure", err)
		return &data, err
	}
	if (len(opts.Season) > 0 && len(idRegex.FindString(opts.Season)) > 0) {
		data.SeasonId = opts.Season
		fmt.Println("Automatically selected season " + data.SeasonId)
	} else {
		var payload *SeasonResponse
		route = baseRoute + "competitions/" + data.CompetitionId + "/seasons"
		payload, err = ApiCall[SeasonResponse](route, opts.ApiKey); if err != nil {
			log.Println("[InteractiveFetchData] ApiCall Season Failure", err)
			return &data, err
		}
		data.Season, err = InteractiveSelectData[Season](payload.Seasons)
		data.SeasonId = data.Season.GetId()
		data.SeasonDataset = payload.Seasons
		fmt.Println(data.Season.Display())
	}

	// Season Competitor Players
	idRegex, err = buildRegex("competitor"); if err != nil {
		log.Println("[InteractiveFetchData] buildRegex Competitor Failure", err)
		return &data, err
	}
	if (len(opts.Competitor) > 0 && len(idRegex.FindString(opts.Competitor)) > 0) {
		data.CompetitorId = opts.Competitor
		fmt.Println("Automatically selected competitor " + data.CompetitorId)
	} else {
		var payload *CompetitorResponse
		route = baseRoute + "seasons/" + data.SeasonId + "/competitor_players"
		payload, err = ApiCall[CompetitorResponse](route, opts.ApiKey); if err != nil {
			log.Println("[InteractiveFetchData] ApiCall Competitor Failure", err)
			return &data, err
		}
		data.Competitor, err = InteractiveSelectData[Competitor](payload.Competitors)
		data.CompetitorId = data.Competitor.Id
		data.CompetitorDataset = payload.Competitors
		fmt.Println(data.Competitor.Display())
	}

	return &data, nil
}

func InteractiveSelectData[T ScoreplayType] (data []T) (selected *T, err error) {
	var question, answer string
	reader := bufio.NewReader(os.Stdin)

	for i := range data {
		question += fmt.Sprintf("%s => %s\n", data[i].GetId(), data[i].GetName())
	}
	if (len(data) == 0) {
		// todo error handling no result, abort
	} else if (len(data) == 1) {
		answer = data[0].GetId()
		question += "\nAutomatically selecting the only result available."
		fmt.Println(question)
		selected = &data[0]
	} else {
		question += fmt.Sprintf("\nPlease, select an option amongst the choices above, by typing the ID of your selected element (e.g '%s')", data[0].GetId())
		resourceIdx := -1
		for (resourceIdx == -1) {
			fmt.Println(question)
			answer, err = reader.ReadString('\n'); if err != nil {
				log.Println("[InteractiveSelectData] Reader failure", err)
			}
			answer = answer[:len(answer)-1]
			resourceIdx = slices.IndexFunc(data, func (d T) bool { return answer == d.GetId() })
		}
		selected = &data[resourceIdx]
	}

	return
}

func buildRegex(resource string) (reg *regexp.Regexp, err error) {
	return regexp.Compile("sr:" + resource + `:\d+`)
}
