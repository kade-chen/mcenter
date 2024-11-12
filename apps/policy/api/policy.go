package api

import (
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/user"
	"github.com/emicklei/go-restful/v3"
)

func (h *policyHandler) check_permission(r *restful.Request, w *restful.Response) {
	//0.get context user
	context_user := r.Request.Context().Value("user").(*user.User)
	h.log.Info().Msgf("CHECK PERMISSION USER , context_user: %s", context_user.Spec.Username)

	role := r.Request.URL.Query().Get("role")
	permission := r.Request.URL.Query().Get("permission")
	h.log.Info().Msgf("CHECK PERMISSION USER , role: %s , permission: %s", role, permission)

	//1.check permission
	user_pms_req := policy.NewUser_Permission_Request(context_user.Spec.Username, role, permission)
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	message, err := h.policy.Check_Permission(r.Request.Context(), user_pms_req)
	if err != nil {
		response.Failed(w, exception.NewInternalServerError("error: %v", err))
		return
	}

	response.Success(w, message)
}
