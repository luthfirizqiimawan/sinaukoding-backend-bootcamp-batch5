package main

import (
	"net/http"
	"strconv"

	_ "go-echo/docs"

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
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,min=0"`
}

var users = []User{
	{ID: 1, Name: "Agus", Age: 15},
	{ID: 2, Name: "Bagus", Age: 25},
	{ID: 3, Name: "Caca", Age: 29},
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the User API")
	})
	e.GET("/users", GetUsers)

	// /users/:id
	e.GET("/users/:id", GetUserByID)

	// update user
	e.PUT("/users/:id", UpdateUser)

	// delete user
	e.DELETE("/users/:id", DeleteUser)

	// insert user
	e.POST("/users", CreateUser)

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

	// Assign a new ID to the user (auto-increment based on current max ID)
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	newUser.ID = maxID + 1

	users = append(users, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

// UpdateUser godoc
// @Summary      Update existing user
// @Description  Updates user data for the given ID
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

	var updated User
	if err := c.Bind(&updated); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := c.Validate(&updated); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	for i, u := range users {
		if u.ID == idInt {
			// ensure ID remains the path ID
			updated.ID = idInt
			users[i] = updated
			return c.JSON(http.StatusOK, updated)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
}

// DeleteUser godoc
// @Summary      Delete user by ID
// @Description  Deletes a user by the given ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	for i, u := range users {
		if u.ID == idInt {
			// remove from slice
			users = append(users[:i], users[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
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
