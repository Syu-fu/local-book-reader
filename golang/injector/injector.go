package injector

import (
	"local-book-reader/domain/repository"
	"local-book-reader/infra"
	"local-book-reader/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

func InjectInmemBookRepository() *infra.Inmem {
	inmem := infra.NewInmem()
	return inmem
}

func InjectDBBookRepository() repository.BookRepository {
	sqlHandler := InjectDB()
	return infra.NewBookRepository(sqlHandler)
}

func InjectDBBookUsecase() usecase.BookUsecase {
	BookRepo := InjectDBBookRepository()
	return usecase.NewBookUsecase(BookRepo)
}

func InjectInmemBookUsecase() usecase.BookUsecase {
	BookRepo := InjectInmemBookRepository()
	return usecase.NewBookUsecase(BookRepo)
}
