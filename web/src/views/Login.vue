<script setup>
import { User, Lock, Message } from '@element-plus/icons-vue' // 导入element图标库
import { ref } from 'vue' // 导入vue响应式数据
import { ElMessage } from 'element-plus' // 导入消息窗口显示模块
import { userRegisterService, userLoginService} from '@/api/user.js' // 导入函数
import { useRouter } from 'vue-router' // 导入router
// 导入token相关信息
import { useAdminStore } from '@/stores/admin.js'
const adminStore = useAdminStore();
import { useUserStore } from '@/stores/user.js'
const userStore = useUserStore();
import { useAccessStore } from '@/stores/access.js'
const accessStore = useAccessStore();
const isRegister = ref(false) //控制注册与登录表单的显示，false显示为登录 true显示为注册
//定义数据模型
const registerData = ref({
    username: '',
    password: '',
    rePassword: '',
    email: '',
});
//定义函数,清空数据模型的数据
const clearRegisterData = ()=>{
    registerData.value={
        username:'',
        password:'',
        rePassword:''
    }
}
const router = useRouter() // 获取路由
//校验密码的函数
const checkRePassword = (rule, value, callback) => {
    if (value === '') {
        callback(new Error('请再次确认密码'))
    } else if (value !== registerData.value.password) {
        callback(new Error('请确保两次输入的密码一样'))
    } else {
        callback()
    }
}
//定义表单校验规则
const rules = {
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 5, max: 16, message: '长度为5~16位非空字符', trigger: 'blur' }
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 5, max: 16, message: '长度为5~16位非空字符', trigger: 'blur' }
    ],
    rePassword: [
        { validator: checkRePassword, trigger: 'blur' }
    ]
}
//调用后台接口,完成注册
const register = async () => {
    let result = await userRegisterService(registerData.value);
    ElMessage.success(result.msg ? result.msg : '注册成功')
    // 如果注册成功，则跳转界面
    if (result.code === 0) {
        isRegister.value = false
    }
}

const login = async ()=>{
    //调用接口,完成登录
   let result =  await userLoginService(registerData.value);
   ElMessage.success(result.msg ? result.msg : '登录成功')
   //把得到的token存储到pinia中
   console.log('登录成功返回的结果为', result.data);
   let data = result.data;
   if (data.admin) {
    // 如果登录的用户是管理员
    adminStore.setToken(data);
   } else {
    userStore.setToken(data);
    // 加载用户权限
    initAccess();
   }
   console.log('adminToken = ', adminStore.adminToken);
   console.log('userToken = ', userStore.userToken);
   router.push('/')
}

import { accessListService } from '@/api/access.js'

const initAccess = async() => {
    let userToken = userStore.userToken;
    let result = await accessListService(userToken.roleid, 0);
    accessStore.setToken(result.data);
    console.log('accessToken = ', accessStore.accessToken);
}
</script>

<template>
    <el-row class="login-page">
        <el-col :span="12" class="bg"></el-col>
        <el-col :span="6" :offset="3" class="form">
            <!-- 注册表单 -->
            <el-form ref="form" size="large" autocomplete="off" v-if="isRegister" :model="registerData" :rules="rules">
                <el-form-item>
                    <h1>注册</h1>
                </el-form-item>
                <el-form-item prop="username">
                    <el-input :prefix-icon="User" placeholder="请输入用户名" v-model="registerData.username"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input :prefix-icon="Lock" type="password" placeholder="请输入密码"
                        v-model="registerData.password"></el-input>
                </el-form-item>
                <el-form-item prop="rePassword">
                    <el-input :prefix-icon="Lock" type="password" placeholder="请输入再次密码"
                        v-model="registerData.rePassword"></el-input>
                </el-form-item>
                <el-form-item prop="eamil">
                    <el-input :prefix-icon="Message" placeholder="请输入邮箱"
                        v-model="registerData.email"></el-input>
                </el-form-item>
                <!-- 注册按钮 -->
                <el-form-item>
                    <el-button class="button" type="primary" auto-insert-space @click="register">
                        注册
                    </el-button>
                </el-form-item>
                <el-form-item class="flex">
                    <el-link type="info" :underline="false" @click="isRegister = false;clearRegisterData()">
                        ← 返回
                    </el-link>
                </el-form-item>
            </el-form>
            <!-- 登录表单 -->
            <el-form ref="form" size="large" autocomplete="off" v-else :model="registerData" :rules="rules">
                <el-form-item>
                    <h1>登录</h1>
                </el-form-item>
                <el-form-item prop="username">
                    <el-input :prefix-icon="User" placeholder="请输入用户名" v-model="registerData.username"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input name="password" :prefix-icon="Lock" type="password" placeholder="请输入密码" v-model="registerData.password"></el-input>
                </el-form-item>
                <el-form-item class="flex">
                    <div class="flex">
                        <el-checkbox>记住我</el-checkbox>
                        <el-link type="primary" :underline="false">忘记密码？</el-link>
                    </div>
                </el-form-item>
                <!-- 登录按钮 -->
                <el-form-item>
                    <el-button class="button" type="primary" auto-insert-space @click="login">登录</el-button>
                </el-form-item>
                <el-form-item class="flex">
                    <el-link type="info" :underline="false" @click="isRegister = true;clearRegisterData()">
                        注册 →
                    </el-link>
                </el-form-item>
            </el-form>
        </el-col>
    </el-row>
</template>

<style lang="scss" scoped>
/* 样式 */
.login-page {
    height: 100vh;
    background-color: #fff;

    .bg {
        // background: url('@/assets/logo2.png') no-repeat 60% center / 240px auto,
        background: url('@/assets/login_bg.jpg') no-repeat center / cover;
        border-radius: 0 20px 20px 0;
    }

    .form {
        display: flex;
        flex-direction: column;
        justify-content: center;
        user-select: none;

        .title {
            margin: 0 auto;
        }

        .button {
            width: 100%;
        }

        .flex {
            width: 100%;
            display: flex;
            justify-content: space-between;
        }
    }
}
</style>