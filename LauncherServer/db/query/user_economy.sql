-- name: GetUserBalanceByUserId :one
SELECT * FROM user_economy WHERE user_id = ? LIMIT 1;