<script setup>
import { Plus, Delete, Edit } from '@element-plus/icons-vue'; // 导入点击框图标
import { ref,onMounted } from 'vue'; // 导入响应式数据
import { userListService } from '@/api/user.js'; // 导入查询所有的角色函数
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import { useTokenStore } from '@/stores/token.js'; //导入token
import {ElMessageBox} from 'element-plus';
import { deleteUserService, updateUserService } from '@/api/user.js';
import { roleListService } from '@/api/role.js'
import {ElMessage} from 'element-plus'
const tokenStore = useTokenStore(); 
const token = tokenStore.token
// =============================== 查询所有用户代码 Start ====================================
const users = ref([]) // 定义接收users的响应式变量
//分页条数据模型
const pageNum = ref(1)//当前页
const total = ref(20)//总条数
const pageSize = ref(3)//每页条数
//当每页条数发生了变化，调用此函数
const onSizeChange = (size) => {
    pageSize.value = size
    userList()
}
//当前页码发生变化，调用此函数
const onCurrentChange = (num) => {
    pageNum.value = num
    userList()
}
// 获取所有用户的函数
const userList = async () => {
    console.log('调用userList')
    users.value = [] // 获取数据之前先清空数据
    // 根据分页获取所有的用户数据
    let params = {
        pageNum: pageNum.value,
        pageSize: pageSize.value,
        userName: option.value.username ? option.value.username : null,
        roleName: option.value.rolename ? option.value.rolename : null,
    }
    let result = await userListService(params); // 向服务器发送get请求
    total.value = result.total
    // 将数据按照格式进行整理
    console.log("后端返回的user数据", result.data)
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let user = {
            "id": result.data[i].ID,
            "username": result.data[i].username,
            "email": result.data[i].email,
            "roleid": result.data[i].roleid,
            "rolename":result.data[i].role.rolename,
        }
        users.value.push(user)
    }
}
// 查询用户条件
const option = ref({
    "username": '',
    "rolename": '',
})

const handleOptionChange = (value) => {
    console.log(option.value.rolename)
    option.value.rolename = value.rolename;
}
// 清除查询用户条件
const clearUserByOption = () => {
    option.value.username = '';
    option.value.rolename = '';
    userList();
}
// =============================== 查询所有用户代码 end ====================================
// =============================== 查询所有角色代码 Start ==================================
const roles= ref([]) // 接收所有的角色名字
const roleList = async() => {
    // 向服务器端请求数据
    let result = await roleListService();
    // 将所有的角色名得到
    console.log("后端返回的role数据", result.data)
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let role = {
            "id": result.data[i].ID,
            "rolename": result.data[i].rolename,
        }
        roles.value.push(role)
    }
}
// =============================== 查询所有角色代码 end ==================================
// =============================== 数据初始化代码 Start ==================================
// 组件渲染结束后加载数据
onMounted(() => {
    // 加载用户数据
    userList();
    console.log('加载用户数据', users)
    // 加载角色数据
    roleList()
    console.log('加载角色数据',roles)
});
// ============================ 编辑用户功能代码 ============================
const dialogVisible = ref(false)

const userModel = ref({
    "id": '',
    "username": '',
    "email": '',
    "rolename": '',
    "roleid":'',
})

const handleRoleChange = (value) => {
    userModel.value.roleid = value.id;
    userModel.value.rolename = value.rolename;
}

const showDialog = (row) => {
    userModel.value.id = row.id;
    userModel.value.username = row.username;
    userModel.value.email = row.email;
    userModel.value.rolename = row.rolename;
    userModel.value.roleid = row.roleid;
    dialogVisible.value = true;
}

const updateUser = async() => {
    let result = await updateUserService(userModel.value.id, userModel.value.roleid);
    ElMessage.success(result.msg ? result.msg : "success");
    dialogVisible.value = false
    userList();
}


// ============================ 删除用户功能代码 ============================
const deleteUser = (row) => {
    //提示用户  确认框
    ElMessageBox.confirm('你确认要删除该角色信息吗?', '温馨提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(async () => {
        //调用接口
        let result = await deleteUserService(row.id);
        ElMessage({ type: 'success', message: '删除成功'});
        userList(); //刷新列表
    }).catch(() => { ElMessage({ type: 'info', message: '管理员取消了删除'})})
}


/*
用户管理搜索功能实现
主要的功能
1. 根据用户名 && 角色进行查询
*/

// import { getUserByOptionService } from '@/api/user.js'

// // 查找用户功能实现
// const getUserByOption = async(getUserByOptionModel) => {
//     let result = await getUserByOptionService(getUserByOptionModel)
    
//     users.value = result.data
// }


// ============================ 添加用户功能代码 ============================

/* 
添加用户实现 
roles存储所有角色的信息
{
    "id"
    "roleName"
}
*/

// 控制添加角色的对话框 false代表不现实

// const addUserDialogVisible = ref(false)

// const showAddUserDialogVisible = () => {
//     addUserDialogVisible.value = true
    
//     console.log("addUserModel = ", addUserModel)
// }


// const handleRoleChange = (value) => {
//     // 根据选中项的值更新整个角色对象
//     addUserModel.value.role = value;
// };

// const addUser = async() => {
//     let result = await addUserService(addUserModel)
// }



</script>

<template>
    <el-card class="page-container">
        
        <template #header>
            <div class="header">
                <span>用户管理</span>
                <!-- <div class="extra">
                    <el-button type="primary" @click="dialogVisible = true ; clearUserModel();">添加用户</el-button>
                </div> -->
            </div>
        </template>
   
        <el-form inline>
            <el-form-item label="用户名：">
                <el-input placeholder="请输入用户名" v-model="option.username"></el-input>
            </el-form-item>
            <el-form-item label="角色名">
                <el-select placeholder="请选择" v-model="option.rolename" @change="handleOptionChange">
                    <el-option 
                    v-for="role in roles" 
                    :key="role.id" 
                    :label="role.rolename" 
                    :value="role"
                    >
                </el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="userList">搜索</el-button>
                <el-button @click="clearUserByOption">重置</el-button>
            </el-form-item>
        </el-form>

        <!-- 用户列表 -->
        <el-table :data="users" style="width: 100%">
            <el-table-column label="用户名" width="300" prop="username"></el-table-column>
            <el-table-column label="用户邮箱" prop="email"></el-table-column>
            <el-table-column label="用户角色" prop="rolename"> </el-table-column>
            <el-table-column label="操作" width="100">
                <template #default="{ row }">
                    <el-button :icon="Edit" circle plain type="primary" round @click="showDialog(row)"></el-button>
                    <el-button :icon="Delete" circle plain type="danger" round @click="deleteUser(row)"></el-button>
                </template>
            </el-table-column>
            <template #empty>
                <el-empty description="没有数据" />
            </template>
        </el-table>
        <!-- 分页条 -->
        <el-pagination v-model:current-page="pageNum" v-model:page-size="pageSize" :page-sizes="[3, 5, 10, 15]"
            layout="jumper, total, sizes, prev, pager, next" background :total="total" @size-change="onSizeChange"
            @current-change="onCurrentChange" style="margin-top: 20px; justify-content: flex-end" />

        
        <!-- 添加角色对话框 -->
        <el-dialog v-model="dialogVisible" title="更改用户角色" width="30%">
            <el-form :model="userModel" label-width="100px" style="padding-right: 30px">
                <el-form-item label="用户名">
                    <el-input v-model="userModel.username" disabled></el-input>
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model="userModel.email" disabled></el-input>
                </el-form-item>
                <el-form-item label="设置角色">
                    <el-select placeholder="请选择" v-model="userModel.rolename" @change="handleRoleChange">
                        <el-option v-for="role in roles" :key="role.id" :label="role.rolename" :value="role"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="updateUser"> 确认 </el-button>
                </span>
            </template>
        </el-dialog>
    </el-card>
</template>

<style lang="scss" scoped>
.page-container {
    min-height: 100%;
    box-sizing: border-box;

    .header {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
}

/* 抽屉样式 */
.avatar-uploader {
    :deep() {
        .avatar {
            width: 178px;
            height: 178px;
            display: block;
        }

        .el-upload {
            border: 1px dashed var(--el-border-color);
            border-radius: 6px;
            cursor: pointer;
            position: relative;
            overflow: hidden;
            transition: var(--el-transition-duration-fast);
        }

        .el-upload:hover {
            border-color: var(--el-color-primary);
        }

        .el-icon.avatar-uploader-icon {
            font-size: 28px;
            color: #8c939d;
            width: 178px;
            height: 178px;
            text-align: center;
        }
    }
}

.editor {
    width: 100%;

    :deep(.ql-editor) {
        min-height: 200px;
    }
}
</style>

