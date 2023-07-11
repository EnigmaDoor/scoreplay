# Scoreplay soccer API browser

This is a short project demonstrating Scoreplay Soccer API usage

## Installing and building the project
Clone the github repository, input your own Scoreplay API_KEY, build and run the project.
> git clone https://github.com/EnigmaDoor/scoreplay.git
> cd scoreplay  
> go build  
> echo "API_KEY=ENTER_YOUR_API_KEY_HERE" > .env  
> ./scoreplay --help  

Available is a .env.example file for an easy first-use example  

## Usage
After building the project, a binary is present to execute it. The main flow of the binary is to prompt the user successively to select a competition, a season and a competitor. There exists various ways and options to pre-emptively select or search for those resources. They can be seen in .env.example ; ./scoreplay --help ; and we provide a couple examples below:  
> ./scoreplay --competition sr:competition:23755  
> ./scoreplay --competition UEFA  
> ./scoreplay --competition sr:competition:23755 --season 22  
> ./scoreplay --competition sr:competition:23755 --season sr:season:89689 --competitor land  
> ./scoreplay --competition UEFA --output test.json  
> ./scoreplay --input test.json

Configuration priority (the prior is overwritten by the later): configs/api.env < .env < CLI flags  

## Structure
* pkg - Contains the code
* configs - Configuration folder, along with repo/.env
* storage - Local folder used for local storage functionalities

## Architecture
A simple architecture was selected. The usual code flow is as such:
* First loading the configuration (package viper) in cli.go
* Setting up the cli parser (package cli) in cli.go
* Executing the main code flow Scoreplay in scoreplay.go
* * This includes an initial local storage input read if required in localStorage.go
* * Followed by interactive fetch of datain scoreplay.go:
* * * For each resource (competition > season > competitor), we will request or use the configured resource option (ID or name) to determine if an API call is needed
* * * Once the resource is obtained (mainly from after an API call), we will automatically determine, or prompt the user to select the resource
* * Once all resources are obtained, we will display a the list of players in that competitor.
* * We finally save the search dataset locally in JSON format when required in localStorage.go
* Most types, structs and interfaces are defined scoreplayTypes.go (that 'condensed' file is only used due to the small scope of the project)

## Design Decisions
* Generics were used instead of a more traditional structuring mainly due to curiosity and to re-discover them.
* InteractiveSelectData was made to easily add a better, more interactive CLI.
* As this is a CLI program, the design reflect a "execute once and terminate" mentality. As such, there is no repeatable search functionality (better to re-execute the program) and most error handling will terminate as error recovery isn't needed.

## Possible ameliorations
I kept the scope small, being more interested in generics usage that another feature in this project. However, if we were to scale up, here's a few avenues of improvements, by order of priority:
* Test coverage. Unit test, notably on localStorage_test.go & apiCall_test.go. Whole CLI test through a testing package, executing the program with the exemples arguments and ensure the results are correct.
* Factorizing InteractiveFetchData to avoid code repetition. The idea is explained more throughly in scoreplay.go, the comment leading to func InteractiveFetchData. This would lead to cleaner code and far easier adding of another resource to manage.
* A (truly) interactive CLI prompt. When prompting the user for a choice, he can either arrow up or arrow down to select one of the options (5 displayed at any point, the selected + 2 below + 2 above), or write to automatically search within the dataset and only display matching names. ENTER to select. Arrow right could display more data related to the highlighted resource, and arrow left to go back to the selection menu.
* Fully implement --input, disabling API calls and instead reusing a previous search (not interesting without the interactive CLI above). The file reading is implemented.
* More statistics displayed, notably seasonal statistics for competitors and players
Others improvements are of course possible, for example: better local storage solutions : schema/types generation from Scoreplay openapi.yaml ; more testing coverage ; auto wait + retry on network failure
