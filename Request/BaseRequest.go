package Request

type BaseRequest interface{
	ToPKIRequestString() string
}
