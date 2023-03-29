package initializers

import "github.com/DeepjyotiSarmah/go_jwt_authentication_gin_gonic/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
