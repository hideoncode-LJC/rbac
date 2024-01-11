import { createRouter, createWebHistory } from 'vue-router'

//导入组件
import LoginVue from '@/views/Login.vue'
import LayoutVue from '@/views/Layout.vue'
import AccessManageVue from '@/views/article/AccessManage.vue'
import UserManageVue from '@/views/article/UserManage.vue'
import RoleManageVue from '@/views/article/RoleManage.vue'
import UserAvatarVue from '@/views/user/UserAvatar.vue'
import UserInfoVue from '@/views/user/UserInfo.vue'
import UserResetPasswordVue from '@/views/user/UserResetPassword.vue'
import AccessVue from '@/views/access/Access.vue'
import IndexVue from '@/views/Index.vue'

//定义路由关系
const routes = [
    { path: '/login', component: LoginVue },
    {
        path: '/', 
        component: LayoutVue,
        redirect:'/index', 
        children: [
            { path: '/index', component: IndexVue, props: true },
            { path: '/article/user', component: UserManageVue },
            { path: '/article/role', component: RoleManageVue },
            { path: '/article/access', component: AccessManageVue },
            { path: '/user/info', component: UserInfoVue },
            { path: '/user/avatar', component: UserAvatarVue },
            { path: '/user/resetPassword', component: UserResetPasswordVue },
            { path: '/access/:page', component: AccessVue },
        ],
    },
]

//创建路由器
const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

//导出路由
export default router
