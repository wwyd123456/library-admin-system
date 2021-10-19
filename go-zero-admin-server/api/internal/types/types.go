// Code generated by goctl. DO NOT EDIT.
package types

type Reply struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUsersReq struct {
	Page     int    `form:"page,default=1"`
	Limit    int    `form:"limit,default=10"`
	Username string `form:"username,optional"`
}

type AddUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar,optional"`
	Info     string `json:"info,optional"`
}

type UpdateUserReq struct {
	Id       uint   `path:"id"`
	Password string `json:"password"`
	Avatar   string `json:"avatar,optional"`
	Info     string `json:"info,optional"`
}

type EditUserRolesReq struct {
	Id      uint   `path:"id"`
	RoleIds []uint `json:"roleIds,optional"`
}

type DeleteUserReq struct {
	Id uint `path:"id"`
}

type UpdateUserByJwtReq struct {
	Avatar string `json:"avatar,optional"`
	Info   string `json:"info,optional"`
}

type ChangePasswordByJwtReq struct {
	OldPassword      string `json:"oldPassword"`
	NewPassword      string `json:"newPassword"`
	NewPasswordAgain string `json:"newPasswordAgain"`
}

type IdReq struct {
	Id uint `path:"id"`
}

type UsernameReq struct {
	Username string `path:"username"`
}

type NameReq struct {
	Name string `path:"name"`
}

type GetBooksReq struct {
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Name   string `form:"name,optional"`
	Author string `form:"author,optional"`
}

type AddBookReq struct {
	Name   string `json:"name"`
	Image  string `json:"image,optional"`
	Author string `json:"author,optional"`
	Info   string `json:"info,optional"`
	TypeId uint   `json:"typeId"`
}

type UpdateBookReq struct {
	Id     uint   `path:"id"`
	Name   string `json:"name"`
	Image  string `json:"image,optional"`
	Author string `json:"author,optional"`
	Info   string `json:"info,optional"`
	TypeId uint   `json:"typeId"`
}

type DeleteBookReq struct {
	Id uint `path:"id"`
}

type GetTypesReq struct {
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
	Name  string `form:"name,optional"`
}

type AddTypeReq struct {
	Name  string `json:"name"`
	Intro string `json:"intro,optional"`
}

type UpdateTypeReq struct {
	Id    uint   `path:"id"`
	Name  string `json:"name"`
	Intro string `json:"intro,optional"`
}

type DeleteTypeReq struct {
	Id uint `path:"id"`
}

type GetVipsReq struct {
	Page       int    `form:"page,default=1"`
	Limit      int    `form:"limit,default=10"`
	CardNumber string `form:"cardNumber,optional"`
	Name       string `form:"name,optional"`
}

type AddVipReq struct {
	Name string `json:"name,optional"`
	Info string `json:"info,optional"`
}

type UpdateVipReq struct {
	Id   uint   `path:"id"`
	Name string `json:"name,optional"`
	Info string `json:"info,optional"`
}

type DeleteVipReq struct {
	Id uint `path:"id"`
}
