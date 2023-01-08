package bean

type VipLevelResponse struct {
	CurrentVipLevel vipLevel `json:"current_vip_level"`
	NextVipLevel    vipLevel `json:"next_vip_level"`
}

type vipLevel struct {
	Level                int     `json:"level"`
	MinimumTradingVolume int     `json:"minimum_trading_volume"`
	MinimumStakingVolume int     `json:"minimum_staking_volume"`
	MakerFee             float64 `json:"maker_fee"`
	TakerFee             float64 `json:"taker_fee"`
}
