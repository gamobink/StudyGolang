package Permission

import (
	"../../dao"
	"../../global"
	"../../models"
	"net/http"
)

//添加博客的权限检查
func PublishBlog(blog *models.Blog,w http.ResponseWriter, r *http.Request) (uid int,err error) {
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp < 1 {
		err = global.NoPermission
		return
	}
	return
}

//修改博客的权限检查
func UpDateBlog(blog *models.Blog, w http.ResponseWriter, r *http.Request) (uid int,err error) {
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	if uid != blog.PublisherId {
		err = global.NoPermission
		return
	}
	return
}

//删除博客的权限检查
func DeleteBlog(blog *models.Blog, w http.ResponseWriter, r *http.Request) (uid int,err error) {
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp == 5 {
		return
	}
	if uid != blog.PublisherId {
		err = global.NoPermission
		return
	}
	return
}
