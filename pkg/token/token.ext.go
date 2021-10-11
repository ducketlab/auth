package token

import "time"

func (t *Token) WithRemoteIp(ip string) {
	t.RemoteIp = ip
}

func (t *Token) WithUserAgent(ua string) {
	t.UserAgent = ua
}

func (t *Token) CheckAccessIsExpired() bool {
	if t.AccessExpiredAt == 0 {
		return false
	}

	return time.Unix(t.AccessExpiredAt/1000, 0).Before(time.Now())
}

func (t *Token) CheckRefreshIsExpired() bool {
	if t.RefreshExpiredAt == 0 {
		return false
	}

	return time.Unix(t.RefreshExpiredAt/1000, 0).Before(time.Now())
}

func (t *Token) Desensitize() {
	t.RefreshToken = ""
}