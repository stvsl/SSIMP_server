package redis

import (
	"context"
	"strconv"
	"time"

	"stvsljl.com/SSIMP/utils"
)

type runningData struct {
	NewViewCount int
	// 七天网站访问量
	SevenDaysViewCount []int
	// 七天文章访问量
	SevenDaysArticleViewCount []int
}

func (a *runningData) WriteToRedis() error {
	// 逐个写入
	status, err := rdb.HSet(context.Background(), "NewViewCount", time.Now().Format("2006-01-02"), a.NewViewCount).Result()
	if err != nil || status == 0 {
		utils.Log.Error("redis写入失败", err)
		return err
	}
	for i, v := range a.SevenDaysViewCount {
		status, err := rdb.HSet(context.Background(), "SevenDaysViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02"), v).Result()
		if err != nil || status == 0 {
			utils.Log.Error("redis写入失败", err)
			return err
		}
	}
	for i, v := range a.SevenDaysArticleViewCount {
		status, err := rdb.HSet(context.Background(), "SevenDaysArticleViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02"), v).Result()
		if err != nil || status == 0 {
			utils.Log.Error("redis写入失败", err)
			return err
		}
	}
	return nil
}

func (a *runningData) ReadAndRemoveFromRedis() (runningData, error) {
	// 逐个读取
	newViewCount, err := rdb.HGet(context.Background(), "NewViewCount", time.Now().Format("2006-01-02")).Result()
	if err != nil {
		utils.Log.Error("键值不存在", err)
		return runningData{}, err
	}
	a.NewViewCount, _ = strconv.Atoi(newViewCount)
	for i := 0; i < 7; i++ {
		viewCount, err := rdb.HGet(context.Background(), "SevenDaysViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02")).Result()
		if err != nil {
			utils.Log.Error("键值不存在", err)
			return runningData{}, err
		}
		a.SevenDaysViewCount[i], _ = strconv.Atoi(viewCount)
	}
	for i := 0; i < 7; i++ {
		articleViewCount, err := rdb.HGet(context.Background(), "SevenDaysArticleViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02")).Result()
		if err != nil {
			utils.Log.Error("键值不存在", err)
			return runningData{}, err
		}
		a.SevenDaysArticleViewCount[i], _ = strconv.Atoi(articleViewCount)
	}
	// 逐个删除
	status, err := rdb.HDel(context.Background(), "NewViewCount", time.Now().Format("2006-01-02")).Result()
	if err != nil || status == 0 {
		utils.Log.Error("redis删除失败：", err)
		return runningData{}, err
	}
	for i := 0; i < 7; i++ {
		status, err := rdb.HDel(context.Background(), "SevenDaysViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02")).Result()
		if err != nil || status == 0 {
			utils.Log.Error("redis删除失败：", err)
			return runningData{}, err
		}
	}
	for i := 0; i < 7; i++ {
		status, err := rdb.HDel(context.Background(), "SevenDaysArticleViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02")).Result()
		if err != nil || status == 0 {
			utils.Log.Error("redis删除失败：", err)
			return runningData{}, err
		}
	}
	return *a, err
}
