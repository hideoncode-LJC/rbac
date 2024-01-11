<script setup>
import {
    Management,
    Promotion,
    UserFilled,
    User,
    Crop,
    EditPen,
    SwitchButton,
    CaretBottom,
    Lock
} from '@element-plus/icons-vue'
import avatar from '@/assets/default.png'
import { accessListService } from '@/api/access.js'
import {ref, onMounted, onUnmounted} from 'vue';
const userInfoStore = useUserStore();
const accesses = ref([]);
import { useRoute } from 'vue-router';
const route = useRoute();

import { useAdminStore } from '@/stores/admin.js'
const adminStore = useAdminStore();
const adminToken = adminStore.adminToken;
import { useUserStore } from '@/stores/user.js'
const userStore = useUserStore();
import { useAccessStore } from '@/stores/access.js'
const accessStore = useAccessStore();
// 获取权限
const accessList = async() => {
    clearAccesses(); // 先清空数据
    let result = await accessListService(0, 0); // 获取权限数据
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let access = {
            "id": result.data[i].ID,
            "title": result.data[i].title,
            "content": result.data[i].content,
            "path":"/access/"
        }
        access.path += access.id
        accesses.value.push(access);
    }
}
// 清空权限数据
const clearAccesses = () => {
    accesses.value = [];
}

import { userLoginService } from '@/api/user.js'

const updateAccessToken = async() => {
    if (adminToken.admin) {
        return 
    }
    let result = await userLoginService(userStore.userToken);
    userStore.setToken(result.data);
    result = await accessListService(userStore.userToken.roleid, 0);
    console.log('result.data = ', result.data);
    accessStore.setToken(result.data);
    console.log('当前用户信息为', userStore.userToken, '对应的角色的权限为', accessStore.accessToken);
}

onMounted(() => {
    router.push('/index');
    accessList();
    updateAccessToken();
    window.addEventListener('refreshMenu', accessList);// 刷新界面时，重新根据用户获取权限
});
onUnmounted(() => {
    window.removeEventListener('refreshMenu', accessList);
})

//条目被点击后,调用的函数
import {useRouter} from 'vue-router'
const router = useRouter();
import { ElMessage, ElMessageBox } from 'element-plus'
const handleCommand = (command)=>{
    //判断指令
    if(command === 'logout'){
        //退出登录
        ElMessageBox.confirm('您确认要退出吗?', '温馨提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        })
        .then(async () => {
            //退出登录
            //1.清空pinia中存储的token以及个人信息
            // tokenStore.removeToken()
            // userInfoStore.removeInfo()
            //2.跳转到登录页面
            router.push('/login')
            ElMessage({
                type: 'success',
                message: '退出登录成功',
            })
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '用户取消了退出登录',
            })
        })
    }else{
        //路由
        router.push('/user/'+command)
    }
}



const handleMenuItemClick = (index, access) => {
    const accessToken = accessStore.accessToken;
    // 默认界面谁都可以访问
    if (index === '/index') {
        router.push(index);
        ElMessage.success('访问成功');
        return 
    }
    const currentPath = route.path
    console.log("path = ", currentPath)
    /*
    如果adminToken.admin 则可以访问所有界面
    */
    if (adminToken.admin) { // 管理员具有所有权限
        console.log('index = ', index);
        router.push(index); 
        ElMessage.success('访问成功');
        return
    }

    let pattern = /^\/article\//;

    if (pattern.test(index)) {
        router.push({ path: currentPath }); 
        ElMessage.warning('您没有访问权限！')
        return 
    } else {
        console.log("字符串不以'/article/'开头");
    }
    /*
    如果是用户
    */  
    let level = 0;
    for (let i = 0 ; i < accessToken.length ; i ++) {
        if (access.id === accessToken[i].ID) {
            level = accessToken[i].level
            break
        }
    }
    console.log('level = ', level);
    if (level > 0) {
        router.push(`/access/${access.id}`)
        ElMessage.success('访问成功')
    } else {
        checkChildLevel(access.id, userStore.userToken.roleid);
    }
};

import { childLevelService } from '@/api/level.js'

const checkChildLevel = async(accessid, roleid) => {
    
    let result = await childLevelService(accessid, roleid);
    console.log('result.data = ', result.data);
    if(result.data > 0) {
        console.log('result.data111 = ', result.data);
        router.push('/access/'+ accessid);
        ElMessage.success('访问成功')
    } else {
        const currentPath = route.path;
        router.push({ path: currentPath }); 
        ElMessage.warning('您没有访问权限！')
    }
}

console.log('adminToken = ', adminStore.adminToken.id)

</script>

<template>
    <!-- element-plus中的容器 -->
    <el-container class="layout-container">
        <!-- 左侧菜单 -->
        <el-aside width="200px">
            <div class="el-aside__logo"></div>
            <!-- element-plus的菜单标签 -->
            <el-menu active-text-color="#ffd04b" background-color="#232323"  text-color="#fff"
                router>
                <el-menu-item index="/index" @click="handleMenuItemClick('/index', '')">
                    <el-icon>
                        <Management />
                    </el-icon>
                    <span>系统首页</span>
                </el-menu-item>
                <el-menu-item index="/article/user" @click="handleMenuItemClick('/article/user', '')">
                    <el-icon>
                        <Management />
                    </el-icon>
                    <span>用户管理</span>
                </el-menu-item>

                <el-menu-item index="/article/role" @click="handleMenuItemClick('/article/role', '')">
                    <el-icon>
                        <Management />
                    </el-icon>
                    <span>角色管理</span>
                </el-menu-item>

                <el-menu-item index="/article/access" @click="handleMenuItemClick('/article/access', '')">
                    <el-icon>
                        <Promotion />
                    </el-icon>
                    <span>权限管理</span>
                </el-menu-item>

                <el-sub-menu >
                    <template #title>
                        <el-icon>
                            <UserFilled />
                        </el-icon>
                        <span>权限界面</span>
                    </template>
                    <!-- 使用 v-for 循环渲染子菜单项 -->
                    <el-menu-item 
                    v-for="(access, index) in accesses" 
                    :key="index" 
                    :index="access.path" 
                    @click="handleMenuItemClick('/access/' + access.id, access)"
                    >
                        <el-icon><Lock /></el-icon>
                        <span>{{ access.title }}</span>
                    </el-menu-item>
                    <!-- <el-menu-item index="/access/page">
                        <el-icon>
                            <User />
                        </el-icon>
                        <span>基本资料</span>
                    </el-menu-item>
                    <el-menu-item index="/user/avatar">
                        <el-icon>
                            <Crop />
                        </el-icon>
                        <span>更换头像</span>
                    </el-menu-item>
                    <el-menu-item index="/user/resetPassword">
                        <el-icon>
                            <EditPen />
                        </el-icon>
                        <span>重置密码</span>
                    </el-menu-item>  -->
                </el-sub-menu>
            </el-menu>
        </el-aside>
        <!-- 右侧主区域 -->
        <el-container>
            <!-- 头部区域 -->
            <el-header>
                <!-- {{ userInfoStore.info.nickname }} -->
                <div>
                    <strong>当前登录的用户为 : {{typeof adminStore.adminToken.id !== 'undefined' ? adminToken.username : userStore.userToken.username }}</strong>
                    <strong>当前登录的用户的角色名为 : {{ typeof adminStore.adminToken.id !== 'undefined' ? adminToken.rolename : userStore.userToken.rolename }}</strong>
                </div> 

                <!-- 下拉菜单 -->
                <!-- command: 条目被点击后会触发,在事件函数上可以声明一个参数,接收条目对应的指令 -->
                <el-dropdown placement="bottom-end" @command="handleCommand">
                    <span class="el-dropdown__box">
                        <!-- userInfoStore.info.userPic? userInfoStore.info.userPic:avatar -->
                        <el-avatar /> 
                        <el-icon>
                            <CaretBottom />
                        </el-icon>
                    </span>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <!-- <el-dropdown-item command="info" :icon="User">基本资料</el-dropdown-item> -->
                            <!-- <el-dropdown-item command="avatar" :icon="Crop">更换头像</el-dropdown-item> -->
                            <!-- <el-dropdown-item command="resetPassword" :icon="EditPen">重置密码</el-dropdown-item> -->
                            <el-dropdown-item command="logout" :icon="SwitchButton">退出登录</el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-header>
            <!-- 中间区域 -->
            <el-main>
                <!-- <div style="width: 1290px; height: 570px;border: 1px solid red;">
                    内容展示区
                </div> -->
                <router-view></router-view>
            </el-main>
            <!-- 底部区域 -->
            <el-footer>RBAC系统DEMO ©2024 Created</el-footer>
        </el-container>
    </el-container>
</template>

<style lang="scss" scoped>
.layout-container {
    height: 100vh;

    .el-aside {
        background-color: #232323;

        &__logo {
            height: 120px;
            background: url('@/assets/logo3.png') no-repeat center / 120px auto;
        }

        .el-menu {
            border-right: none;
        }
    }

    .el-header {
        background-color: #fff;
        display: flex;
        align-items: center;
        justify-content: space-between;

        .el-dropdown__box {
            display: flex;
            align-items: center;

            .el-icon {
                color: #999;
                margin-left: 10px;
            }

            &:active,
            &:focus {
                outline: none;
            }
        }
    }

    .el-footer {
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 14px;
        color: #666;
    }
}
</style>