package handler

import (
	"fmt"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Create a new order
func (h *Handler) addOrder(c echo.Context) (err error) {
	r := new(models.Order)
	return h.Add(c, r, func(r models.Model) (string, error) {
		return h.orderStore.Insert(r.(*models.Order), "UID_HERE")
	})
}

func (h *Handler) getOrders(c echo.Context) (err error) {
	return h.getFiltered(c, h.orderStore.GetAll)
}

func (h *Handler) getUserOrders(c echo.Context) (err error) {
	m := map[string]interface{}{}
	if err = c.Bind(&m); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	if m["uid"] == nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	l, err := h.orderStore.GetByUser(m["uid"].(string))

	fmt.Println(l)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	if len(l) <= 0 {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "User dont have any orders",
		})
	}
	return c.JSON(http.StatusOK, l)
}
