package usecase

//go:generate mockgen -source=match.go -destination=mock/match.go
//go:generate mockgen -source=gameserver.go -destination=mock/gameserver.go
//go:generate mockgen -source=get5.go -destination=mock/get5.go
//go:generate mockgen -source=mapstats.go -destination=mock/mapstats.go
//go:generate mockgen -source=match.go -destination=mock/match.go
//go:generate mockgen -source=player.go -destination=mock/player.go
//go:generate mockgen -source=playerstat.go -destination=mock/playerstat.go
//go:generate mockgen -source=team.go -destination=mock/team.go
//go:generate mockgen -source=user.go -destination=mock/user.go
//go:generate mockgen -source=validate_jwt.go -destination=mock/validate_jwt.go
