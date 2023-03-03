package response

type AppHeaderAuthorization struct {
	Authorization string `header:"Authorization,required"`
}
