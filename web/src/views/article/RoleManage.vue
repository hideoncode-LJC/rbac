<script setup>

// 导入按钮样式 
import { View, Edit, Delete } from '@element-plus/icons-vue'
// 导入响应式数据模块
import { ref } from 'vue'

// 导入函数模块
import { roleListService, addRoleService } from '@/api/role.js'
import { accessesListService, getAccessByRoleIdService } from '@/api/access.js'
import {ElMessageBox} from 'element-plus'
// ------------------- 接收所有角色代码 start 
// 接收所有角色 type role {"id":, "rolename":}
// 定义接收数据的变量
const roles = ref([])

// 声明获取所有角色的函数
const rolesList = async() =>  {
    let result = await roleListService()
    // 遍历循环获取数据
    for (let i = 0 ; i < result.data.length ; i ++) {
        let role = {
            "id": result.data[i].ID,
            "name": result.data[i].Name,
        }
        roles.value.push(role)
    }
}
// 一显示该界面就直接获取数据
rolesList()
// --------------------------- 获取所有角色代码 end --------------------------

// --------------------------- 添加角色逻辑代码 Start -------------------------
// 定义控制窗口的参数 默认为false， 当点击添加角色时 修改为 true
const addRoleDialogVisible = ref(false)
// 定义添加角色模块，用于存储新添加的角色的对应权限
const addRoleModel = ref ({
    roleName: '',
    accesses: [],
})
// 定义表单的校验规则
const addRoleRules = {
    roleName :[{ required: true, message: '请输入分类别名', trigger: 'blur' }],
    accesses: [{}]
}

//展示添加角色弹窗
const showAddRoleDialog = () => {
    // 显示弹窗
    addRoleDialogVisible.value = true; 
    // 设置对话框的标题
    title.value = '添加角色'
    //清空添加角色模型的数据
    addRoleModel.value.roleName = ''
    addRoleModel.value.accesses = []
    // 获取所有的权限名字
    accessesList()
}

const accessesList = async() => {
    let result = await accessesListService();
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let access = {
            "id" : result.data[i].ID,
            "title": result.data[i].Title,
            "content": result.data[i].Content,
            "level":['无权限']
        }
        if (i === 0) {
            addRoleSelectedAccess.value = access
        }
        addRoleModel.value.accesses.push(access)
    }
}
// 添加权限的选项
const accessOptions = ref(["无权限","读", "写", "删除"])
// 绑定每个下拉菜单的选项
const addRoleSelectedAccess = ref({
    "id":'',
    "title":'',
    "content":'',
    "level": ['无权限'],
})
// 调用函数
const addRole = async ()=> {
    // 调用接口
    let result = await addRoleService(addRoleModel.value)
    // 显示返回信息
    ElMessage.success(result.msg ? result.msg : '添加成功')
    // 获取所有角色
    rolesList()
    // 关闭窗口
    addRoleDialogVisible = false
}
// 监控权限表单的修改
const handlerAddRoleSelectedAccessLevel = (value) => {
    if (value.includes('删除')) {
        value = ['读','写','删除']
    } else if (value.includes('写')) {
        value = ['读','写']
    } else if (value.includes('读')) {
        value = ['读']
    } else {
        value = ['无权限']
    }
    addRoleSelectedAccess.value.level = value
} 
// ------------------------------ 添加角色逻辑代码 End ---------------------------

// ------------------------------ 修改角色逻辑代码 End ---------------------------

// 控制编辑窗口的显示
const updateRoleDialogVisible = ref(false)

// 定义接收该角色的权限数据

const updateRoleModel = ref({
    "id": '',
    "roleName": '',
    "accesses": '',
})

// 显示编辑窗口
const showUpdateRoleDialogVisible = (row) => {
    updateRoleDialogVisible.value = true;
    // 将接收数据的变量清空
    updateRoleModel.value.id = '';
    updateRoleModel.value.roleName = '';
    updateRoleModel.value.accesses = [];
    getAccessByRoleId(row.id)
}



const viewRoleDialogVisible = ref(false)

const roleModel = ref({
    roleID: '',
    roleName: '',
})

const roleRules = {
    roleName :[{
        required: true, message: '请输入分类别名', trigger: 'blur'
    }],
    accesses: [{}]
}

//添加分类数据模型
const categoryModel = ref({
    categoryName: '',
    categoryAlias: '',
    roleName: ''
})


//添加分类表单校验
const rules = {
    categoryName: [
        { required: true, message: '请输入分类名称', trigger: 'blur' },
    ],
    categoryAlias: [
        { required: true, message: '请输入分类别名', trigger: 'blur' },
    ],
    rolename: [
        { required: true, message: '请输入分类名称', trigger: 'blur' },
    ],
}


//调用接口,添加表单
import { ElMessage } from 'element-plus'
const addCategory = async () => {
    //调用接口
    let result = await articleCategoryAddService(categoryModel.value);
    ElMessage.success(result.msg ? result.msg : '添加成功')

    //调用获取所有文章分类的函数
    articleCategoryList();
    dialogVisible.value = false;
}



//定义变量,控制标题的展示
const title = ref('')



const showViewRoleDialog = (row) => {
    // 先清空数据
    roleModel.value.roleID = ''
    roleModel.value.roleName = ''
    // 获取权限数据
    getAccessByRoleId(row.id)
    title.value = '查看角色权限'
    roleModel.value.roleID = row.id
    roleModel.value.roleName = row.name
    // 显示界面
    viewRoleDialogVisible.value = true;
}

const showDialog = (row) => {

}

//编辑分类
const updateCategory = async () => {
    //调用接口
    let result = await articleCategoryUpdateService(categoryModel.value);

    ElMessage.success(result.msg ? result.msg : '修改成功')

    //调用获取所有分类的函数
    articleCategoryList();

    //隐藏弹窗
    dialogVisible.value = false;
}

//清空模型的数据
const clearData = () => {
    categoryModel.value.categoryName = '';
    categoryModel.value.categoryAlias = '';
}

//删除角色

const deleteCategory = (row) => {
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
            let result = await articleCategoryDeleteService(row.id);
            ElMessage({
                type: 'success',
                message: '删除成功',
            })
            //刷新列表
            articleCategoryList();
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '用户取消了删除',
            })
        })
}



// 定义接收返回权限的参数
const accessesById = ref([])

const getAccessByRoleId = async (id) => {
    // 先清空数据
    accessesById.value = [];
    let result = await getAccessByRoleIdService(id);
    for (let i = 0 ; i < result.data.length ; i ++ ) {

        let accessById = {
            "id":result.data[i].Access.ID,
            "title":result.data[i].Access.Title,
            "level":'',
        }
        switch (result.data[i].Level) {
            case 0:accessById.level = '无权限';break;
            case 1:accessById.level = '读权限';break;
            case 2:accessById.level = '读、写权限';break;
            case 3:accessById.level = '读、写、删除权限';break;
        }
        accessesById.value.push(accessById)
    } 
}

const accessForPage = ref({
    "id": '',
    "title": '',
    "level":'',
})

</script>
<template>
    <el-card class="page-container">
        <template #header>
            <div class="header">
                <span>角色分类</span>
                <div class="extra">
                    <el-button type="primary" @click="showAddRoleDialog()">添加角色</el-button>
                </div>
            </div>
        </template>
        <el-table :data="roles" style="width: 100%">
            <el-table-column label="序号" width="100" type="index"> </el-table-column>
            <el-table-column label="角色名称" prop="name" align="center"></el-table-column>
            <el-table-column label="操作" width="200">
                <template #default="{ row }">
                    <el-button :icon="View" circle plain type="info" @click="showViewRoleDialog(row)"></el-button>
                    <el-button :icon="Edit" circle plain type="primary" @click="showEditRoleDialog()"></el-button>
                    <el-button :icon="Delete" circle plain type="danger" @click="deleteCategory(row)"></el-button>
                </template>
            </el-table-column>
            <template #empty>
                <el-empty description="没有数据" />
            </template>
        </el-table>





        <!-- 添加新的角色弹窗 -->
        <el-dialog v-model="addRoleDialogVisible" :title="title" width="30%">
            <el-form :model="addRoleModel" :rules="addRoleRules" label-width="100px" style="padding-right: 30px">
                <el-form-item label="角色名称" prop="roleName">
                    <el-input v-model="addRoleModel.roleName" minlength="1" maxlength="10"></el-input>
                </el-form-item>
                <el-form-item label="权限界面">
                    <el-select v-model="addRoleSelectedAccess" value-key="id" placeholder="Select">
                        <el-option v-for="addRoleAccess in addRoleModel.accesses" 
                        :key="addRoleAccess.id" 
                        :label="addRoleAccess.title" 
                        :value="addRoleAccess"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="权限选择">
                    <el-checkbox-group v-model="addRoleSelectedAccess.level" size="large" @change="handlerAddRoleSelectedAccessLevel">
                        <el-checkbox v-for="option in accessOptions" 
                        :key="option" 
                        :label="option">
                        {{ option }}
                        </el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="addRoleDialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="addRole"> 确认 </el-button>
                </span>
            </template>
        </el-dialog>




        

        <!-- 添加分类弹窗 -->
        <el-dialog v-model="dialogVisible" :title="title" width="30%">
            <el-form :model="categoryModel" :rules="rules" label-width="100px" style="padding-right: 30px">
                <el-form-item label="分类名称" prop="categoryName">
                    <el-input v-model="categoryModel.categoryName" minlength="1" maxlength="10"></el-input>
                </el-form-item>
                <el-form-item label="分类别名" prop="categoryAlias">
                    <el-input v-model="categoryModel.categoryAlias" minlength="1" maxlength="15"></el-input>
                </el-form-item>
                <el-form-item label="角色别名" prop="rolename">
                    <el-input v-model="categoryModel.roleName" minlength="1" maxlength="15"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="title == '添加分类' ? addCategory() : updateCategory()"> 确认 </el-button>
                </span>
            </template>
        </el-dialog>

        <!-- 查看角色权限弹窗 -->
        <el-dialog v-model="viewRoleDialogVisible" :title="title" width="30%">
            <el-form :model="roleModel" :rules="roleRules" label-width="100px" style="padding-right: 30px">
                <!-- 第一行 显示角色名和数据 -->
                <el-form-item label="角色名" prop="roleName">
                    <el-input v-model="roleModel.roleName" disabled>{{ roleModel.roleName }}</el-input>
                </el-form-item>
                <!-- 第二行 左右两个下拉菜单 -->
                <el-form-item label="权限">
                    <el-select v-model="accessForPage" value-key="id" placeholder="Select">
                        <el-option v-for="access in accessesById" :key="access.id" :label="access.title" :value="access"/>
                    </el-select>
                    <p>
                    {{ roleModel.roleName }}对{{ accessForPage.title}}的权限为{{ accessForPage.level }}
                    </p>
                </el-form-item>
            </el-form>
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
</style>

<!-- <script setup>
// 存放没用的数据

// const isUpdating = ref(false) 

// watch(addRoleSelectedAccess, (newValue, oldValue) => {
//     console.log("isUpdateing = ", isUpdating.value)
//   // 检查是否已经在更新，以避免递归
//   if (isUpdating.value) {
//     return;
//   }

//   const newLevel = newValue.level;
//   const oldLevel = oldValue.level;



//   let updatedLevel;

//   if (newLevel.includes('删除')) {
//     updatedLevel = ['读', '写', '删除'];
//   } else if (newLevel.includes('写')) {
//     updatedLevel = ['读', '写'];
//   } else if (newLevel.includes('读')) {
//     updatedLevel = ['读'];
//   } else {
//     updatedLevel = ['无权限'];
//   }

//   // 直接设置属性
//   isUpdating.value = true;
//   console.log("isUpdateing = ", isUpdating.value)
//   addRoleSelectedAccess.value.level = updatedLevel;
//   isUpdating.value = false;

//   // 执行你的逻辑，比如更新权限选择
//   // ...

//   console.log('Level changed:', oldLevel, '->', newLevel);

// });

// const level = ref([])

// const handleAddRoleSelectedAccessChange = () => {
//     updateSelectedAccessLevel();
// }

// const updateSelectedAccessLevel = () => {
//     level = addRoleSelectedAccess.value.level;
//     const currentIndex = accessOptions.value.indexOf(selectedAccessLevel[0]);
//     // 更新权限选择
//     addRoleSelectedAccess.level = accessOptions.value.slice(0, currentIndex + 1);
// }

// const categorys = ref([
//     {
//         "id": 3,
//         "categoryName": "美食",
//         "categoryAlias": "my",
//         "createTime": "2023-09-02 12:06:59",
//         "updateTime": "2023-09-02 12:06:59"
//     },
//     {
//         "id": 4,
//         "categoryName": "娱乐",
//         "categoryAlias": "yl",
//         "createTime": "2023-09-02 12:08:16",
//         "updateTime": "2023-09-02 12:08:16"
//     },
//     {
//         "id": 5,
//         "categoryName": "军事",
//         "categoryAlias": "js",
//         "createTime": "2023-09-02 12:08:33",
//         "updateTime": "2023-09-02 12:08:33"
//     }
// ])

//声明一个异步的函数
// import { 
//     articleCategoryListService, 
//     articleCategoryAddService, 
//     articleCategoryUpdateService,
//     articleCategoryDeleteService } from '@/api/article.js'


// const articleCategoryList = async () => {
//     let result = await articleCategoryListService();
//     categorys.value = result.data;
// }

// articleCategoryList();
//控制添加分类弹窗

</script> -->