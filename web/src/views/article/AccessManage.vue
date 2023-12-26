<script setup>
import { Edit, Delete } from '@element-plus/icons-vue'; // 导入图标 用作编辑框、删除框的图标显示
import {ElMessage} from 'element-plus'; // 消息提示框
import { QuillEditor } from '@vueup/vue-quill'; // content内容编写框导入
import '@vueup/vue-quill/dist/vue-quill.snow.css'; // content 内容编写框导入样式
import { ElMessageBox } from 'element-plus'; // 导入删除框 根据选择进行提示
import { onMounted, ref} from 'vue'; // 导入响应式数据和挂载函数
import { accessesListService, addAccessService, deleteAccessService } from '@/api/access.js' // 导入函数
const title = ref(''); // 设置添加权限、更改权限的标题显示
const submitTitle = ref(''); // 设置添加权限、更改权限的按钮显示
const accesses = ref([]); // 定义接收权限的数据
const drawerVisible = ref(false); // 设置添加权限、更改权限的抽屉显示
const accessModel = ref({
    "id": '',
    "title": '',
    "content":'',
}) // 定义添加权限、修改权限的数据模型
const accessRules = {
    title: [
        { required: true, message: '请输入文章标题', trigger: 'blur' },
    ],
} // 添加权限时的数据校验规则

// 清空数据模型
const clearAccessModel = () => {
    // 清空 addAccessModel 数据
    accessModel.value.id = ''
    accessModel.value.title = '';
    accessModel.value.content = '';
}
// 清空权限数据函数
const clearAccesses = () => {
    accesses.value = [];
}
// 挂在函数 获取收据
onMounted(() => {
  accessesList();
});
// 控制抽屉的显示
const showDrawer = (drawerTitle, row) => {
    if (drawerTitle === '添加权限') {
        clearAccessModel();
        submitTitle.value = '添加';
    } else {
        accessModel.value.id = row.id;
        accessModel.value.title = row.title;
        accessModel.value.content = row.content;
        submitTitle.value = '修改';
    } 
    drawerVisible.value = true;
    title.value = drawerTitle;
    console.log("accessModel = ", accessModel.value)
}
// 获取权限
const accessesList = async() => {
    clearAccesses(); // 先清空数据
    let result = await accessesListService(); // 获取权限数据
    for (let i = 0 ; i < result.data.length ; i ++ ) {
        let access = {
            "id": result.data[i].ID,
            "title": result.data[i].title,
            "content": result.data[i].content,
        }
        accesses.value.push(access);
    }
}
// 添加权限
const addAccess = async() => {
    let result = await addAccessService(accessModel.value);
    // 显示返回信息
    ElMessage.success(result.msg ? result.msg : '添加权限成功');
    // 获取所有角色
    if (result.code === 0) {
        accessesList();
        // 关闭窗口
        drawerVisible.value = false;
    } 
}
// 更新权限
const updateAccess = async() => {
    let result = await updateAccessService(accessModel);
    // 显示返回信息
    ElMessage.success(result.msg ? result.msg : '修改权限成功');
    // 获取所有角色
    accessesList();
    // 关闭窗口
    drawerVisible.value = false;
}
// 删除权限
const deleteAccess = (row) => {
    //提示用户  确认框
    ElMessageBox.confirm(
        '你确认要删除该权限信息吗?',
        '温馨提示',
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
        .then(async () => {
            //调用接口
            let result = await deleteAccessService(row.id);
            ElMessage({
                type: 'success',
                message: '删除成功',
            })
            //刷新列表
            accessesList();
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '管理员取消了删除',
            })
        })
}
</script>

<template>
    <el-card class="page-container">

        <!-- header显示 -->
        <template #header>
            <div class="header">
                <span>权限管理</span>
                <div class="extra">
                    <el-button type="primary" @click="showDrawer('添加权限')">添加权限</el-button>
                </div>
            </div>
        </template>

        <!-- 权限列表 -->
        <el-table :data="accesses" style="width: 100%">
            <el-table-column label="序号" width="100" type="index"></el-table-column>
            <el-table-column label="标题"  width="300" prop="title" align="center"></el-table-column>
            <el-table-column label="内容"  prop="content"></el-table-column>
            <el-table-column label="操作" width="100">
                <template #default="{ row }">
                    <el-button :icon="Edit" circle plain type="primary" @click="showDrawer('修改权限', row)"></el-button>
                    <el-button :icon="Delete" circle plain type="danger" @click="deleteAccess(row)"></el-button>
                </template>
            </el-table-column>
            <template #empty>
                <el-empty description="没有数据" />
            </template>
        </el-table>

        <!-- 添加权限抽屉 -->
        <el-drawer v-model="drawerVisible" :title="title" direction="rtl" size="50%">
            <!-- 添加权限表单 -->
            <el-form :model="accessModel" :rules="accessRules" label-width="100px" >
                <el-form-item label="权限标题" prop="title">
                    <el-input v-model="accessModel.title" placeholder="请输入标题"></el-input>
                </el-form-item>
                <el-form-item label="文章内容">
                    <div class="editor">
                        <quill-editor 
                        v-model:content="accessModel.content" 
                        placeholder="请输入权限内容"
                        contentType="html">
                        </quill-editor>
                    </div>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="title === '添加权限' ? addAccess() : updateAccess()">
                    {{ submitTitle }}
                    </el-button>
                    <el-button type="info" @click="drawerVisible = false">取消</el-button>
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