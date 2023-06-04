package middleware

import (
	"crm_service/entity"
	db2 "crm_service/utils/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vicanso/go-axios"
	"log"
)

func CustomerBulk(c *gin.Context) {
	db := db2.GormMysql()
	url := "https://reqres.in/api/users?page=2"

	// Send HTTP GET request using axios
	resp, err := axios.Get(url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the response body into a struct
	var responseData Response
	err = json.Unmarshal(resp.Data, &responseData)
	if err != nil {
		log.Fatal(err)
	}

	// Access the retrieved data
	for _, customer := range responseData.Data {
		var existingCustomer entity.Customer

		err := db.First(&existingCustomer, "email = ?", customer.Email).Error
		if err == nil {
			// FirstName already exists, return an error
			continue
		}
		// Email does not exist, proceed with creating the customer
		db.Table("customer").Create(&entity.Customer{FirstName: customer.FirstName, LastName: customer.LastName, Email: customer.Email, Avatar: customer.Avatar})

	}
	// Call c.Next() to pass control to the next middleware or route handler
	c.Next()
}
