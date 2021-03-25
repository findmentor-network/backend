package cmd

import (
	"fmt"
	"github.com/findmentor-network/backend/internal/generate"
	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "Get the information of your beloved findmentor.network people",
}

func init() {
	SPREADSHEET_ID := "1x_W7Z2o_TGmEjL5cLTFbjO1R3KzQOqIhQKu9RQ4a_P4"
	API_KEY := "AIzaSyA5el9Fo8rMSYkcMjUqLfJi4tDB5_n0bzY"
	rootCmd.AddCommand(generatorCmd)

	generatorCmd.Flags().StringVarP(&SPREADSHEET_ID, "spreadsheet-id", "s", SPREADSHEET_ID, "spreadsheet id, get that from @cagataycali please don't ask me anymore")
	generatorCmd.Flags().StringVarP(&API_KEY, "api-key", "a", API_KEY, "github api key, everybody has one get yours")

	generatorCmd.RunE = func(cmd *cobra.Command, args []string) error {
		url := generateUrl(SPREADSHEET_ID, API_KEY)
		bundle := process(url)
		fmt.Println(bundle)
		return nil
	}
}

func generateUrl(spreadsheetId, apiKey string) string {
	return fmt.Sprint("https://sheets.googleapis.com/v4/spreadsheets/", spreadsheetId, "/values:batchGet?key=", apiKey, "&fields=valueRanges(range,values)&ranges=Mentees&ranges=Aktif%20Mentorluklar&ranges=Jobs&ranges=Interns")
}

func process(url string) generate.Bundle {
	res, err := generate.GetData(url)
	if err != nil {
		panic("failed getting data from google sheets")
	}

	// parsing
	bundle := generate.ParseBundle(&res)

	// post process
	bundle.AggregateMentorships()

	return bundle
}
