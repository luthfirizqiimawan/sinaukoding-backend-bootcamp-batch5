package main

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type User struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,min=0"`
}

var users = []User{
	{ID: 1, Name: "Alim", Age: 21},
	{ID: 2, Name: "Bagus", Age: 18},
	{ID: 3, Name: "Bambang", Age: 22},
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the User API")
	})
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUserByID)
	e.POST("/users", CreateUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Creates a new user with the provided details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      User  true  "User to create"
// @Success      201   {object}  User
// @Failure      400   {object}  map[string]string
// @Router       /users [post]
func CreateUser(c echo.Context) error {
	var newUser User

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := c.Validate(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Assign a new ID to the user
	users = append(users, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieves a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func GetUserByID(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id) // Convert string to int
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	for _, user := range users {
		if user.ID == idInt {
			c.Logger().Debug("Fetching user by ID")
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
}

// GetUsers godoc
// @Summary      Get all users
// @Description  Retrieves a list of all users
// @Tags         users
// @Produce      json
// @Success      200  {array}   User
// @Router       /users [get]
func GetUsers(c echo.Context) error {
	c.Logger().Debug("Fetching all users")
	return c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary      Update user by ID
// @Description  Updates an existing user's data by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int   true  "User ID"
// @Param        user  body      User  true  "Updated user data"
// @Success      200   {object}  User
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /users/{id} [put]
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := c.Validate(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	for i, user := range users {
		if user.ID == idInt {
			updatedUser.ID = idInt // Ensure ID stays the same
			users[i] = updatedUser
			return c.JSON(http.StatusOK, updatedUser)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
}

// DeleteUser godoc
// @Summary      Delete user by ID
// @Description  Deletes a user by their ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	for i, user := range users {
		if user.ID == idInt {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, echo.Map{"message": "User deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
}
