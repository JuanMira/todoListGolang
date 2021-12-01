package middlewares

import "github.com/gofiber/fiber/v2/middleware/basicauth"

type CheckJWT struct {
	basicauth.Config
}
