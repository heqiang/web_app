package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 423 //每一票的分数
)

var (
	ctx               = context.Background()
	ErrVoteTimeExpire = errors.New("超出投票时间")
)

func VoteForPost(userid, postId string, value float64) (err error) {
	// 1 获取帖子的发布时间
	postTime := rdb.ZScore(ctx, "KeyPostTimeZSet", postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2更新帖子的分数
	//先获取用户给这个帖子的投票纪录 看是-1 0 还是1
	voteValue := rdb.ZScore(ctx, KeyPostVoted+postId, userid).Val()
	var op float64
	if value > voteValue {
		op = 1
	} else {
		op = -1
	}
	// 计算两次投票的差值
	diff := math.Abs(voteValue - value)
	_, err = rdb.ZIncrBy(ctx, KeyPostScoreZSet, op*diff*scorePerVote, postId).Result()
	if err != nil {
		return err
	}
	//3 记录用户为该帖子投票的数据
	if value == 0 {
		rdb.ZRem(ctx, KeyPostVoted+postId, postId)
	}
	_, err = rdb.ZAdd(ctx, KeyPostVoted+postId, &redis.Z{
		Score:  value,
		Member: userid,
	}).Result()
	return err
}
