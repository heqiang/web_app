package redis

const (
	// KeyPostTimeZSet 帖子及发帖时间
	KeyPostTimeZSet = "webApp:post:time"
	// KeyPostScoreZSet 帖子投票的分数
	KeyPostScoreZSet = "webApp:post:score"
	// 用户及帖子投票类型
	KeyPostVoted = "webApp:post:voted"
)
