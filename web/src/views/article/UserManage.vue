<script setup>
import { Delete } from '@element-plus/icons-vue'; // 导入点击框图标
import { ref,onMounted } from 'vue'; // 导入响应式数据
import { userListService } from '@/api/user.js'; // 导入查询所有的角色函数
// 定义用户名、角色名查询的模型，主要用于传输查询条件以及接收数据
const getUserByOptionModel = ref({
    "userName": '',
    "roleName": '',
})
const users = ref([]) // 定义接收user的响应式变量
//分页条数据模型
const pageNum = ref(1)//当前页
const total = ref(20)//总条数
const pageSize = ref(3)//每页条数
//当每页条数发生了变化，调用此函数
const onSizeChange = (size) => {
    pageSize.value = size
    console.log("onSizeChange调用userList")
    userList()
}
//当前页码发生变化，调用此函数
const onCurrentChange = (num) => {
    pageNum.value = num
    console.log("onCurrentChange调用userList")
    userList()
}


//文章分类数据模型
const categorys = ref([
    {
        "id": 3,
        "categoryName": "美食",
        "categoryAlias": "my",
        "createTime": "2023-09-02 12:06:59",
        "updateTime": "2023-09-02 12:06:59"
    },
    {
        "id": 4,
        "categoryName": "娱乐",
        "categoryAlias": "yl",
        "createTime": "2023-09-02 12:08:16",
        "updateTime": "2023-09-02 12:08:16"
    },
    {
        "id": 5,
        "categoryName": "军事",
        "categoryAlias": "js",
        "createTime": "2023-09-02 12:08:33",
        "updateTime": "2023-09-02 12:08:33"
    }
])

//用户搜索时选中的分类id
const categoryId = ref('')

//用户搜索时选中的发布状态
const state = ref('')

//文章列表数据模型
const articles = ref([
    {
        "id": 5,
        "title": "陕西旅游攻略",
        "content": "兵马俑,华清池,法门寺,华山...爱去哪去哪...",
        "coverImg": "https://big-event-gwd.oss-cn-beijing.aliyuncs.com/9bf1cf5b-1420-4c1b-91ad-e0f4631cbed4.png",
        "state": "草稿",
        "categoryId": 2,
        "createTime": "2023-09-03 11:55:30",
        "updateTime": "2023-09-03 11:55:30"
    },
    {
        "id": 5,
        "title": "陕西旅游攻略",
        "content": "兵马俑,华清池,法门寺,华山...爱去哪去哪...",
        "coverImg": "https://big-event-gwd.oss-cn-beijing.aliyuncs.com/9bf1cf5b-1420-4c1b-91ad-e0f4631cbed4.png",
        "state": "草稿",
        "categoryId": 2,
        "createTime": "2023-09-03 11:55:30",
        "updateTime": "2023-09-03 11:55:30"
    },
    {
        "id": 5,
        "title": "陕西旅游攻略",
        "content": "兵马俑,华清池,法门寺,华山...爱去哪去哪...",
        "coverImg": "https://big-event-gwd.oss-cn-beijing.aliyuncs.com/9bf1cf5b-1420-4c1b-91ad-e0f4631cbed4.png",
        "state": "草稿",
        "categoryId": 2,
        "createTime": "2023-09-03 11:55:30",
        "updateTime": "2023-09-03 11:55:30"
    },
])

// =============================== 查询所有用户代码 Start ====================================


// 获取所有的用户
const userList = async () => {
    users.value = [] // 获取数据之前先清空数据
    // 根据分页获取所有的用户数据
    let params = {
        pageNum: pageNum.value,
        pageSize: pageSize.value,
        userName: getUserByOptionModel.value.userName ? getUserByOptionModel.value.userName : null,
        roleName: getUserByOptionModel.value.roleName ? getUserByOptionModel.value.roleName : null,
    }
    let result = await userListService(params); // 向服务器发送get请求
    total.value = result.total
    console.log("获取从服务器返回的数据 : ", result.data, result.access)
    // 将数据按照格式进行整理
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let user = {
            "id": result.data[i].ID,
            "username": result.data[i].username,
            "email": result.data[i].email,
            "rolename":result.data[i].role.Name,
        }
        users.value.push(user)
    }
}

onMounted(() => {
    console.log("初始化加载userList");
    userList();
});


// =============================== 查询所有用户代码 End ====================================

//回显文章分类
// import { articleCategoryListService, articleListService,articleAddService } from '@/api/article.js'

// const articleCategoryList = async () => {
//     let result = await articleCategoryListService();
//     categorys.value = result.data;
// }


// //获取文章列表数据
// const articleList = async () => {
//     let params = {
//         pageNum: pageNum.value,
//         pageSize: pageSize.value,
//         categoryId: categoryId.value ? categoryId.value : null,
//         state: state.value ? state.value : null
//     }
//     let result = await articleListService(params);

//     //渲染视图
//     total.value = result.data.total;
//     articles.value = result.data.items;

//     //处理数据,给数据模型扩展一个属性categoryName,分类名称
//     for (let i = 0; i < articles.value.length; i++) {
//         let article = articles.value[i];
//         for (let j = 0; j < categorys.value.length; j++) {
//             if (article.categoryId == categorys.value[j].id) {
//                 article.categoryName = categorys.value[j].categoryName;
//             }
//         }
//     }
// }


// articleCategoryList()
// articleList();

import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css'
import { Plus } from '@element-plus/icons-vue'
//控制抽屉是否显示
const visibleDrawer = ref(false)
//添加表单数据模型
const articleModel = ref({
    title: '',
    categoryId: '',
    coverImg: '',
    content: '',
    state: ''
})


//导入token
import { useTokenStore } from '@/stores/token.js';
const tokenStore = useTokenStore();
console.log("token = ", tokenStore.token);

//上传成功的回调函数
const uploadSuccess = (result)=>{
    articleModel.value.coverImg = result.data;
    console.log(result.data);
}

//添加文章
import {ElMessage} from 'element-plus'
const addArticle = async (clickState)=>{
    //把发布状态赋值给数据模型
    articleModel.value.state = clickState;

    //调用接口
    let result = await articleAddService(articleModel.value);

    ElMessage.success(result.msg? result.msg:'添加成功');

    //让抽屉消失
    visibleDrawer.value = false;

    //刷新当前列表
    articleList()


}

// ============================ 搜索用户功能代码 ============================

/*
用户管理搜索功能实现
主要的功能
1. 根据用户名 && 角色进行查询
*/


// 接收所有的角色名字
const roles= ref([])

// 获取所有的角色名
import { roleListService } from '@/api/role.js'

const getAllRoles = async() => {
    // 向服务器端请求数据
    let result = await roleListService();
    // 将所有的角色名得到
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let role = {
            "id": result.data[i].ID,
            "roleName": result.data[i].Name,
        }
        roles.value.push(role)
    }
}

// 一加载界面直接调用函数
getAllRoles();

// 清空数据实现
const clearUserByOption = () => {
    getUserByOptionModel.value.userName = '';
    getUserByOptionModel.value.roleName = '';
    userList();
}

// import { getUserByOptionService } from '@/api/user.js'

// // 查找用户功能实现
// const getUserByOption = async(getUserByOptionModel) => {
//     let result = await getUserByOptionService(getUserByOptionModel)
    
//     users.value = result.data
// }

import {ElMessageBox} from 'element-plus'
import { deleteUserService } from '@/api/user.js'
const deleteUser = (row) => {
    //提示用户  确认框
    ElMessageBox.confirm(
        '你确认要删除该角色信息吗?',
        '温馨提示',
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
        .then(async () => {
            //调用接口
            let result = await deleteUserService(row.id);
            ElMessage({
                type: 'success',
                message: '删除成功',
            })
            //刷新列表
            userList();
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '管理员取消了删除',
            })
        })
}

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

const addUserDialogVisible = ref(false)

const showAddUserDialogVisible = () => {
    addUserDialogVisible.value = true
    console.log("addUserModel = ", addUserModel)
}

// 定义添加角色模型
const addUserModel = ref({
    "userName": '',
    "email": '',
    "role": {},
})


const handleRoleChange = (value) => {
    // 根据选中项的值更新整个角色对象
    addUserModel.value.role = value;
};

const addUser = async() => {
    let result = await addUserService(addUserModel)
}

const addUserRules = {
    userName: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
    ],
    email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
    ],
}



</script>

<template>
    <el-card class="page-container">
        
        <template #header>
            <div class="header">
                <span>用户管理</span>
                <div class="extra">
                    <el-button type="primary" @click="showAddUserDialogVisible">添加用户</el-button>
                </div>
            </div>
        </template>
   
        <el-form inline>
            <el-form-item label="用户名：">
                <el-input placeholder="请输入用户名" v-model="getUserByOptionModel.userName"></el-input>
            </el-form-item>
            <el-form-item label="角色名">
                <el-select placeholder="请选择" v-model="getUserByOptionModel.roleName">
                    <el-option v-for="role in roles" :key="role.id" :label="role.roleName" :value="role.roleName"></el-option>
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
                    <el-button :icon="Delete" circle plain type="danger" @click="deleteUser(row)"></el-button>
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
        <el-dialog v-model="addUserDialogVisible" :title="添加角色" width="30%">
            <el-form :model="addUserModel" :rules="addUserRules" label-width="100px" style="padding-right: 30px">
                <el-form-item label="用户名" prop="userName">
                    <el-input v-model="addUserModel.userName" minlength="1" maxlength="10"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="addUserModel.email"></el-input>
                </el-form-item>
                <el-form-item label="设置角色">
                    <el-select placeholder="请选择" v-model="addUserModel.role.roleName" @change="handleRoleChange">
                        <el-option v-for="role in roles" :key="role.id" :label="role.roleName" :value="role"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="addUser"> 确认 </el-button>
                </span>
            </template>
        </el-dialog>

        <!-- 抽屉 -->
        <el-drawer v-model="visibleDrawer" title="添加文章" direction="rtl" size="50%">
            <!-- 添加文章表单 -->
            <el-form :model="articleModel" label-width="100px">
                <el-form-item label="文章标题">
                    <el-input v-model="articleModel.title" placeholder="请输入标题"></el-input>
                </el-form-item>
                <el-form-item label="文章分类">
                    <el-select placeholder="请选择" v-model="articleModel.categoryId">
                        <el-option v-for="c in categorys" :key="c.id" :label="c.categoryName" :value="c.id">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="文章封面">

                    <!-- 
                        auto-upload:设置是否自动上传
                        action:设置服务器接口路径
                        name:设置上传的文件字段名
                        headers:设置上传的请求头
                        on-success:设置上传成功的回调函数
                     -->
                   
                    <el-upload class="avatar-uploader" :auto-upload="true" :show-file-list="false"
                    action="/api/upload"
                    name="file"
                    :headers="{'Authorization':tokenStore.token}"
                    :on-success="uploadSuccess"
                    >
                        <img v-if="articleModel.coverImg" :src="articleModel.coverImg" class="avatar" />
                        <el-icon v-else class="avatar-uploader-icon">
                            <Plus />
                        </el-icon>
                    </el-upload>
                </el-form-item>
                <el-form-item label="文章内容">
                    <div class="editor">
                        <quill-editor theme="snow" v-model:content="articleModel.content" contentType="html">
                        </quill-editor>
                    </div>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="addArticle('已发布')">发布</el-button>
                    <el-button type="info" @click="addArticle('草稿')">草稿</el-button>
                </el-form-item>
            </el-form>
        </el-drawer>
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

