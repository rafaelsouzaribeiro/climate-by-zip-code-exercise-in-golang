package di

import (
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/infra/web/handler"
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/infra/web/request"
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/infra/web/server"
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/usecase"
)

func Newid() *server.Server {
	server := server.New(":8080")
	request := request.NewRequest()
	usecases := usecase.NewClimateUseCase(request)
	handler := handler.NewClimateHandler(usecases)
	server.AddHandler("GET /climate/{cep}", handler.GetClimateByZipCode)

	return server
}
