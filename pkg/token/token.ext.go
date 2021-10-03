package token

func (t *Token) WithRemoteIp(ip string) {
	t.RemoteIp = ip
}

func (t *Token) WithUserAgent(ua string) {
	t.UserAgent = ua
}