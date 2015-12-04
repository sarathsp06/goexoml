package goexoml

//SetSendDigits sets SendDigits
func (__number__ *Number) SetSendDigits(senddigits string) *Number {
	__number__.SendDigits = senddigits
	return __number__
}

//SetURL sets URL
func (__number__ *Number) SetURL(url string) *Number {
	__number__.URL = url
	return __number__
}

//SetMethod sets Method
func (__number__ *Number) SetMethod(method string) *Number {
	__number__.Method = method
	return __number__
}

//SetStatusCallbackEvent sets StatusCallbackEvent
func (__number__ *Number) SetStatusCallbackEvent(statuscallbackevent string) *Number {
	__number__.StatusCallbackEvent = statuscallbackevent
	return __number__
}

//SetStatusCallback sets StatusCallback
func (__number__ *Number) SetStatusCallback(statuscallback string) *Number {
	__number__.StatusCallback = statuscallback
	return __number__
}

//SetStatusCallbackMethod sets StatusCallbackMethod
func (__number__ *Number) SetStatusCallbackMethod(statuscallbackmethod string) *Number {
	__number__.StatusCallbackMethod = statuscallbackmethod
	return __number__
}

//SetNoun sets Noun
func (__number__ *Number) SetNoun(noun string) *Number {
	__number__.Noun = noun
	return __number__
}

//NewNumber return a new Number pointer
func NewNumber() *Number {
	return new(Number)
}

//AddNumber appends the verb to response
func (r *Response) AddNumber(number Number) *Response {
	r.Response = append(r.Response, number)
	return r
}