package main

import (
	"context"
	"developer/any/dal"
	database "developer/any/db"
	dbmodels "developer/any/dbmodels/models"
	"fmt"
	"math/rand"
)

// games := []string{"xkdk2edm", "m1zjmz60", "j1lq8v6g", "46wr4n1r", "46wlwl1r", "26840ydp", "9d3r97ed"}
func main() {
	ctx := context.Background()
	db := database.NewSQLiteDatabase()

	runDAL := dal.NewRunDAL(db.GetDb())
	runs, err := runDAL.GetAll(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printStruct(runs[0])

	// gameResp, err := speedrun.GetGames()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// gameDAL := dal.NewGameDAL(db.GetDb())

	// games, err := gameDAL.GetAll(ctx)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// err = gameDAL.Create(ctx, gameResp)
	// if err != nil {
	// 	return
	// }

	// spr, err := speedrun.GetRecords()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// randRecord := getRandomNumber(0, len(spr.Data))
	// selectedRecord := spr.Data[randRecord]

	// randRun := getRandomNumber(0, len(selectedRecord.Runs))
	// run := selectedRecord.Runs[randRun]

	// fmt.Println(selectedRecord.Category.Data.Name, run.Run.Values, selectedRecord.Variables)

	// for _, a := range selectedRecord.Variables.Data {
	// 	fmt.Println(a.Id, a.Name, a.Values.Values[run.Run.Values[a.Id]])
	// }

	// Add check for empty video links. Yes it can happen.
	// cmd := exec.Command("open", "-a", "Google Chrome", run.Run.Videos.Links[0].URI)
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// insertSQLite()
}

func setupTable(db database.Database) {
	// Migrate the schema
	if err := db.GetDb().AutoMigrate(&dbmodels.Run{}, &dbmodels.Game{}, &dbmodels.Variable{}, &dbmodels.Category{}); err != nil {
		panic("failed migration")
	}
}

func insertSQLite() {
	// db, err := sqlite.OpenDB("test.db")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // Migrate the schema
	// if err = db.AutoMigrate(&dbmodels.Run{}, &dbmodels.Game{}, &dbmodels.Variable{}, &dbmodels.Category{}); err != nil {
	// 	panic("failed migration")
	// }

	// g := dao.NewGameDAO(db)

	// c := &dao.CategoryDAO{
	// 	DB: db,
	// }

	// r := &dao.RunDAO{
	// 	DB: db,
	// }

	// v := &dao.VariablesDAO{
	// 	DB: db,
	// }

	// ctx := context.Background()
	// game := dbmodels.Game{Id: "m1zjmz60", Name: "Resident Evil 2"}
	// g.Create(ctx, game)

	// c.Create(ctx, dbmodels.Category{Id: "82466qnd", Name: "PC", Type: "per-game"})
	// c.Create(ctx, dbmodels.Category{Id: "jdz88r6d", Name: "GCN", Type: "per-game"})

	// run1 := dbmodels.Run{Id: "znj2xg7m", CategoryID: "82466qnd", GameID: "m1zjmz60", PlayerId: "48gn05yj", Level: nil,
	// 	VideoURI: "https://www.youtube.com/watch?v=I4gKsQM34mM", Runtime: 2915}

	// run2 := dbmodels.Run{Id: "me3ll79z", CategoryID: "jdz88r6d", GameID: "m1zjmz60", PlayerId: "48ggnp28", Level: nil,
	// 	VideoURI: "https://www.twitch.tv/videos/1504049347", Runtime: 3659}

	// runs := []dbmodels.Run{run1, run2}
	// r.Create(ctx, runs)

	// // Category: 82466qnd (PC)
	// v.Create(ctx, dbmodels.Variable{Id: "kn033dn3", RunID: "znj2xg7m", Name: "Scenario", LabelId: "5le3736q", Label: "Leon A"})
	// v.Create(ctx, dbmodels.Variable{Id: "789qw3lw", RunID: "znj2xg7m", Name: "Category", LabelId: "21dk9dpl", Label: "Any%"})
	// v.Create(ctx, dbmodels.Variable{Id: "jlz197n2", RunID: "znj2xg7m", Name: "Difficulty", LabelId: "gq750dy1", Label: "Normal"})

	// // Category jdz88r6d (GCN)
	// v.Create(ctx, dbmodels.Variable{Id: "kn033dn3", RunID: "me3ll79z", Name: "Scenario", LabelId: "5le3736q", Label: "Leon A"})
	// v.Create(ctx, dbmodels.Variable{Id: "789qw3lw", RunID: "me3ll79z", Name: "Category", LabelId: "21dk9dpl", Label: "Any%"})
}

func getRandomNumber(min, max int) int {
	if max == 0 {
		return 0
	}
	return min + rand.Intn(max-min)
}

// Choose random category
// func randCategory(game models.GamesResponse) models.CategoryData {
// 	total_categories := len(game.GamesData.Categories.Data)
// 	fmt.Printf("Total Categories %v\n", total_categories)
// 	random_cat_idx := getRandomNumber(0, total_categories)
// 	return game.GamesData.Categories.Data[random_cat_idx]
// }

// func randVariable(category models.CategoryData) (map[string]map[string]string, error) {
// 	var variables *models.Variable
// 	variable := make(map[string]map[string]string)

// 	//hunk_id := "xk986y20"
// 	//psn := "7kjpplzk"

// 	resp, err := http.Get(fmt.Sprintf("https://www.speedrun.com/api/v1/categories/%v/variables", category.Id))
// 	if err != nil {
// 		return variable, err
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return variable, err
// 	}

// 	if err := json.Unmarshal(body, &variables); err != nil {
// 		return variable, err
// 	}

// 	// Check for variables, not every game has one
// 	// I think I can actually do this before when I first
// 	// call the Game endpoint with "variables" embed so
// 	// I don't have to call this function at all.

// 	//var chosenVariables map[string]string
// 	chosenVariables := make(map[string]map[string]string)
// 	//printStruct(variables.Data)

// 	if len(variables.Data) != 0 {
// 		fmt.Println("===Variables===")
// 		for _, variableData := range variables.Data {
// 			var varKeys []string

// 			// get all keys
// 			for key := range variableData.Values.Variable {
// 				varKeys = append(varKeys, key)
// 			}

// 			// randomize key selection
// 			randKeyIdx := getRandomNumber(0, len(varKeys)-1)
// 			randKey := varKeys[randKeyIdx]

// 			// Get value
// 			value := variableData.Values.Variable[randKey]

// 			answer := make(map[string]string)
// 			answer[randKey] = value.Label

// 			chosenVariables[variableData.Id] = answer
// 		}
// 	} else {
// 		// RE2 Tofu/Hunk (xk986y20) doesn't have variables
// 		fmt.Println("TOFU!!")
// 	}

// 	printStruct(chosenVariables)
// 	return chosenVariables, nil
// }

// func buildURI(category string, variables map[string]map[string]string) (string, error) {
// 	// https://www.speedrun.com/api/v1/leaderboards/m1zjmz60/category/jdrxxwld?var-kn033dn3=0q5xjxv1&var-789qw3lw=5q8kjzyq&max=1&top=1

// 	uri := fmt.Sprintf("https://www.speedrun.com/api/v1/leaderboards/%v/category/%v", GAME_ID, category)

// 	if len(variables) == 0 {
// 		return uri, nil
// 	}

// 	uri = uri + "?"

// 	for key, value := range variables {
// 		for val_key := range value {
// 			uri += fmt.Sprintf("var-%v=%v&", key, val_key)
// 		}
// 	}

// 	uri = uri + "max=1&top=1"

// 	return uri, nil
// }

func printStruct(v interface{}) {
	fmt.Printf("%+v\n", v)
}
