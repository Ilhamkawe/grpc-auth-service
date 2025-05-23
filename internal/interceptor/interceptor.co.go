package interceptor

import (
	"context"
	"fmt"
	"github.com/Ilhamkawe/grpc-auth-service/internal/port"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type contextKey string

const CurrentUserKey = contextKey("currentUser")

func AuthUnaryInterceptor(rdb port.JwtRedisPort, authService port.AuthServicePort, jwtService port.JwtServicePort, protectedMethods map[string]bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		// Cek apakah method memerlukan otentikasi
		if _, ok := protectedMethods[info.FullMethod]; !ok {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(http.StatusUnauthorized, "Unauthorized")
		}

		authHeader := md["authorization"]

		if len(authHeader) == 0 || !strings.Contains(authHeader[0], "Bearer") {
			return nil, status.Errorf(http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := strings.TrimSpace(strings.Split(authHeader[0], " ")[1])
		token, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			return nil, status.Errorf(http.StatusUnauthorized, "Unauthorized")
		}

		_, err = rdb.GetKey(ctx, fmt.Sprintf("blacklist_token:%s", tokenString))

		if err == nil {
			return nil, status.Errorf(http.StatusUnauthorized, "Invalid Token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return nil, status.Errorf(http.StatusUnauthorized, "Unauthorized")
		}

		userIdFloat, ok := claims["user_id"].(float64)
		userID := int(userIdFloat)

		user, err := authService.GetUserByID(userID)
		user.Token = tokenString
		if err != nil {
			return nil, status.Errorf(http.StatusUnauthorized, "Unauthorized")
		}

		newCtx := context.WithValue(ctx, CurrentUserKey, user)

		return handler(newCtx, req)
	}

}
