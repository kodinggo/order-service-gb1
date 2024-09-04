package console

import (
	"order-service-gb1/db/db"
	"order-service-gb1/internal/handler"
	"order-service-gb1/internal/repository"
	"order-service-gb1/internal/services"
	"order-service-gb1/internal/utils"

	"order-service-gb1/messaging"

	"github.com/go-playground/validator"
	productPb "github.com/kodinggo/product-service-gb1/pb/product"
	authPb "github.com/kodinggo/user-service-gb1/pb/auth"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/nats-io/nats.go"
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

	authConn, err := grpc.NewClient("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	continueOrFatal(err)

	defer authConn.Close()

	productConn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	continueOrFatal(err)

	defer productConn.Close()

	auth := authPb.NewJWTValidatorClient(authConn)
	productClient := productPb.NewProductServiceClient(productConn)

	nc, err := nats.Connect(nats.DefaultURL)
	continueOrFatal(err)
	defer nc.Close()

	jetstream, err := messaging.NewJetStreamRepository(nc)
	continueOrFatal(err)

	err = jetstream.AddStream(nil, "ORDERS", []string{"ORDER.create"})
	continueOrFatal(err)

	e := echo.New()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}

	cartRepo := repository.NewCartsRepository(msql)
	orderRepo := repository.NewOrderRepository(msql)

	cartService := services.NewCartsRepository(cartRepo)
	orderService := services.NewOrderService(cartRepo, orderRepo, productClient, jetstream)

	authMiddleware := utils.NewJWTMiddleware(auth)

	handler := handler.NewHandler()
	handler.RegisterCartsServices(cartService)
	handler.RegisterAuthClient(auth)
	handler.RegisterOrderService(orderService)
	handler.Routes(e, authMiddleware.ValidateJWT)

	err = e.Start(":3232")
	continueOrFatal(err)
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
