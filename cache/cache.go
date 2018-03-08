package cache

import "github.com/konojunya/HEW2018/model"

var db = model.NewDBConn()

// RefreshAll cacheをリフレッシュする
func RefreshAll() {
	Product.Reload()
}
