package cmd

import (
	"fmt"
	"log"

	"github.com/apeiron242/todo/db"
	"github.com/apeiron242/todo/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show full list of Todo",
	Run: func(cmd *cobra.Command, args []string) {
		getList()
	},
}

func getList() {
	var dataList []models.Data
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	row, err := db.Query("SELECT * FROM data", "")
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		var id int
		var title string
		var details string
		var time string
		row.Scan(&id, &title, &details, &time)
		data := models.Data{Id: id, Title: title, Details: details, Time: time}
		dataList = append(dataList, data)
	}

	if len(dataList) == 0 {
		fmt.Println("\nNo saved Data\n")
		return
	}

	fmt.Printf("\nIndex\tTitle\tTime\t\t\tDetail\n\n")

	for index, elem := range dataList {
		var trimDetail string = elem.Details
		if len(trimDetail) > 30 {
			trimDetail = elem.Details[0:30]
		}
		fmt.Printf("%v\t%v\t%v\t%v\n", index+1, elem.Title, elem.Time, trimDetail)
	}
	fmt.Println("")
}

func init() {
	rootCmd.AddCommand(listCmd)
}
