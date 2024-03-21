package repository

import (
	"LauncherServer/internal/database"
	"context"
)

func GetUserBaseProfileByUserName(username string) (int32, string, bool, error) {
	user, err := database.Store.GetUserByUsername(context.Background(), username)
	if err != nil {
		return 0, "", false, err
	}
	return user.ID, user.Username, user.Isadmin, nil
}
