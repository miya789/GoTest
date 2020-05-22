package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Server struct {
	e   *echo.Echo
	cnf *Config
}

type Config struct {
	Port string
}

type InvalidConfig struct{}

func (*InvalidConfig) Error() string {
	return fmt.Sprintf("invalid config")
}

func New(cnf *Config) (*Server, error) {
	if cnf == nil {
		return nil, new(InvalidConfig)
	}

	svr := new(Server)
	// echo.New() initializes properly.
	e := echo.New()

	svr.e = setRoute(e)
	svr.cnf = cnf

	return svr, nil
}

func setRoute(e *echo.Echo) *echo.Echo {
	e.GET("/", mainFunc)
	return e
}

func (s *Server) listenAddr() string {
	return fmt.Sprintf(":%v", s.cnf.Port)
}

func (s *Server) Start() error {
	// s.e.Use(middleware.Recover())
	// s.e.Use(middleware.Logger())

	fmt.Printf("start listening server at %s\n", s.cnf.Port)
	return s.e.Start(s.listenAddr())
}

func mainFunc(context echo.Context) error {

	type Person struct {
		Age    int
		Height int
	}
	p := new(Person)
	p.Age = 26
	p.Height = 170
	str := fmt.Sprintf("年齢:%d\n身長:%d\n", p.Age, p.Height)

	return context.String(http.StatusOK, str)
}
