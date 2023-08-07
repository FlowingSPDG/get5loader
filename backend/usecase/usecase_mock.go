package usecase

//go:generate mockgen -source=get_match.go -destination=mock/get_match.go
//go:generate mockgen -source=user_login.go -destination=mock/user_login.go
//go:generate mockgen -source=user_register.go -destination=mock/user_register.go
//go:generate mockgen -source=validate_jwt.go -destination=mock/validate_jwt.go
