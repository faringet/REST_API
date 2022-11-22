package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if errors.As(err, &appErr) { // смотрим наша ли эта ошибка (apperror), то пишем 418 ошибку
				if errors.Is(err, ErrNotFound) { //сравнение ошибки с моими ошибками
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}

			w.WriteHeader(http.StatusTeapot)    // почему 418 ? если выбрасывает 418 то косяк в коде
			w.Write(systemError(err).Marshal()) // получаем все системные ошибки красиво обернутые
		}
	}
}
