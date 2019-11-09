package vm

// BaseViewModel struct
type BaseViewModel struct {
	Title string
}

// SetTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}


//// 由于_base.html 基础模板中有 Title 字段，所以 Title是每个view都必有的字段，
//// 我们将它单独设成个 BaseViewStruct，方便用 匿名组合
//// Deprecated IndexViewModel struct
//type BaseViewModel struct {
//	Title string
//	User
//	Posts []Post
//}