package loader

import (
	"bufio"
	"context"
	"developer/any/clients/speedrun"
	"developer/any/dal"
	database "developer/any/db"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
)

func LoadRecords(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db := database.NewSQLiteDatabase()
	categoryDAL := dal.NewCategoryDAL(db.GetDb())
	runDAL := dal.NewRunDAL(db.GetDb())

	var wg sync.WaitGroup

	// logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// we need to call SQLite to get the games
	// so we can load their respective records
	gameDAL := dal.NewGameDAL(db.GetDb())
	games, err := gameDAL.GetAll(ctx)
	if err != nil {
		return err
	}

	// TODO: Maybe add this as a config?
	// Speedrun API rate limit: 100 requests per minute
	// https://github.com/speedruncomorg/api/blob/master/throttling.md
	requestsPerMinute := 100
	duration := 60 * time.Second
	rateLimit := rate.Limit(requestsPerMinute) / rate.Limit(duration.Seconds())
	limiter := rate.NewLimiter(rateLimit, 1)

	// open success file
	successFile, err := openFile("cmd/loader/loader/success.txt", Success)
	if err != nil {
		return err
	}
	defer successFile.Close()

	// open error file
	errorFile, err := openFile("cmd/loader/loader/error.txt", Error)
	if err != nil {
		return err
	}
	defer errorFile.Close()

	// put every game in the file in a map as key
	loadedGames := map[string]struct{}{}
	scanner := bufio.NewScanner(successFile)
	for scanner.Scan() {
		ln := scanner.Text()
		loadedGames[ln] = struct{}{}
	}

	// loop through every game
	start := time.Now()
	for _, game := range games {

		if _, ok := loadedGames[game.Name]; ok {
			log.Info().Str("gameId:", game.Id).Str("game:", game.Name).Msg("game already loaded. skipping...")
			continue
		}

		// rate limiting wait until allowed to send request
		if err := limiter.Wait(ctx); err != nil {
			log.Error().AnErr("loader.LoadRecords ", err).Msg("error")
			return err
		}

		log.Info().Str("gameId:", game.Id).Str("game: ", game.Name).Msg("loading game...")
		records, err := speedrun.GetRecords(game)
		if err != nil {
			log.Error().AnErr("speedrun.GetRecords ", err).Msg("error")
			return err
		}

		if len(records.Data) != 0 {
			for _, record := range records.Data {
				// category
				categoryId, err := categoryDAL.Create(ctx, game.Id, record.Category)
				if err != nil {
					return err
				}
				log.Debug().Str("categoryId", record.Category.Data.Id).Str("category:", record.Category.Data.Name).Msg("loaded category")

				// runs
				for _, run := range record.Runs {
					if err := runDAL.Create(ctx, game.Id, categoryId, run.Run); err != nil {
						go writer(&wg, errorFile, game.Name, run.Run.Id, Error, err.Error())
						continue
					}
				}
			}

			go writer(&wg, successFile, game.Name, "", Success, "")
			log.Info().Str("gameId:", game.Id).Str("game:", game.Name).Msg("loaded game!")
		} else {
			log.Info().Str("gameId:", game.Id).Str("game:", game.Name).Msg("NO RUNS!")
		}
	}

	wg.Wait()
	elapsed := time.Since(start)
	log.Info().Str("time elapsed to process 300 requests", elapsed.String()).Msg("time")
	return nil
}

func openFile(filePath string, fileType FileType) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Error().AnErr("loader.LoadRecords ", err).Msg(fmt.Sprintf("Failed to open %v file", fileType))
		return nil, err
	}

	return file, nil
}

// TODO: Need to handle error better
func writer(wg *sync.WaitGroup, file *os.File, game, runId string, fileType FileType, errMsg string) {
	defer wg.Done()
	wg.Add(1)

	if fileType == Error {
		if err := writeToErrorFile(file, game, runId, errMsg); err != nil {
			return
		}
	} else {
		if err := writeToSuccessFile(file, game); err != nil {
			return
		}
	}
}

func writeToSuccessFile(file *os.File, game string) error {
	_, err := file.WriteString(game + "\n")
	if err != nil {
		log.Error().AnErr("loader.LoadRecords ", err).Msg("error writing to success file")
		return err
	}

	return nil
}

func writeToErrorFile(file *os.File, game, runId string, errorMsg string) error {
	_, err := file.WriteString(game + " " + runId + " " + errorMsg + "\n")
	if err != nil {
		log.Error().AnErr("loader.LoadRecords ", err).Msg("error writing to error file")
		return err
	}

	return nil
}
