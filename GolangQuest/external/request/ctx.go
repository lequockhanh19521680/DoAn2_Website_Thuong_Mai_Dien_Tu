package request

import (
	"encoding/json"
	"fmt"
	"github.com/eNViDAT0001/Backend/api"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type context struct {
	*gin.Context
}

func FromContext(c *gin.Context) *context {
	return &context{
		Context: c,
	}
}

func (s *context) Response(statusCode int, response interface{}) {
	s.JSON(statusCode, response)
}

func (s *context) Ok(data interface{}) {
	s.JSON(http.StatusOK, NewSuccessResponse(data))
}

func (s *context) Created(data interface{}) {
	s.JSON(http.StatusCreated, NewSuccessResponse(data))
}

func (s *context) Accepted(data interface{}) {
	s.JSON(http.StatusAccepted, NewSuccessResponse(data))
}

func (s *context) Success(statusCode int, data interface{}) {
	s.JSON(statusCode, NewSuccessResponse(data))
}

func (s *context) Err(statusCode int, errs ...error) {
	s.JSON(statusCode, NewErrResponse(errs...))
}

func (s *context) BadRequest(errs ...error) {
	s.Err(http.StatusBadRequest, errs...)
}

func (s *context) UnprocessableEntity(errs ...error) {
	s.Err(http.StatusUnprocessableEntity, errs...)
}

func (s *context) StatusBadGateway(errs ...error) {
	s.Err(http.StatusBadGateway, errs...)
}

func (s *context) NotFound(errs ...error) {
	s.Err(http.StatusNotFound, errs...)
}

func (s *context) Conflict(errs ...error) {
	s.Err(http.StatusConflict, errs...)
}

func (s *context) Forbidden(errs ...error) {
	s.Err(http.StatusForbidden, errs...)
}

func (s *context) InternalServer(errs ...error) {
	s.Err(http.StatusInternalServerError, errs...)
}

func (s *context) NoContent(errs ...error) {
	s.Err(http.StatusNoContent, errs...)
}

func (s *context) Unauthorized(errs ...error) {
	s.Err(http.StatusUnauthorized, errs...)
}

// BindError request error when parse json
func (s *context) BindError(err error) {
	key := ""
	if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
		key = unmarshalErr.Field
	}

	message := fmt.Sprintf("%s invalid params", key)

	s.BadRequest(api.NewBadRequestError(key, "", message))
}

func (s *context) OkPaging(paging paging.Paging, data interface{}) {
	s.JSON(http.StatusOK, NewSuccessResponse(data).WithPaging(paging))
}

func (s *context) OkIndicator(indicator *Indicator, data interface{}) {
	s.JSON(http.StatusOK, NewSuccessResponse(data).WithIndicator(indicator))
}

func (s *context) OkMetaData(meta *Meta, data interface{}) {
	s.JSON(http.StatusOK, NewSuccessResponse(data).WithMeta(meta))
}

// ResponseError handle error
func (s *context) ResponseError(errs ...error) {
	if len(errs) < 1 {
		s.JSON(http.StatusInternalServerError, NewErrResponse(api.NewInternalServerError("", "", "Oops.. Unexpected error occurred")))
		return
	}

	errResult := make([]error, 0)
	statusCode := http.StatusInternalServerError

	for i, err := range errs {
		if e, ok := err.(*api.Error); ok {
			if i == 0 {
				statusCode = e.StatusCode
			}

			errResult = append(errResult, e)
		} else {
			errResult = append(errResult, api.NewInternalServerError("", "", err.Error()))
		}
	}

	s.JSON(statusCode, NewErrResponse(errResult...))
}
