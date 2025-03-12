package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	generated "github.com/nickyrolly/tree-drone/generated"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) GetHello(ctx echo.Context, params generated.GetHelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostEstate(ctx echo.Context) error {
	var resp generated.EstatePostResponse

	newUUID := uuid.New()
	resp.Id = &newUUID // Convert UUID to string
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostEstateIdTree(ctx echo.Context, id uuid.UUID) error {
	var resp generated.EstateTreePostResponse

	newUUID := uuid.New()
	resp.Id = &newUUID // Convert UUID to string
	return ctx.JSON(http.StatusOK, resp)
}
