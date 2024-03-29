//导入request.js请求工具
import request from '@/utils/request.js'
//提供调用注册接口的函数
export const userRegisterService = (registerData)=>{
    return request.post('/register',registerData);
}
//提供调用登录接口的函数
export const userLoginService = (data)=>{
    return request.post('/login', data);
}
//获取用户详细信息
export const userInfoService = ()=>{
    return request.get('/user/userInfo')
}
//修改个人信息
export const userInfoUpdateService = (userInfoData)=>{
   return request.put('/user/update',userInfoData)
}

//修改头像
export const userAvatarUpdateService = (avatarUrl)=>{
    const params = new URLSearchParams();
    params.append('avatarUrl',avatarUrl)
    return request.patch('/user/updateAvatar',params)
}

// 获取所有的user数据
export const userListService = (requestData) => {
    return request.get("/user", {params : requestData})
}

export const getUserByOptionService = (Data) => {
    return request.get("/user/option",Data)
}

export const deleteUserService = (id) => {
    return request.delete("/user?id="+id)
}

export const addUserService = (data) => {
    const requestData = {
        "userName": data.userName,
        "email": data.email,
        "roleID": data.role.id,
        "role": data.role,
    }
}

export const updateUserService = (id, roleid) => {
    return request.put("/user?id=" + id + "&roleid=" + roleid);
}