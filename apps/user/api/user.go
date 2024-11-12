package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/user"
	"github.com/emicklei/go-restful/v3"
)

func (h *userHandler) create_user(r *restful.Request, w *restful.Response) {
	//0.get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	h.log.Info().Msgf("CREATE USER , context_user: %s", context_user.Spec.Username)
	// result = roles.NewRoleUserBinding()
	//1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, "User Admin", "user.create")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	_, err := h.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}
	//2.read the request body parametars
	req := user.NewCreateUserRequest()
	// req.Username = "kade6"
	// req.Password = "123456"
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, exception.NewInternalServerError("read request(create_user): %s; error: %v", req, err))
		return
	}
	//3.create user
	_, err = h.user.CreateUser(r.Request.Context(), req)
	h.log.Error().Msgf("err: %s", req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("create user error:%v", err))
		return
	}
	response.Success(w, "create user success")
}

func (h *userHandler) list_user(r *restful.Request, w *restful.Response) {
	//0.get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	h.log.Info().Msgf("LIST USER , context_user: %s", context_user.Spec.Username)

	//1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, "User Admin", "user.list")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	_, err := h.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}
	//2.read the request body parametars
	req := NewQueryUserRequest()
	// var myType user.TYPE = user.TYPE_SUB
	// req.Type = &myType
	pageSizeStr := r.Request.URL.Query().Get("page_size")
	pageNumberStr := r.Request.URL.Query().Get("page_number")

	if pageSizeStr != "" {
		ps, _ := strconv.ParseUint(pageSizeStr, 10, 64)
		req.Page.PageSize = ps
	}

	if pageNumberStr != "" {
		pn, _ := strconv.ParseUint(pageNumberStr, 10, 64)
		req.Page.PageNumber = pn
	}
	//3.list users
	if req.Page.PageSize == 0 {
		req.Page.PageSize = 10
	}
	h.log.Info().Msgf("LIST USER , req: %s", req)
	l, err := h.user.ListUser(context.Background(), req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("read request(list_user): %s; error: %v", req, err))
		return
	}

	response.Success(w, l)
}

func (h *userHandler) delete_user(r *restful.Request, w *restful.Response) {
	//0.get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	h.log.Info().Msgf("LIST USER , context_user: %s", context_user.Spec.Username)

	//1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, "User Admin", "user.delete")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	_, err := h.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}
	//2.read the request body parametars
	req := user.NewDeleteUserRequest()
	idsParam := r.QueryParameters("user_ids")
	// 将 ID 字符串按逗号分隔，并转换为切片
	// 将查询参数按逗号分隔
	if len(idsParam) > 0 {
		idsList := strings.Split(idsParam[0], ",")
		// 去除每个 ID 的前后空白
		for i := range idsList {
			idsList[i] = strings.TrimSpace(idsList[i])
		}

		req.UserIds = idsList
	}
	// req.UserIds = append(req.UserIds, r.QueryParameters("id")...)
	// pageSizeStr := r.Request.URL.Query().Get("page_size")
	// pageNumberStr := r.Request.URL.Query().Get("page_number")
	//3.list users
	// 打印 ID 列表
	if len(req.UserIds) == 0 {
		response.Failed(w, exception.NewInternalServerError("error: %v", "user_ids is empty"))
		return
	}
	h.log.Info().Msgf("LIST USER , req: %v ", req)
	l, err := h.user.DeleteUser(context.Background(), req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("The user is %s; error: %v", req, err))
		return
	}

	response.Success(w, l)
}
