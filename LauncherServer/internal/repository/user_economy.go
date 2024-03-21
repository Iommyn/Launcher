package repository

import (
	"LauncherServer/internal/database"
	"context"
)

func GetUserBalanceByUserID(userID int32) (int32, error) {
	userBalance, err := database.Store.GetUserBalanceByUserId(context.Background(), userID)
	if err != nil {
		return 0, err
	}
	return userBalance.Balance, nil
}
