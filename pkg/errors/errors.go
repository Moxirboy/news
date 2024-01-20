package errors
//
//import (
//	"errors"
//	"net/http"

//)
//
//func Auth(err error) (int, *dto.BaseResponse) {
//	if err == nil {
//		return http.StatusForbidden, &dto.BaseResponse{
//			Status:  Forbidden,
//			Message: messages[Forbidden],
//		}
//	}
//
//	errStatus := status.Convert(err)
//
//	switch errStatus.Code() {
//	case codes.InvalidArgument:
//		return http.StatusForbidden, &dto.BaseResponse{
//			Status:  FieldInvalid,
//			Message: messages[FieldInvalid],
//		}
//	case codes.Unauthenticated:
//		return http.StatusUnauthorized, &dto.BaseResponse{
//			Status:  TokenInvalid,
//			Message: messages[TokenInvalid],
//		}
//	case codes.Canceled, codes.NotFound:
//		return http.StatusUnauthorized, &dto.BaseResponse{
//			Status:  TokenExpired,
//			Message: messages[TokenExpired],
//		}
//	}
//
//	return http.StatusBadGateway, &dto.BaseResponse{
//		Status:  Unknown,
//		Message: err.Error(),
//	}
//}
//
//func Parse(err error) (int, *dto.BaseResponse) {
//	switch {
//	case errors.Is(err, jwt.ErrExpValidation):
//		return http.StatusUnauthorized, &dto.BaseResponse{
//			Status:  TokenExpired,
//			Message: "access",
//		}
//	case errors.Is(err, jwt.ErrHMACVerification):
//		return http.StatusUnauthorized, &dto.BaseResponse{
//			Status:  TokenInvalid,
//			Message: "access",
//		}
//	}
//
//	s := status.Convert(err)
//
//	return 200, &dto.BaseResponse{
//		Message: s.Message(),
//	}
//}
