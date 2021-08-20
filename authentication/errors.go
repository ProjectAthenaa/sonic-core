package authentication

import "github.com/gin-gonic/gin"

var (
	ipDoesNotMatchSessionError = gin.H{"status": "ip_does_not_match_session"}
	unauthorizedError = gin.H{"status": "unauthorized"}
)
