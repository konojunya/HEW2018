package main

import (
	"fmt"

	"github.com/konojunya/HEW2018/model"
	"github.com/konojunya/HEW2018/service"
)

func main() {
	create(model.Product{
		Title:  "HEW投票システム",
		Author: "河野 純也",
		Votes:  0,
	})
	create(model.Product{
		Title:     "Review Lot",
		Author:    "石田 一馬",
		Votes:     0,
		Thumbnail: "/image/01.png",
	})
	create(model.Product{
		Title:     "Wordy",
		Author:    "大江 和摩",
		Votes:     0,
		Thumbnail: "/image/02.png",
	})
	create(model.Product{
		Title:     "Funny Diary",
		Author:    "大塚 翔太",
		Votes:     0,
		Thumbnail: "/image/03.png",
	})
	create(model.Product{
		Title:     "クラブのサポート・クラボ",
		Author:    "尾崎 由芽",
		Votes:     0,
		Thumbnail: "/image/05.png",
	})
	create(model.Product{
		Title:     "S Store",
		Author:    "川相 拓哉",
		Votes:     0,
		Thumbnail: "/image/06.png",
	})
	create(model.Product{
		Title:     "BOOK SEEKER",
		Author:    "小泉 優馬",
		Votes:     0,
		Thumbnail: "/image/07.png",
	})
	create(model.Product{
		Title:     "photori",
		Author:    "鴻上 萌",
		Votes:     0,
		Thumbnail: "/image/08.png",
	})
	create(model.Product{
		Title:     "グラブル募集掲示板",
		Author:    "菅 優祟",
		Votes:     0,
		Thumbnail: "/image/10.png",
	})
	create(model.Product{
		Title:     "岩瀬唯奈SITE（偽）",
		Author:    "高鉾 大貴",
		Votes:     0,
		Thumbnail: "/image/11.png",
	})
	create(model.Product{
		Title:     "単語帳アプリ コネクト",
		Author:    "竹林 寛晃",
		Votes:     0,
		Thumbnail: "/image/12.jpg",
	})
	create(model.Product{
		Title:     "L'Atelier",
		Author:    "田村 勇貴",
		Votes:     0,
		Thumbnail: "/image/13.png",
	})
	create(model.Product{
		Title:     "VietFood",
		Author:    "トラン ヴ ホアン",
		Votes:     0,
		Thumbnail: "/image/14.png",
	})
	create(model.Product{
		Title:     "レシピを保存するwebapp",
		Author:    "中久保 彰伽",
		Votes:     0,
		Thumbnail: "/image/15.png",
	})
	create(model.Product{
		Title:     "Dinner Match",
		Author:    "西川 和希",
		Votes:     0,
		Thumbnail: "/image/16.png",
	})
	create(model.Product{
		Title:     "グッズ作り共有支援サービス",
		Author:    "東村 美るり",
		Votes:     0,
		Thumbnail: "/image/17.png",
	})
	create(model.Product{
		Title:     "Skype Map",
		Author:    "森 真樹史",
		Votes:     0,
		Thumbnail: "/image/18.png",
	})
	create(model.Product{
		Title:     "Study Node",
		Author:    "綿野 拓也",
		Votes:     0,
		Thumbnail: "/image/19.jpg",
	})
}

func create(p model.Product) {
	product, err := service.Create(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("product created: %v\n", product.CreatedAt)
}
