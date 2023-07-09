package scoreplay

import (
	"log"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

// Struct used to remember CLI options
type Options struct {
	// Used in CLI or interactive requested search
	Competition string `mapstructure:"COMPETITION_ID"`
	Season string `mapstructure:"SEASON_ID"`
	Competitor string `mapstructure:"COMPETITOR_ID"`
	Player string `mapstructure:"PLAYER_ID"`

	// Force input/output to this LocalFolder
	LocalFolder string `mapstructure:"LOCAL_FOLDER"`
	// Will we fetch from the API, or local read
	Input string
	// Will we local write
	Output string

	// api data
	ApiKey string `mapstructure:"API_KEY"`
	ApiRoute string `mapstructure:"API_ROUTE"`
	ApiEnv string `mapstructure:"API_ENV"`
	ApiVer string `mapstructure:"API_VER"`
	ApiLoc string `mapstructure:"API_LOC"`
}

func CLI(args []string) {
	opts, err := LoadConf(); if err != nil {
		log.Fatal("[CLI] Config load failure", err)
	}

	app := &cli.App{
		Name: "scoreplay",
		Usage: "Scoreplay Test App",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "competition",
				Usage: "Input a competition ID (sr:competition:17) or name to match",
				Destination: &opts.Competition,
				Value: opts.Competition,
			},
			&cli.StringFlag{
				Name: "season",
				Usage: "Input a season ID (sr:season:17) or name to match",
				Destination: &opts.Season,
				Value: opts.Season,
			},
			&cli.StringFlag{
				Name: "competitor",
				Usage: "Input a competitor ID (sr:competitor:17) or name to match",
				Destination: &opts.Competitor,
				Value: opts.Competitor,
			},
			&cli.StringFlag{
				Name: "output",
				Usage: "Required to save the resulting dataset into local storage",
				Destination: &opts.Output,
				Value: opts.Output,
			},
		},
		Action: func(*cli.Context) error {
			Scoreplay(opts)
			return nil
		},
	}

	if err := app.Run(args); err != nil {
		log.Fatal("[CLI] CLI app failure")
	}
}

func LoadConf() (*Options, error) {
	opts := Options{
		LocalFolder: "./storage",
		ApiKey: "",
		ApiRoute: "https://api.sportradar.com/soccer-extended",
		ApiEnv: "production",
		ApiVer: "v4",
		ApiLoc: "en",
	}

	// First read the configs file
	viper.AddConfigPath("./configs")
	viper.SetConfigName("api")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("[LoadConf] Reading config env file failure", err)
		return &opts, err
	}
	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&opts); err != nil {
		log.Println("[LoadConf] Unmarshal config env file failure", err)
		return &opts, err
	}

	// And overwrite with the .env file
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("[LoadConf] Reading local env file failure", err)
		return &opts, err
	}
	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&opts); err != nil {
		log.Println("[LoadConf] Unmarshal local env file failure", err)
		return &opts, err
	}

	return &opts, nil
}
