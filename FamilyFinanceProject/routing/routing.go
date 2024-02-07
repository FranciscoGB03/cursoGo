package routing

import (
	"net/http"
	"strconv"

	"github.com/cursoGo/FamilyFinanceProject/common"
	"github.com/gin-gonic/gin"
)

// Declare global variables to store user and income data
var users []common.User
var incomes []common.Incoming

// Check if a user already exists based on first and last name
func CheckUserExists(user common.User) bool {
	for _, u := range users {
		if u.FirstName == user.FirstName && u.LastName == user.LastName {
			return true
		}
	}
	return false
}

// Update user's income and balance when a new income transaction is created
func UpdateNewIncomeTransaction(income *common.Incoming) bool {
	for i, u := range users {
		if u.ID == income.Account_ID {
			user := users[i]
			user.Income += income.Deposit
			user.Balance += income.Deposit
			users[i] = user
			income.Balance = user.Balance
			return true
		}
	}
	return false
}

// Update user's income and balance when an existing income transaction is updated
func UpdateIncomeTransaction(update *common.Incoming, old common.Incoming) bool {
	for i, u := range users {
		if u.ID == update.Account_ID {
			user := users[i]
			user.Income += (update.Deposit - old.Deposit)
			user.Balance += (update.Deposit - old.Deposit)
			users[i] = user
			update.Balance = user.Balance
			return true
		}
	}
	return false
}

// Update user's income and balance when an income transaction is deleted
func UpdateDeletedIncome(income common.Incoming) {
	for i, u := range users {
		if u.ID == income.Account_ID {
			user := users[i]
			user.Income -= income.Deposit
			user.Balance -= income.Deposit
			users[i] = user
			break
		}
	}
}

// Create routes for user-related functionality
func CreateUsersRouting(r *gin.Engine) {
	// Create User
	r.POST("/users", func(c *gin.Context) {
		var user common.User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if CheckUserExists(user) {
			c.AbortWithStatusJSON(http.StatusConflict, "user already exists")
			return
		}
		user.ID = len(users) + 1
		users = append(users, user)
		c.JSON(http.StatusCreated, user)
	})

	// Read all users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Read single user
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if strconv.Itoa(user.ID) == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Update user
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser common.User
		if err := c.BindJSON(&updatedUser); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		for i, user := range users {
			if strconv.Itoa(user.ID) == id {
				updatedUser.ID = user.ID
				users[i] = updatedUser
				c.JSON(http.StatusOK, updatedUser)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Delete user
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, user := range users {
			if strconv.Itoa(user.ID) == id {
				users = append(users[:i], users[i+1:]...)
				c.Status(http.StatusAccepted)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})
}

// Create routes for income-related functionality
func CreateIncomeRouting(r *gin.Engine) {
	// Create Income transaction
	r.POST("/incomes", func(c *gin.Context) {
		var income common.Incoming
		if err := c.BindJSON(&income); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if !UpdateNewIncomeTransaction(&income) {
			c.AbortWithStatusJSON(http.StatusConflict, "Account ID does not exist")
			return
		}
		income.Income_ID = len(incomes) + 1
		incomes = append(incomes, income)
		c.JSON(http.StatusCreated, income)
	})

	// Read all income transactions
	r.GET("/incomes", func(c *gin.Context) {
		c.JSON(http.StatusOK, incomes)
	})

	// Read single income transaction
	r.GET("/incomes/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, income := range incomes {
			if strconv.Itoa(income.Account_ID) == id {
				c.JSON(http.StatusOK, income)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Update particular income transaction
	r.PUT("/incomes/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updateIncome common.Incoming
		if err := c.BindJSON(&updateIncome); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		for i, income := range incomes {
			if strconv.Itoa(income.Income_ID) == id {
				updateIncome.Income_ID = income.Income_ID
				UpdateIncomeTransaction(&updateIncome, income)
				incomes[i] = updateIncome
				c.JSON(http.StatusOK, updateIncome)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Delete transaction
	r.DELETE("/incomes/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, income := range incomes {
			if strconv.Itoa(income.Income_ID) == id {
				deletedIncome := incomes[i]
				incomes = append(incomes[:i], incomes[i+1:]...)
				UpdateDeletedIncome(deletedIncome)
				c.Status(http.StatusNoContent)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})
}
