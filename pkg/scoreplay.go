package scoreplay

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
)

type SrData struct {
	CompetitionId string
	CompetitionDataset []Competition

	SeasonId string
	SeasonDataset []Season

	CompetitorId string
	CompetitorDataset []Competition
}

func Scoreplay(opts *Options) {
	fmt.Println("boom! I say!", opts, len(opts.Competition))

	// Get input: interactive competition > season > competitor > player
	// Or from CLI
	// Or read from file
	if (len(opts.Input) > 0) {

	} else {
		InteractiveFetchData(opts)
	}


	// Apply search if necessary

	// Write output if requested
}

func InteractiveFetchData(opts *Options) (err error) {
	var idRegex *regexp.Regexp
	var data SrData
	var route string
	baseRoute := opts.ApiRoute + "/" + opts.ApiEnv + "/" + opts.ApiVer + "/" + opts.ApiLoc + "/"

	// Competition
	idRegex, err = buildRegex("competition"); if err != nil {
		// todo err handle
		return
	}
	if (len(opts.Competition) > 0 && len(idRegex.FindString(opts.Competition)) > 0) {
		data.CompetitionId = opts.Competition
	} else {
		var payload *CompetitionResponse
		route = baseRoute + "competitions"
		payload, err = ApiCall[CompetitionResponse](route, opts.ApiKey); if err != nil {
			// todo error handling
			return
		}
		data.CompetitionId = InteractiveSelectData[Competition](payload.Competitions)
	}
	fmt.Println("COMPETITION SELECTED ", data.CompetitionId)

	// Season
	idRegex, err = buildRegex("season"); if err != nil {
		// todo err handle
		return
	}
	if (len(opts.Season) > 0 && len(idRegex.FindString(opts.Season)) > 0) {
		data.SeasonId = opts.Season
	} else {
		var payload *SeasonResponse
		route = baseRoute + "competitions/" + data.CompetitionId + "/seasons"
		payload, err = ApiCall[SeasonResponse](route, opts.ApiKey); if err != nil {
			// todo error handling
			return
		}
		data.SeasonId = InteractiveSelectData[Season](payload.Seasons)
	}
	fmt.Println("SEASON SELECTED ", data.SeasonId)

	// Season Competitor
	// idRegex, err = buildRegex("competitor"); if err != nil {
	// 	// todo err handle
	// 	return
	// }
	// if (len(opts.Competitor) > 0 && len(idRegex.FindString(opts.Competitor)) > 0) {
	// 	data.CompetitorId = opts.Competitor
	// } else {
	// 	var payload *CompetitorResponse
	// 	route = baseRoute + "seasons/" + data.SeasonId + "/competitors"
	// 	payload, err = ApiCall[CompetitorResponse](route, opts.ApiKey); if err != nil {
	// 		// todo error handling
	// 		return
	// 	}
	// 	data.CompetitorId = InteractiveSelectData[Competitor](payload.Competitors)
	// }
	// fmt.Println("COMPETITOR SELECTED ", data.CompetitorId)

	// Season Competitor Players
	idRegex, err = buildRegex("player"); if err != nil {
		// todo err handle
		return
	}
	if (len(opts.Player) > 0 && len(idRegex.FindString(opts.Player)) > 0) {
		data.PlayerId = opts.Player
	} else {
		var payload *PlayerResponse
		route = baseRoute + "seasons/" + data.SeasonId + "/competitor_players"
		payload, err = ApiCall[PlayerResponse](route, opts.ApiKey); if err != nil {
			// todo error handling
			return
		}
		data.PlayerId = InteractiveSelectData[Player](payload.Players)
	}
	fmt.Println("PLAYER SELECTED ", data.PlayerId)
	return
}

func InteractiveSelectData[T ScoreplayType] (data []T) string {
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
	} else {
		question += "\nPlease, select an option amongst the choices above, by typing the ID of your selected element (e.g 'sr:resource:17')"
		validInput := false
		for (validInput != true) {
			fmt.Println(question)
			answer, _ = reader.ReadString('\n')
			answer = answer[:len(answer)-1]
			for i := 0 ; i < len(data) && validInput != true ; i++ {
				validInput = answer == data[i].GetId()
			}
		}
	}

	return answer
}

func buildRegex(resource string) (reg *regexp.Regexp, err error) {
	return regexp.Compile("sr:" + resource + `:\d+`)
}
