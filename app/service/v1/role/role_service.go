package roleService

import (
    model "gin-test/app/models"
    "log"
)

type RoleStruct struct {
    PageNum  int
    PageSize int
}

type NewRoleStruct struct {
    RoleKey string `json:"role_key" form:"role_key" validate:"required,min=4,max=10"`
}

type UpdateRoleStruct struct {
     // 角色名称
    RoleName string `json:"role_name" form:"role_name" validate:"required,min=4,max=10" minLength:"4" maxLength:"10"`
    // 备注
    Remark string `json:"remark" form:"remark" validate:"required,min=4,max=100" minLength:"4" maxLength:"100"`
}

func UpdateRole(roleId int, u UpdateRoleStruct) bool{
    wheres := make(map[string]interface{})
    wheres["role_id"] = roleId

    updates := make(map[string]interface{})
    updates["role_name"] = u.RoleName
    updates["remark"] = u.Remark
    _, rowsAffected := model.Update(&model.Role{}, wheres, updates)
    if rowsAffected == 0 {
        log.Println("修改Role失败！")
        return false
    }
    return true
}

func (u *RoleStruct) getConditionMaps() map[string]interface{} {
    maps := make(map[string]interface{})
    //maps["deleted_at"] = nil
    return maps
}

func (u *RoleStruct) Count() (int, error) {
    return model.GetTotal(model.Role{}, u.getConditionMaps())
}

func (u *RoleStruct) GetAll() ([]*model.Role, error) {
    Roles, err := model.GetRoles(u.PageNum, u.PageSize, u.getConditionMaps())
    if err != nil {
        return nil, err
    }

    return Roles, nil
}