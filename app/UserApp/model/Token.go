package UserAppModel

type Token struct {
	UserId   int64  `json:"UserId"`
	TenantId int64  `json:"TenantId"`
	Mobile   string `json:"Mobile"`
	Iss      string `json:"iss"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}
