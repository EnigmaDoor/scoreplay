package scoreplay

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"regexp"
	"strings"
	"golang.org/x/exp/slices"
)

// Scoreplay Data structure, contains the search dataset and allows for
// easy (un)marshalling into JSON local files
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

// Main function after Config & CLI setup. Will deal with input/output matters
// And call on InteractiveFetchData
func Scoreplay(opts *Options) {
	var err error
	var data SrData

	// Ensure local storage folder
	err = EnsureLocalFolder(opts); if err != nil {
		log.Fatal("[Scoreplay] Fatal failure setting up local folder", err)
	}

	// if (len(opts.Input) > 0) {
	// 	err = ReadInput(opts, &data); if err != nil {
	// 		log.Fatal("[Scoreplay] Fatal failure on read input", err)
	// 	}
	// }

	err = InteractiveFetchData(opts, &data); if err != nil {
		log.Fatal("[Scoreplay] Fatal failure", err)
	}

	if (len(opts.Output) > 0) {
		err = WriteOutput(opts, &data); if err != nil {
			log.Fatal("[Scoreplay] Fatal failure on write output", err)
		}
	}
}

// TODO Factory to reduce repetition. Requires a GetPayload() on ScoreplayResponseType, and double generic (ScoreplayType and ScoreplayResponseType), a way to write in SrData generically. Example
// func GetRoute(baseRoute string, args []string) for ScoreplayType interface
// func InteractiveFetchData[T ScoreplayType, R ScoreplayResponseType] (opts *Options, data *SrData, resource string) (selected *T, dataset *[]T)
func InteractiveFetchData(opts *Options, data *SrData) (error) {
	var route string
	baseRoute := opts.ApiRoute + "/" + opts.ApiEnv + "/" + opts.ApiVer + "/" + opts.ApiLoc + "/"

	// Competition
	idRegex, err := buildRegex("competition"); if err != nil {
		log.Println("[InteractiveFetchData] buildRegex Competition Failure", err)
		return err
	}
	if (len(opts.Competition) > 0 && len(idRegex.FindString(opts.Competition)) > 0) {
		data.CompetitionId = opts.Competition
		fmt.Println("Automatically selected competition " + data.CompetitionId)
	} else {
		var competitions []Competition
		route = baseRoute + "competitions"
		payload, err := ApiCall[CompetitionResponse](route, opts.ApiKey); if err != nil {
			log.Println("[InteractiveFetchData] ApiCall Competition Failure", err)
			return err
		}
		if len(opts.Competition) > 0 {
			for i := range payload.Competitions {
				if strings.Contains(payload.Competitions[i].GetName(), opts.Competition) {
					competitions = append(competitions, payload.Competitions[i])
				}
			}
		} else {
			competitions = payload.Competitions
		}
		data.Competition, err = InteractiveSelectData[Competition](competitions); if err != nil {
			log.Println("[InteractiveFetchData] InteractiveSelectData Competition Failure", err)
			return err
		}
		data.CompetitionId = data.Competition.Id
		data.CompetitionDataset = payload.Competitions
		fmt.Println(data.Competition.Display())
	}

	// Season
	idRegex, err = buildRegex("season"); if err != nil {
		log.Println("[InteractiveFetchData] buildRegex Season Failure", err)
		return err
	}
	if (len(opts.Season) > 0 && len(idRegex.FindString(opts.Season)) > 0) {
		data.SeasonId = opts.Season
		fmt.Println("Automatically selected season " + data.SeasonId)
	} else {
		var seasons []Season
		route = baseRoute + "competitions/" + data.CompetitionId + "/seasons"
		payload, err := ApiCall[SeasonResponse](route, opts.ApiKey); if err != nil {
			log.Println("[InteractiveFetchData] ApiCall Season Failure", err)
			return err
		}
		if len(opts.Season) > 0 {
			for i := range payload.Seasons {
				if strings.Contains(payload.Seasons[i].GetName(), opts.Season) {
					seasons = append(seasons, payload.Seasons[i])
				}
			}
		} else {
			seasons = payload.Seasons
		}
		data.Season, err = InteractiveSelectData[Season](seasons); if err != nil {
			log.Println("[InteractiveFetchData] InteractiveSelectData Season Failure", err)
			return err
		}
		data.SeasonId = data.Season.GetId()
		data.SeasonDataset = payload.Seasons
		fmt.Println(data.Season.Display())
	}

	// Season Competitor Players
	idRegex, err = buildRegex("competitor"); if err != nil {
		log.Println("[InteractiveFetchData] buildRegex Competitor Failure", err)
		return err
	}
	if (len(opts.Competitor) > 0 && len(idRegex.FindString(opts.Competitor)) > 0) {
		data.CompetitorId = opts.Competitor
		fmt.Println("Automatically selected competitor " + data.CompetitorId)
		// As this is the final resource endpoint (for now), TODO implement GET on route
		// baseRoute + "seasons/" + data.SeasonId + "/competitors/" + data.CompetitorId + "/statistics"
	} else {
		var competitors []Competitor;
		route = baseRoute + "seasons/" + data.SeasonId + "/competitor_players"
		payload, err := ApiCall[CompetitorResponse](route, opts.ApiKey); if err != nil {
			log.Println("[InteractiveFetchData] ApiCall Competitor Failure", err)
			return err
		}
		if len(opts.Competitor) > 0 {
			for i := range payload.Competitors {
				if strings.Contains(payload.Competitors[i].GetName(), opts.Competitor) {
					competitors = append(competitors, payload.Competitors[i])
				}
			}
		} else {
			competitors = payload.Competitors
		}
		data.Competitor, err = InteractiveSelectData[Competitor](competitors); if err != nil {
			log.Println("[InteractiveFetchData] InteractiveSelectData Competitor Failure", err)
			return err
		}
		data.CompetitorId = data.Competitor.Id
		data.CompetitorDataset = payload.Competitors
		fmt.Println(data.Competitor.Display())
	}

	return nil
}

// InteractiveSelectData will list the possibles choices, present in data.
// If a single choice is present, it is automatically selected.
// Otherwise, prompt the user until an ID matching in the dataset is selected.
func InteractiveSelectData[T ScoreplayType] (data []T) (selected *T, err error) {
	var question, answer string
	reader := bufio.NewReader(os.Stdin)

	for i := range data {
		question += fmt.Sprintf("%s => %s\n", data[i].GetId(), data[i].GetName())
	}
	if (len(data) == 0) {
		log.Fatal("[InteractiveSelectData] No result found. Please verify your search and try again.")
	} else if (len(data) == 1) {
		answer = data[0].GetId()
		question += fmt.Sprintf("Automatically selecting the only result available: %s (%s)\n", data[0].GetName(), data[0].GetId())
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

// Helper func to ensure the string format is of a correct Scoreplay ID format
func buildRegex(resource string) (reg *regexp.Regexp, err error) {
	return regexp.Compile("sr:" + resource + `:\d+`)
}
