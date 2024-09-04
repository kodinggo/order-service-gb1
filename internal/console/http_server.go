package console

import (
	"order-service-gb1/db/db"
	"order-service-gb1/internal/handler"
	"order-service-gb1/internal/repository"
	"order-service-gb1/internal/services"
	"order-service-gb1/internal/utils"

	"github.com/go-playground/validator"
	authPb "github.com/kodinggo/user-service-gb1/pb/auth"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "Start the HTTP server",
	Run:   httpServer,
}

func httpServer(cmd *cobra.Command, args []string) {
	msql := db.ConfigDB()

	db, err := msql.DB()
	continueOrFatal(err)

	defer db.Close()

	conn, err := grpc.NewClient("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	continueOrFatal(err)

	defer conn.Close()

	auth := authPb.NewJWTValidatorClient(conn)

	continueOrFatal(err)

	e := echo.New()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}

	cartRepo := repository.NewCartsRepository(msql)

	cartService := services.NewCartsRepository(cartRepo)

	authMiddleware := utils.NewJWTMiddleware(auth)

	handler := handler.NewcartsHandler()
	handler.RegisterCartsServices(cartService)
	handler.RegisterAuthClient(auth)
	handler.Routes(e, authMiddleware.ValidateJWT)

	err = e.Start(":3232")
	continueOrFatal(err)
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
