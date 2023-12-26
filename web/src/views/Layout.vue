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
import { accessesListService } from '@/api/access.js'
import useUserInfoStore from '@/stores/userInfo.js'
import {useTokenStore} from '@/stores/token.js'
import {ref, onMounted} from 'vue';
const tokenStore = useTokenStore();
const userInfoStore = useUserInfoStore();
const accesses = ref([]);
const selectedPage = ref({});
// 获取权限
const accessList = async() => {
    clearAccesses(); // 先清空数据
    let result = await accessesListService(); // 获取权限数据
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
onMounted(() => {
    console.log("加载权限数据 = ", accesses);
    accessList();
});
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
            tokenStore.removeToken()
            userInfoStore.removeInfo()
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

import { useRoute } from 'vue-router';
const route = useRoute();

const handleMenuItemClick = (index, access) => {
    const currentPath = route.path
    console.log("path = ", currentPath)
    const token = tokenStore.token
    if (token.user.admin) {
        // 管理员具有所有权限
        router.push({ path: index }); 
        ElMessage.success('访问成功');
        return
    }
    if (access === null) {
        router.push({ path: currentPath }); 
        ElMessage.warning('您没有访问权限！')
        return 
    }
    let level = 0;
    for (let i = 0 ; i < token.access.length ; i ++) {
        if (access.id === token.access[i].Access.ID) {
            level = token.access[i].Level
            break
        }
    }
    if (level >= 0) {
        router.push(`/access/${access.id}`)
        ElMessage.success('访问成功')
    } else {
        router.push({ path: currentPath }); 
        ElMessage.warning('您没有访问权限！')
    }
};
    

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
                <el-menu-item index="/article/user" @click="handleMenuItemClick('/article/user', null)">
                    <el-icon>
                        <Management />
                    </el-icon>
                    <span>用户管理</span>
                </el-menu-item>

                <el-menu-item index="/article/role" @click="handleMenuItemClick('/article/role', null)">
                    <el-icon>
                        <Management />
                    </el-icon>
                    <span>角色管理</span>
                </el-menu-item>

                <el-menu-item index="/article/access" @click="handleMenuItemClick('/article/access', null)">
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
                    @click="handleMenuItemClick(index, access)"
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
                <div>黑马程序员：<strong>{{ userInfoStore.info.nickname }}</strong></div>
                <!-- 下拉菜单 -->
                <!-- command: 条目被点击后会触发,在事件函数上可以声明一个参数,接收条目对应的指令 -->
                <el-dropdown placement="bottom-end" @command="handleCommand">
                    <span class="el-dropdown__box">
                        <el-avatar :src="userInfoStore.info.userPic? userInfoStore.info.userPic:avatar" />
                        <el-icon>
                            <CaretBottom />
                        </el-icon>
                    </span>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item command="info" :icon="User">基本资料</el-dropdown-item>
                            <el-dropdown-item command="avatar" :icon="Crop">更换头像</el-dropdown-item>
                            <el-dropdown-item command="resetPassword" :icon="EditPen">重置密码</el-dropdown-item>
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
            <el-footer>大事件 ©2023 Created by 黑马程序员</el-footer>
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
            background: url('@/assets/logo.png') no-repeat center / 120px auto;
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