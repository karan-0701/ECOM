package cart

import (
	"fmt"

	"github.com/karan-0701/ecom/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}

		productIds[i] = item.ProductID
	}

	return productIds, nil
}

func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, userID int) (prodQuant int, totalAmt float64, err error) {
	// create a map of type products
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}
	// calculate the total price
	totalPrice := calculateTotalPrice(productMap, items)
	// reduce quantity product in db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity
		h.store.UpdateProduct(product)
	}
	// create the order
	OrderID, err := h.orderStore.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, err
	}
	// create order items
	for _, item := range items {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}
	return OrderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	// check if the number of items are zero
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}
	for _, item := range cartItems {
		product, ok := products[item.ProductID]

		if !ok {
			return fmt.Errorf("product %d is not available in the store, please check again", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product is not available in the quantity %d", product.Quantity)
		}

	}
	return nil
}

func calculateTotalPrice(products map[int]types.Product, cartItems []types.CartItem) float64 {
	var totalPrice float64
	for _, item := range cartItems {
		product := products[item.ProductID]
		totalPrice += product.Price * float64(item.Quantity)
	}
	return totalPrice
}
