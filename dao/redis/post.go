package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 423 //每一票的分数
)

/*
投票的几种情况
	direction = 1 时,有两种情况：
			1.之前没有投过票，现在投赞成票   差值的绝对值为 1  1-0  +432
			2.之前投反对票 现在改投赞成票    差值的绝对值为2  1-(-1)  +432*2
	direction = 0 时,有两种情况：
			1.之前投过赞成票，现在要取消   差值的绝对值为1   |0-1| -432
			2.之前投反对票 现在要取消     差值的绝对值为1	 |0-(1)| +432
	direction = -1 时,有两种情况：
			1.之前没投过票，现在投反对票   差值的绝对值1	|-1-0| -432
			2.之前投过赞成票 现在改投反对票  差值的绝对值2  |-1-1|  -432*2
*/
var (
	ctx               = context.Background()
	ErrVoteTimeExpire = errors.New("超出投票时间")
)

// CreatePostTime 这里面的两个操作都必须同时成功 要么都失败
func CreatePostTime(postId int64) error {
	pipeline := rdb.Pipeline()
	// 帖子时间
	rdb.ZAdd(ctx, KeyPostTimeZSet, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	// 帖子分数
	rdb.ZAdd(ctx, KeyPostScoreZSet, &redis.Z{
		Score:  0,
		Member: postId,
	})
	_, err := pipeline.Exec(ctx)
	return err
}

func VoteForPost(userid, postId string, value float64) (err error) {
	// 1 获取帖子的发布时间
	postTime := rdb.ZScore(ctx, "KeyPostTimeZSet", postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2更新帖子的分数
	//先获取用户给这个帖子的投票纪录 看是-1 0 还是1
	voteValue := rdb.ZScore(ctx, KeyPostVoted+postId, userid).Val()
	fmt.Println(voteValue)
	var op float64
	if value > voteValue {
		op = 1
	} else {
		op = -1
	}
	// 计算两次投票的差值
	diff := math.Abs(voteValue - value)
	pipeline := rdb.TxPipeline()
	rdb.ZIncrBy(ctx, KeyPostScoreZSet, op*diff*scorePerVote, postId)
	//3 记录用户为该帖子投票的数据
	if value == 0 {
		rdb.ZRem(ctx, KeyPostVoted+postId, postId)
	}
	rdb.ZAdd(ctx, KeyPostVoted+postId, &redis.Z{
		Score:  value,
		Member: userid,
	})
	_, err = pipeline.Exec(ctx)
	return err
}
