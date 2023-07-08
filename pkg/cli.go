package scoreplay

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Struct used to remember CLI options
type Options struct {
	// Used in CLI or interactive requested search
	Competition string;
	Season string;
	Competitor string;
	Player string;

	// Will we fetch from the API, or local read
	Input string;
	// Will we local write
	Output string;

	// api data
	ApiKey string;
	ApiRoute string;
	ApiEnv string;
	ApiVer string;
	ApiLoc string;
}

func CLI(args []string) {
	opts := LoadConf()

	app := &cli.App{
		Name: "scoreplay",
		Usage: "Scoreplay Test App",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "competition",
				Usage: "Input a competition ID (sr:competition:17) or name to match",
				Destination: &opts.Competition,
			},
			&cli.StringFlag{
				Name: "season",
				Usage: "Input a season ID (sr:season:17) or name to match",
				Destination: &opts.Season,
			},
			&cli.StringFlag{
				Name: "competitor",
				Usage: "Input a competitor ID (sr:competitor:17) or name to match",
				Destination: &opts.Competitor,
			},
		},
		Action: func(*cli.Context) error {
			Scoreplay(opts)
			return nil
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Println("Error on CLI app running")
	}
}

func LoadConf() *Options {
	opts := Options{
		ApiKey: "",
		ApiRoute: "https://api.sportradar.com/soccer-extended",
		ApiEnv: "production",
		ApiVer: "v4",
		ApiLoc: "en",
	}
	return &opts
}
