package scoreplay

import (
	"log"
	"fmt"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

// Struct used to remember CLI options
type Options struct {
	// Used in CLI or interactive requested search
	Competition string
	Season string
	Competitor string
	Player string

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
		// todo error handling
		return
	}

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

func LoadConf() (*Options, error) {
	var err error
	opts := Options{
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
		log.Fatal("Error reading env file", err)
	}
	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&opts); err != nil {
		log.Fatal(err)
	}

	// And overwrite with the .env file
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&opts); err != nil {
		log.Fatal(err)
	}

	fmt.Println("OPTIONS", opts)
	return &opts, err
}
