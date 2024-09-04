package services

import (
	"context"
	"errors"
	"fmt"
	"order-service-gb1/internal/model"
	"order-service-gb1/internal/utils"
	"order-service-gb1/messaging"
	"strconv"
	"strings"

	productPb "github.com/kodinggo/product-service-gb1/pb/product"
	"github.com/sirupsen/logrus"
)

type orderService struct {
	cartRepo      model.ICartsRepository
	orderRepo     model.OrderRepository
	productClient productPb.ProductServiceClient
	natsJS        messaging.JetStreamRepository
}

func NewOrderService(cartRepo model.ICartsRepository, orderRepo model.OrderRepository, productCLient productPb.ProductServiceClient, natsJS messaging.JetStreamRepository) model.OrderService {
	return &orderService{
		cartRepo:      cartRepo,
		orderRepo:     orderRepo,
		productClient: productCLient,
		natsJS:        natsJS,
	}
}

func (s *orderService) Create(ctx context.Context, order model.Order) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.Dump(ctx),
		"order": utils.Dump(order),
	})

	order.Status = "waiting_for_payment"

	// Generate invoice number if not exist
	if order.InvoiceNo == "" {
		order.InvoiceNo = utils.GenerateInvoice()
	}

	// TODO: get items/products from cart by user id
	// BLOCKED: waiting for auth implementation
	// carts, err := s.cartRepo.FindByUserID(ctx, order.UserID)
	// if err != nil {
	// 	logger.Error(err)
	// 	return err
	// }

	// if len(carts) == 0 {
	// 	return errors.New("cart is empty")
	// }

	// Get product information
	var productIDs []int32
	for _, item := range order.OrderItems {
		productIDs = append(productIDs, int32(item.ProductID))
	}

	request := &productPb.ProductRequest{
		Ids: productIDs,
	}

	resp, err := s.productClient.FindProductByIDs(ctx, request)
	if err != nil {
		logger.Error(err)
		return err
	}

	if len(resp.Products) == 0 {
		return errors.New("no product found")
	}

	// Set product information to order items
	for i, item := range order.OrderItems {
		for _, product := range resp.Products {
			if int(product.Id) == item.ProductID {
				if item.Qty > int(product.Stock) {
					return fmt.Errorf("product with id '%d' out of stock", item.ProductID)
				}

				order.OrderItems[i].ProductName = product.Name
				order.OrderItems[i].ProductPrice = product.Price
				order.OrderItems[i].SubTotal = product.Price * float64(item.Qty)
			}
		}
	}

	var notFoundIds []string
	for _, item := range order.OrderItems {
		if item.ProductName == "" {
			notFoundIds = append(notFoundIds, strconv.Itoa(item.ProductID))
		}
	}

	if len(notFoundIds) > 0 {
		return fmt.Errorf("product with id '%s' not found", strings.Join(notFoundIds, ", "))
	}

	order.SetGrandTotal()

	// Calculate total price
	err = s.orderRepo.Create(ctx, order)
	if err != nil {
		logger.Error(err)
		return err
	}

	var reservedRequest []*productPb.ReserveProduct
	for _, item := range order.OrderItems {
		reservedRequest = append(reservedRequest, &productPb.ReserveProduct{
			Id:  int32(item.ProductID),
			Qty: int32(item.Qty),
		})
	}

	reserveRequest := &productPb.ReserveProductRequest{
		Products: reservedRequest,
	}

	res, err := s.productClient.ReserveProduct(ctx, reserveRequest)
	if err != nil {
		logger.Error(err)
		return err
	}

	if res.Error != "" {
		logger.Error(res.Error)
		return errors.New(res.Error)
	}

	err = s.natsJS.Publish(ctx, "ORDER.create", order.ToJSON())
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
