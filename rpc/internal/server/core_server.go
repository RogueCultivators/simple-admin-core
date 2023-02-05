// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package server

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/api"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/authority"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/base"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/department"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/menu"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/oauth"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/post"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/role"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/token"
	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

type CoreServer struct {
	svcCtx *svc.ServiceContext
	core.UnimplementedCoreServer
}

func NewCoreServer(svcCtx *svc.ServiceContext) *CoreServer {
	return &CoreServer{
		svcCtx: svcCtx,
	}
}

func (s *CoreServer) CreateOrUpdateApi(ctx context.Context, in *core.ApiInfo) (*core.BaseResp, error) {
	l := api.NewCreateOrUpdateApiLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateApi(in)
}

func (s *CoreServer) DeleteApi(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := api.NewDeleteApiLogic(ctx, s.svcCtx)
	return l.DeleteApi(in)
}

func (s *CoreServer) GetApiList(ctx context.Context, in *core.ApiListReq) (*core.ApiListResp, error) {
	l := api.NewGetApiListLogic(ctx, s.svcCtx)
	return l.GetApiList(in)
}

func (s *CoreServer) GetMenuAuthority(ctx context.Context, in *core.IDReq) (*core.RoleMenuAuthorityResp, error) {
	l := authority.NewGetMenuAuthorityLogic(ctx, s.svcCtx)
	return l.GetMenuAuthority(in)
}

func (s *CoreServer) CreateOrUpdateMenuAuthority(ctx context.Context, in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	l := authority.NewCreateOrUpdateMenuAuthorityLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateMenuAuthority(in)
}

func (s *CoreServer) InitDatabase(ctx context.Context, in *core.Empty) (*core.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// Department management
func (s *CoreServer) CreateOrUpdateDepartment(ctx context.Context, in *core.DepartmentInfo) (*core.BaseResp, error) {
	l := department.NewCreateOrUpdateDepartmentLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateDepartment(in)
}

func (s *CoreServer) GetDepartmentList(ctx context.Context, in *core.DepartmentListReq) (*core.DepartmentListResp, error) {
	l := department.NewGetDepartmentListLogic(ctx, s.svcCtx)
	return l.GetDepartmentList(in)
}

func (s *CoreServer) DeleteDepartment(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := department.NewDeleteDepartmentLogic(ctx, s.svcCtx)
	return l.DeleteDepartment(in)
}

func (s *CoreServer) BatchDeleteDepartment(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := department.NewBatchDeleteDepartmentLogic(ctx, s.svcCtx)
	return l.BatchDeleteDepartment(in)
}

func (s *CoreServer) UpdateDepartmentStatus(ctx context.Context, in *core.StatusCodeReq) (*core.BaseResp, error) {
	l := department.NewUpdateDepartmentStatusLogic(ctx, s.svcCtx)
	return l.UpdateDepartmentStatus(in)
}

func (s *CoreServer) CreateOrUpdateDictionary(ctx context.Context, in *core.DictionaryInfo) (*core.BaseResp, error) {
	l := dictionary.NewCreateOrUpdateDictionaryLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateDictionary(in)
}

func (s *CoreServer) DeleteDictionary(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := dictionary.NewDeleteDictionaryLogic(ctx, s.svcCtx)
	return l.DeleteDictionary(in)
}

func (s *CoreServer) GetDictionaryList(ctx context.Context, in *core.DictionaryListReq) (*core.DictionaryList, error) {
	l := dictionary.NewGetDictionaryListLogic(ctx, s.svcCtx)
	return l.GetDictionaryList(in)
}

func (s *CoreServer) GetDetailByDictionaryName(ctx context.Context, in *core.DictionaryDetailReq) (*core.DictionaryDetailList, error) {
	l := dictionary.NewGetDetailByDictionaryNameLogic(ctx, s.svcCtx)
	return l.GetDetailByDictionaryName(in)
}

func (s *CoreServer) CreateOrUpdateDictionaryDetail(ctx context.Context, in *core.DictionaryDetail) (*core.BaseResp, error) {
	l := dictionary.NewCreateOrUpdateDictionaryDetailLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateDictionaryDetail(in)
}

func (s *CoreServer) DeleteDictionaryDetail(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := dictionary.NewDeleteDictionaryDetailLogic(ctx, s.svcCtx)
	return l.DeleteDictionaryDetail(in)
}

func (s *CoreServer) CreateOrUpdateMenu(ctx context.Context, in *core.CreateOrUpdateMenuReq) (*core.BaseResp, error) {
	l := menu.NewCreateOrUpdateMenuLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateMenu(in)
}

func (s *CoreServer) DeleteMenu(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := menu.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

func (s *CoreServer) GetMenuListByRole(ctx context.Context, in *core.IDReq) (*core.MenuInfoList, error) {
	l := menu.NewGetMenuListByRoleLogic(ctx, s.svcCtx)
	return l.GetMenuListByRole(in)
}

func (s *CoreServer) GetMenuList(ctx context.Context, in *core.PageInfoReq) (*core.MenuInfoList, error) {
	l := menu.NewGetMenuListLogic(ctx, s.svcCtx)
	return l.GetMenuList(in)
}

func (s *CoreServer) CreateOrUpdateMenuParam(ctx context.Context, in *core.CreateOrUpdateMenuParamReq) (*core.BaseResp, error) {
	l := menu.NewCreateOrUpdateMenuParamLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateMenuParam(in)
}

func (s *CoreServer) DeleteMenuParam(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := menu.NewDeleteMenuParamLogic(ctx, s.svcCtx)
	return l.DeleteMenuParam(in)
}

func (s *CoreServer) GetMenuParamListByMenuId(ctx context.Context, in *core.IDReq) (*core.MenuParamListResp, error) {
	l := menu.NewGetMenuParamListByMenuIdLogic(ctx, s.svcCtx)
	return l.GetMenuParamListByMenuId(in)
}

func (s *CoreServer) CreateOrUpdateProvider(ctx context.Context, in *core.ProviderInfo) (*core.BaseResp, error) {
	l := oauth.NewCreateOrUpdateProviderLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateProvider(in)
}

func (s *CoreServer) DeleteProvider(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := oauth.NewDeleteProviderLogic(ctx, s.svcCtx)
	return l.DeleteProvider(in)
}

func (s *CoreServer) GetProviderList(ctx context.Context, in *core.PageInfoReq) (*core.ProviderListResp, error) {
	l := oauth.NewGetProviderListLogic(ctx, s.svcCtx)
	return l.GetProviderList(in)
}

func (s *CoreServer) OauthLogin(ctx context.Context, in *core.OauthLoginReq) (*core.OauthRedirectResp, error) {
	l := oauth.NewOauthLoginLogic(ctx, s.svcCtx)
	return l.OauthLogin(in)
}

func (s *CoreServer) OauthCallback(ctx context.Context, in *core.CallbackReq) (*core.LoginResp, error) {
	l := oauth.NewOauthCallbackLogic(ctx, s.svcCtx)
	return l.OauthCallback(in)
}

// Post management
func (s *CoreServer) CreateOrUpdatePost(ctx context.Context, in *core.PostInfo) (*core.BaseResp, error) {
	l := post.NewCreateOrUpdatePostLogic(ctx, s.svcCtx)
	return l.CreateOrUpdatePost(in)
}

func (s *CoreServer) GetPostList(ctx context.Context, in *core.PostListReq) (*core.PostListResp, error) {
	l := post.NewGetPostListLogic(ctx, s.svcCtx)
	return l.GetPostList(in)
}

func (s *CoreServer) DeletePost(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := post.NewDeletePostLogic(ctx, s.svcCtx)
	return l.DeletePost(in)
}

func (s *CoreServer) BatchDeletePost(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := post.NewBatchDeletePostLogic(ctx, s.svcCtx)
	return l.BatchDeletePost(in)
}

func (s *CoreServer) UpdatePostStatus(ctx context.Context, in *core.StatusCodeReq) (*core.BaseResp, error) {
	l := post.NewUpdatePostStatusLogic(ctx, s.svcCtx)
	return l.UpdatePostStatus(in)
}

func (s *CoreServer) CreateOrUpdateRole(ctx context.Context, in *core.RoleInfo) (*core.BaseResp, error) {
	l := role.NewCreateOrUpdateRoleLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateRole(in)
}

func (s *CoreServer) DeleteRole(ctx context.Context, in *core.IDReq) (*core.BaseResp, error) {
	l := role.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

func (s *CoreServer) GetRoleById(ctx context.Context, in *core.IDReq) (*core.RoleInfo, error) {
	l := role.NewGetRoleByIdLogic(ctx, s.svcCtx)
	return l.GetRoleById(in)
}

func (s *CoreServer) GetRoleList(ctx context.Context, in *core.PageInfoReq) (*core.RoleListResp, error) {
	l := role.NewGetRoleListLogic(ctx, s.svcCtx)
	return l.GetRoleList(in)
}

func (s *CoreServer) UpdateRoleStatus(ctx context.Context, in *core.StatusCodeReq) (*core.BaseResp, error) {
	l := role.NewUpdateRoleStatusLogic(ctx, s.svcCtx)
	return l.UpdateRoleStatus(in)
}

func (s *CoreServer) CreateOrUpdateToken(ctx context.Context, in *core.TokenInfo) (*core.BaseResp, error) {
	l := token.NewCreateOrUpdateTokenLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateToken(in)
}

func (s *CoreServer) DeleteToken(ctx context.Context, in *core.UUIDReq) (*core.BaseResp, error) {
	l := token.NewDeleteTokenLogic(ctx, s.svcCtx)
	return l.DeleteToken(in)
}

func (s *CoreServer) BatchDeleteToken(ctx context.Context, in *core.UUIDsReq) (*core.BaseResp, error) {
	l := token.NewBatchDeleteTokenLogic(ctx, s.svcCtx)
	return l.BatchDeleteToken(in)
}

func (s *CoreServer) GetTokenList(ctx context.Context, in *core.TokenListReq) (*core.TokenListResp, error) {
	l := token.NewGetTokenListLogic(ctx, s.svcCtx)
	return l.GetTokenList(in)
}

func (s *CoreServer) UpdateTokenStatus(ctx context.Context, in *core.StatusCodeUUIDReq) (*core.BaseResp, error) {
	l := token.NewUpdateTokenStatusLogic(ctx, s.svcCtx)
	return l.UpdateTokenStatus(in)
}

func (s *CoreServer) BlockUserAllToken(ctx context.Context, in *core.UUIDReq) (*core.BaseResp, error) {
	l := token.NewBlockUserAllTokenLogic(ctx, s.svcCtx)
	return l.BlockUserAllToken(in)
}

func (s *CoreServer) Login(ctx context.Context, in *core.LoginReq) (*core.LoginResp, error) {
	l := user.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *CoreServer) ChangePassword(ctx context.Context, in *core.ChangePasswordReq) (*core.BaseResp, error) {
	l := user.NewChangePasswordLogic(ctx, s.svcCtx)
	return l.ChangePassword(in)
}

func (s *CoreServer) CreateOrUpdateUser(ctx context.Context, in *core.CreateOrUpdateUserReq) (*core.BaseResp, error) {
	l := user.NewCreateOrUpdateUserLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateUser(in)
}

func (s *CoreServer) GetUserById(ctx context.Context, in *core.UUIDReq) (*core.UserInfoResp, error) {
	l := user.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *CoreServer) GetUserList(ctx context.Context, in *core.GetUserListReq) (*core.UserListResp, error) {
	l := user.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}

func (s *CoreServer) DeleteUser(ctx context.Context, in *core.UUIDReq) (*core.BaseResp, error) {
	l := user.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *CoreServer) BatchDeleteUser(ctx context.Context, in *core.UUIDsReq) (*core.BaseResp, error) {
	l := user.NewBatchDeleteUserLogic(ctx, s.svcCtx)
	return l.BatchDeleteUser(in)
}

func (s *CoreServer) UpdateProfile(ctx context.Context, in *core.UpdateProfileReq) (*core.BaseResp, error) {
	l := user.NewUpdateProfileLogic(ctx, s.svcCtx)
	return l.UpdateProfile(in)
}

func (s *CoreServer) UpdateUserStatus(ctx context.Context, in *core.StatusCodeUUIDReq) (*core.BaseResp, error) {
	l := user.NewUpdateUserStatusLogic(ctx, s.svcCtx)
	return l.UpdateUserStatus(in)
}
