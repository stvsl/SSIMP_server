package redis

import (
	"context"
	"strconv"
	"time"

	"stvsljl.com/SSIMP/utils"
)

type RunningData struct {
	NewViewCount int
	// 七天网站访问量
	SevenDaysViewCount []int
	// 七天文章访问量
	SevenDaysArticleViewCount []int
}

func GetRunningDataStruct() *RunningData {
	return &RunningData{
		SevenDaysViewCount:        make([]int, 7),
		SevenDaysArticleViewCount: make([]int, 7),
	}
}

func (a *RunningData) ReadAllFromRedis() (RunningData, error) {
	// 逐个读取
	newViewCount, err := rdb.HGet(context.Background(), "NewViewCount", time.Now().Format("2006-01-02")).Result()
	if err != nil {
		utils.Log.Error("键值不存在", err)
		newViewCount = "0"
	}
	a.NewViewCount, _ = strconv.Atoi(newViewCount)
	for i := 0; i < 7; i++ {
		viewCount, err := rdb.HGet(context.Background(), "SevenDaysViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02")).Result()
		if err != nil {
			utils.Log.Error("键值不存在", err)
			a.SevenDaysViewCount[i] = 0
		}
		a.SevenDaysViewCount[i], _ = strconv.Atoi(viewCount)
	}
	for i := 0; i < 7; i++ {
		articleViewCount, err := rdb.HGet(context.Background(), "SevenDaysArtilceViewCount", time.Now().AddDate(0, 0, -i).Format("2006-01-02")).Result()
		if err != nil {
			utils.Log.Error("键值不存在", err)
			a.SevenDaysArticleViewCount[i] = 0
		}
		a.SevenDaysArticleViewCount[i], _ = strconv.Atoi(articleViewCount)
	}
	return *a, err
}

func (a *RunningData) AddNewViewCount() {
	// 更新当日新增浏览量
	newViewCount, err := rdb.HGet(context.Background(), "NewViewCount", time.Now().Format("2006-01-02")).Result()
	if err != nil {
		newViewCount = "0"
		syncYesterdayArticleViewCount()
	}
	// 设置过期时间
	rdb.Expire(context.Background(), "NewViewCount", time.Hour*24)
	// 数据写回
	a.NewViewCount, _ = strconv.Atoi(newViewCount)
	a.NewViewCount++
	_, err = rdb.HSet(context.Background(), "NewViewCount", time.Now().Format("2006-01-02"), a.NewViewCount).Result()
	if err != nil {
		utils.Log.Error("redis写入失败", err)
	}
}

func syncYesterdayArticleViewCount() {
	// 同步昨日浏览量
	yesterdayViewCount, err := rdb.HGet(context.Background(), "NewViewCount", time.Now().AddDate(0, 0, -1).Format("2006-01-02")).Result()
	if err != nil {
		utils.Log.Error("键值不存在", err)
		yesterdayViewCount = "0"
	}
	// 数据写回
	_, err = rdb.HSet(context.Background(), "SevenDaysArtilceViewCount", time.Now().Format("2006-01-02"), yesterdayViewCount).Result()
	if err != nil {
		utils.Log.Error("redis写入失败", err)
	}
}

func SetUpdateWebViewCount() {
	// 更新今日网站访问量
	viewCount, err := rdb.HGet(context.Background(), "SevenDaysViewCount", time.Now().Format("2006-01-02")).Result()
	if err != nil {
		viewCount = "0"
	}
	viewCountInt, _ := strconv.Atoi(viewCount)
	viewCountInt++
	// 数据写回
	_, err = rdb.HSet(context.Background(), "SevenDaysViewCount", time.Now().Format("2006-01-02"), viewCountInt).Result()
	if err != nil {
		utils.Log.Error("redis写入失败", err)
	}
	// 设置过期时间
	rdb.Expire(context.Background(), "SevenDaysViewCount", time.Hour*24*7)
}
