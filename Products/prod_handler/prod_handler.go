package prod_handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SubhamMurarka/microService/Products/models"
	"github.com/SubhamMurarka/microService/Products/prod_service"
	"github.com/SubhamMurarka/microService/Products/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	prod_service.Service
}

func NewHandler(s prod_service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var p models.CreateProduct
	if err := c.BodyParser(&p); err != nil {
		fmt.Println("Error parsing request body for creating product : ", err)
		return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(err))
	}

	res, err := h.Service.CreateProduct(c.Context(), &p)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrEmptyFields):
			return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(err))
		default:
			return c.Status(http.StatusInternalServerError).JSON(utils.EncodeErrorResponse(err))
		}
	}

	return c.Status(http.StatusOK).JSON(utils.EncodeSuccessResponse(res))
}

func (h *Handler) GetProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	product, err := h.Service.GetProduct(c.Context(), productID)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrEmptyProductID):
			return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(err))
		default:
			return c.Status(http.StatusInternalServerError).JSON(utils.EncodeErrorResponse(err))
		}
	}

	return c.Status(http.StatusOK).JSON(utils.EncodeSuccessResponse(product))
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var updateProduct models.CreateProduct
	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(err))
	}

	err := h.Service.UpdateProduct(c.Context(), productID, &updateProduct)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrEmptyFields):
			return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(err))
		default:
			return c.Status(http.StatusInternalServerError).JSON(utils.EncodeErrorResponse(err))
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "product updated successfully"})
}

func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	err := h.Service.DeleteProduct(c.Context(), productID)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrEmptyProductID):
			return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(err))
		default:
			return c.Status(http.StatusInternalServerError).JSON(utils.EncodeErrorResponse(err))
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}

func (h *Handler) GetAllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, err := h.Service.GetAllProducts(c.Context(), page, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.EncodeErrorResponse(err))
	}

	return c.Status(http.StatusOK).JSON(utils.EncodeSuccessResponse(products))
}

func (h *Handler) Purchase(c *fiber.Ctx) error {
	var req models.PurchaseReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(err))
	}
	if len(req.ProductID) == 0 {
		return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(utils.ErrEmptyProductID))
	}
	event := models.KafkaEvent{
		UserID:    c.Locals("ID").(string),
		ProductID: req.ProductID,
	}
	fmt.Printf("my user id is : %v", event.UserID)
	res, err := h.Service.Purchase(c.Context(), &event)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.EncodeErrorResponse(err))
	}

	return c.Status(http.StatusOK).JSON(utils.EncodeSuccessResponse(res))
}
