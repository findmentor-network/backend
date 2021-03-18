package cmd

import (
	"fmt"
	_ "github.com/findmentor-network/backend/docs"
	"github.com/findmentor-network/backend/internal/healthcheck"
	"github.com/findmentor-network/backend/internal/person"
	"github.com/findmentor-network/backend/internal/person/controller"
	"github.com/findmentor-network/backend/pkg/echoextention"
	"github.com/findmentor-network/backend/pkg/log"
	mongohelper "github.com/findmentor-network/backend/pkg/mongoextentions"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"time"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "find mentor api",
}

// @title Find Mentor API
// @version 1.0
// @description Find Mentor API.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @schemes http
func init() {
	port := "5000"
	dbconn := "mongodb://root:example@127.0.0.1:27017/"
	dbName := "findmentor"
	rootCmd.AddCommand(apiCmd)
	rootCmd.Flags().StringVarP(&port, "port", "p", port, "service port")
	rootCmd.Flags().StringVarP(&dbconn, "dbconn", "c", dbconn, "Database connection string")
	rootCmd.Flags().StringVarP(&dbName, "dbname", "d", dbName, "Database name")

	apiCmd.RunE = func(cmd *cobra.Command, args []string) error {

		instance := echo.New()
		instance.HideBanner = true
		instance.HidePort = true
		instance.Logger = log.SetupLogger()
		echoextention.RegisterGlobalMiddlewares(instance)

		mongoDb, err := mongohelper.NewDatabase(dbconn, dbName)
		if err != nil {
			log.Logger.Fatalf("Failed to connect database. Error :%s", err.Error())
		}

		repository := person.NewRepository(mongoDb)
		apiController := controller.NewController(repository)
		controller.NewHandlers(instance, apiController)
		healthcheck.RegisterHandlers(instance, mongoDb)

		instance.GET("/swagger/*", echoSwagger.WrapHandler)
		go func() {
			if err := instance.Start(fmt.Sprintf(":%s", port)); err != nil {
				log.Logger.Fatalf("Failed to shutting down the service. Error: %s", err.Error())
			}

		}()

		echoextention.Shutdown(instance, time.Second*3)

		return nil
	}

}
