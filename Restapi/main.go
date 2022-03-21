package main
	
import (
    "net/http"
	
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type msg string

func main() {
    e := echo.New()
	e.GET("/", hello)
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}