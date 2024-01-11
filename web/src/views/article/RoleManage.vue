<script setup>
import { View, Edit, Delete } from '@element-plus/icons-vue' // 导入按钮样式 
import { ref, onMounted, watch} from 'vue' // 导入响应式数据模块
import { roleListService, addRoleService, deleteRoleService, updateRoleService } from '@/api/role.js' // 导入函数模块
import { accessListService} from '@/api/access.js'
import { ElMessage, ElMessageBox } from 'element-plus'
const roles = ref([]); // 定义接收角色的数据
const roleList = async() =>  { // 声明获取所有角色的函数
    roles.value = [];
    let result = await roleListService();
    for (let i = 0 ; i < result.data.length ; i ++) { // 遍历循环获取数据
        if (result.data[i].rolename === '管理员') {
            continue;
        }
        let role = {
            "id": result.data[i].ID,
            "rolename": result.data[i].rolename,
        };
        roles.value.push(role);
    }
}

const selectedRole = ref([]);


onMounted(() => {
    console.log("从后端获取到的角色数据为", roles)
    roleList(); // 一显示该界面就直接获取数据
});
const title = ref('')
const dialogVisible = ref(false); // 定义对话框的检测
const roleModel = ref({
    id: '',
    rolename: '',
    childrole: [],
})
const rules = {
    rolename: [{ required: true, message: '请输入分类别名', trigger: 'blur' }],
};

const setRoleModel = (id, rolename) => {
    roleModel.value.id = id;
    roleModel.value.rolename = rolename;
};
const accesses = ref([]);
const accessList = async(id) => {
    accesses.value = []; // 清空数据
    let result = await accessListService(id, 0);
    console.log('后台返回的权限数据为 = ', result.data);
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let access = {
            "id" : result.data[i].ID,
            "title": result.data[i].title,
            "content": result.data[i].content,
            "level": result.data[i].level,
        }
        if (id !== 0 && i === 0) {
            selectedAccess.value = access.title;
            levelToOption(parseInt(access.level));
            console.log(selectedAccess.value, selectedOption.value);
        }
        accesses.value.push(access);
    }
}
const isDisabled = ref(false);
const isInputDisabled = ref(false);
const editVisible = ref(false)
const showDialog = (row, i) => {
    if (i === 0) {
        title.value = '添加角色'; // 设置title
        setRoleModel('', ''); // 清空数据
        accessList(0);
        selectedRole.value = [];
        editVisible.value = true; 
        isInputDisabled.value = false;
        isDisabled.value = false;
    } else if (i === 1) { // 查看权限
        title.value = '查看角色';
        setRoleModel(row.id, row.rolename); // 设置RoleModel
        accessList(row.id) // 获取该用户的权限数据
        // 获取该用户的子角色
        childRoleList(row.id);
        editVisible.value = false;
        isInputDisabled.value = true;
        isDisabled.value = true;
        // 设置初始值
    } else {
        title.value = '修改角色';
        setRoleModel(row.id, row.rolename);
        accessList(row.id) // 获取该用户的权限数据
        childRoleList(row.id);
        editVisible.value = true;
        isInputDisabled.value = false;
        isDisabled.value = false;
    }
    dialogVisible.value = true; // 显示对话框
}
// 当前选中的权限名 用字符串表示
const selectedAccess = ref('');

// 当前选中权限的权限
const selectedOption = ref(['无权限']);

const handleAccessChange = (value) => {
    selectedAccess.value = value.title;
}
const handleOption = (value) => {
    if (value.includes('删除')) {
        selectedOption.value = ['读','写','删除'];
    } else if (value.includes('写')) {
        selectedOption.value = ['读','写'];
    } else if (value.includes('读')) {
        selectedOption.value = ['读'];
    } else {
        selectedOption.value = ['无权限'];
    }
}
watch(() => selectedAccess.value, (newTitle, oldTitle) => {
    if (oldTitle === '') {
        return 
    }
    const oldAccess = accesses.value.find(access => access.title === oldTitle);
    oldAccess.level = optionToLevel();
    const newAccess = accesses.value.find(access => access.title === newTitle);
    levelToOption(newAccess.level);
});

import {childRoleListSerivice} from '@/api/role.js'

const childRoleList = async(id) => {
    let result = await childRoleListSerivice(id);
    console.log('该角色的子角色为', result.data)
    selectedRole.value = [];
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        selectedRole.value.push(result.data[i].ID);
    }
    console.log('selectedRole = ', selectedRole.value);
}
const optionToLevel = () => {
    if (selectedOption.value.includes('删除')) {
        return 3;
    } else if (selectedOption.value.includes('写')) {
        return 2;
    } else if (selectedOption.value.includes('读')) {
        return 1;
    } else {
        return 0;
    }
}
const levelToOption = (level) => {
    if (level === 3) {
        selectedOption.value = ['删除', '读', '写'];
    } else if (level === 2) {
        selectedOption.value = ['读', '写'];
    } else if (level === 1) {
        selectedOption.value = ['读'];
    } else {
        selectedOption.value = ['无权限'];
    }
    console.log('调用函数', selectedOption.value);
}
// 添加权限的选项
const options = ref(['无权限','读', '写', '删除'])
// 调用函数
const addRole = async ()=> {
    console.log('添加角色对话框的标题:', title.value)
    // 最后一次没有监听到select的改变，可能会不改变值 将提交时最后选中的设置
    const lastAccess = accesses.value.find(access => access.title === selectedAccess.value);
    lastAccess.level = optionToLevel();
    // 调用接口
    if(title.value === '添加角色') {
        const postData = {
            "rolename": roleModel.value.rolename,
            "accesses": accesses.value,
            "childrole": selectedRole.value,
        }
        console.log('PostData = ', postData);
        let result = await addRoleService(postData);
        ElMessage.success(result.msg ? result.msg : '添加成功') // 显示返回信息
    } else if(title.value === '修改角色') {
        const putData = {
            "rolename": roleModel.value.rolename,
            "accesses": accesses.value,
            "childrole": selectedRole.value,
        }
        console.log('putData = ', roleModel.value.id, putData);
        let result = await updateRoleService(roleModel.value.id, putData);
        ElMessage.success(result.msg ? result.msg : '更新成功') // 显示返回信息
    }
    // 获取所有角色
    roleList();
    // 关闭窗口
    dialogVisible.value = false;
}
// 删除角色
const deleteRole = (row) => {
    //提示用户  确认框
    ElMessageBox.confirm('你确认要删除该权限信息吗?', '温馨提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(async () => { //调用接口
        let result = await deleteRoleService(row.id);
        ElMessage({ type: 'success', message: '删除成功'})
            //刷新列表
        roleList();
    }).catch(() => {
        ElMessage({ type: 'info', message: '管理员取消了删除'})
    })
}
</script>
<template>
    <el-card class="page-container">
        <template #header>
            <div class="header">
                <span>角色分类</span>
                <div class="extra">
                    <el-button type="primary" @click="showDialog(null, 0)">添加角色</el-button>
                </div>
            </div>
        </template>
        <el-table :data="roles" style="width: 100%">
            <el-table-column label="序号" width="100" type="index"> </el-table-column>
            <el-table-column label="角色名称" prop="rolename" align="center"></el-table-column>
            <el-table-column label="操作" width="200">
                <template #default="{ row }">
                    <!-- 查看角色权限对话框 -->
                    <!-- 步骤：显示对话框 传值  -->
                    <el-button :icon="View" circle plain type="info" @click="showDialog(row, 1)"></el-button>
                    <el-button :icon="Edit" circle plain type="primary" @click="showDialog(row, 2)"></el-button>
                    <el-button :icon="Delete" circle plain type="danger" @click="deleteRole(row)"></el-button>
                </template>
            </el-table-column>
            <template #empty>
                <el-empty description="没有数据" />
            </template>
        </el-table>

        <!-- 添加新的角色弹窗 -->
        <el-dialog v-model="dialogVisible" :title="title" width="30%">
            <el-form :model="roleModel" :rules="rules" label-width="100px" style="padding-right: 30px">
                <el-form-item label="角色名称" prop="rolename">
                    <el-input 
                    v-model="roleModel.rolename" 
                    minlength="1" 
                    maxlength="10"
                    :disabled="isInputDisabled"
                    ></el-input>
                </el-form-item>
                <el-form-item label="角色权限">
                    <el-select v-model="selectedAccess" @change="handleAccessChange" placeholder="Select">
                        <el-option 
                        v-for="access in accesses" 
                        :key="access.id" 
                        :label="access.title" 
                        :value="access"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="权限选择">
                    <el-checkbox-group v-model="selectedOption" size="large" @change="handleOption" >
                        <el-checkbox v-for="option in options" 
                        :key="option" 
                        :label="option"
                        :disabled="isDisabled"
                        >
                        {{ option }}
                        </el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="选择子角色">
                    <el-select
                        v-model="selectedRole"
                        multiple
                        placeholder="Select"
                        >
                        <el-option
                            v-for="role in roles"
                            :key="role.id"
                            :label="role.rolename"
                            :value="role.id"/>
                        </el-select>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button 
                    v-if="editVisible"
                    type="primary" 
                    @click="addRole"
                    round
                    > 提交 </el-button>
                    <el-button 
                    type="info" 
                    @click="dialogVisible = false"
                    round
                    >取消</el-button>
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
</style>

