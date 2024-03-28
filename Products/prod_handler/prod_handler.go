package prod_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SubhamMurarka/microService/Products/models"
	"github.com/SubhamMurarka/microService/Products/prod_service"
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
	var p models.Product
	if err := c.BodyParser(&p); err != nil {
		fmt.Println("Error parsing request body for creating product : ", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.Service.CreateProduct(c.Context(), &p)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) GetProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	product, err := h.Service.GetProduct(c.Context(), productID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(product)
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var updateProduct models.Product
	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.UpdateProduct(c.Context(), productID, &updateProduct)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "product updated successfully"})
}

func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	err := h.Service.DeleteProduct(c.Context(), productID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}

func (h *Handler) GetAllProducts(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page number"})
	}

	products, err := h.Service.GetAllProducts(c.Context(), page)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(products)
}

func (h *Handler) Purchase(c *fiber.Ctx) error {
	var req models.PurchaseReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := h.Service.Purchase(c.Context(), &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(res)
}
