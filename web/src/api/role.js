import request from '@/utils/request.js'


// 获取所有的user数据
export const roleListService = () => {
    return request.get("/role")
}

/*添加角色函数
RoleModel {
    roleName: '',\
    accesses: [],
}

*/

export const addRoleService = (addRoleData) => {
    // const params = new URLSearchParams();
    // console.log(addRoleData.roleName)
    // params.append("roleName", addRoleData.RoleName)
    // params.append("accesses", addRoleData.accesses)
    // console.log(params)
    const requestData = {
        "roleName": addRoleData.roleName,
        "accesses": addRoleData.accesses
    }
    console.log(requestData)
    return request.post("/role", requestData)
}