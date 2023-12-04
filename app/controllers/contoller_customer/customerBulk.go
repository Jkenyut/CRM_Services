package contoller_customer

import (
	"context"
	"crm_service/app/model/origin"
	db2 "crm_service/utils/clients"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vicanso/go-axios"
	"net/http"
	"time"
)

func CustomerBulk(c *gin.Context) {
	db := db2.GormMysql()
	url := "https://reqres.in/api/users?page=2"

	// Send HTTP GET request using axios
	resp, err := axios.Get(url, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("uri error", http.StatusBadRequest))
		return
	}

	// Unmarshal the response body into a struct
	var responseUri ResponseBulkData
	err = json.Unmarshal(resp.Data, &responseUri)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("data decode error", http.StatusBadRequest))
		return
	}

	// Access the retrieved data
	for _, customer := range responseUri.Data {
		//timeout
		var cancel context.CancelFunc
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
		defer cancel()
		//query
		queryCreateBulkCustomer := "INSERT INTO contoller_customer(first_name, last_name, email, avatar)  SELECT ?,?,?,?  WHERE NOT EXISTS (SELECT email from contoller_customer where email=?)"
		result := db.WithContext(ctx).Exec(queryCreateBulkCustomer, customer.FirstName, customer.LastName, customer.Email, customer.Avatar, customer.Email)

		//check
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("failed exec query create bulk repository-entity_actor", http.StatusBadRequest))
			break
		}
	}
	// Call c.Next() to pass control to the next middleware or route handler
	c.Next()
}
