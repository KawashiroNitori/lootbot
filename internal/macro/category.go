//go:generate go run github.com/dmarkham/enumer -type Category -trimprefix Category -transform lower -sql -json -output category_string.go
package macro

type Category int

const (
	_                     Category = iota
	CategoryWeapon                 // 武器
	CategoryCoffer                 // 装备箱
	CategoryCoating                // 首饰药
	CategoryTomestone              // 神典石
	CategoryRoborant               // 武器药
	CategorySpool                  // 纤维
	CategoryMount                  // 坐骑
	CategoryOrchestraRoll          // 乐谱
	CategoryCompanion              // 宠物
)

var _categoryNameMap = map[Category]string{
	CategoryWeapon:        "武器",
	CategoryCoffer:        "装备箱",
	CategoryCoating:       "首饰药",
	CategoryTomestone:     "神典石",
	CategoryRoborant:      "武器药",
	CategorySpool:         "纤维",
	CategoryMount:         "坐骑",
	CategoryOrchestraRoll: "乐谱",
	CategoryCompanion:     "宠物",
}

func (i Category) Name() string {
	return _categoryNameMap[i]
}

func (i Category) Values() []string {
	return CategoryStrings()
}
