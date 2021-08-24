package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apeiron242/todo/db"
	"github.com/apeiron242/todo/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new Todo",
	Run: func(cmd *cobra.Command, args []string) {
		addData(args)
	},
}

func addData(args []string) {
	var title string
	var details string
	time := time.Now().Format("2006-01-02 15:04:05")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nWhat is the title? : ")
	scanner.Scan()
	title = scanner.Text()
	fmt.Printf("Write some Details : ")
	scanner.Scan()
	details = scanner.Text()

	data := models.Data{Title: title, Details: details, Time: time}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO data (title, details, time) VALUES (?, ?, ?)", data.Title, data.Details, time)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Successfully added %v\n", title)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
