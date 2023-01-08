package usecase

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Backend/api"
	ioSto "github.com/eNViDAT0001/Backend/internal/verification/domain/jwt/storage/io"
	ioUC "github.com/eNViDAT0001/Backend/internal/verification/domain/jwt/usecase/io"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtUseCase) RefreshToken(ctx context.Context, refreshToken string) (*ioUC.JwtToken, error) {
	token, err := s.tokenSto.VerifyToken(refreshToken)
	if err != nil {
		return nil, api.NewUnauthorizedError("refresh_token", refreshToken, "Invalid Refresh Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, api.NewUnauthorizedError("refresh_token", refreshToken, "Invalid Token")
	}

	userID, _ := strconv.Atoi(claims["user_id"].(string))
	user, err := s.userSto.GetUserDetailByID(ctx, uint(userID))

	if err != nil {
		return nil, api.NewUnauthorizedError("refresh_token", refreshToken, "Invalid Refresh Token")
	}

	baseToken := ioSto.GenerateTokenInput{
		UserID:   strconv.Itoa(int(user.ID)),
		Username: *user.Username,
		Salt:     *user.Salt,
	}

	minutes, _ := ioUC.GetMinute()

	accessToken, err := s.tokenSto.GenerateToken(baseToken, minutes)
	if err != nil {
		return nil, err
	}

	output := ioUC.JwtToken{
		AccessToken:        accessToken.Token,
		AccessTokenExpiry:  accessToken.TokenExpiry,
		RefreshToken:       refreshToken,
		RefreshTokenExpiry: int64(claims["exp"].(float64)),
	}

	return &output, nil
}
