<script setup>
import { 
    Edit,
    Delete,
} from '@element-plus/icons-vue';
import { routerKey, useRoute } from 'vue-router';
import { useTokenStore } from '@/stores/token.js'
import { useAccessStore } from '@/stores/access.js'
import { onMounted, ref, watch, watchEffect} from 'vue';
import { updateAccessService, deleteAccessService, accessListService} from '@/api/access';
import { ElMessage, ElMessageBox, roleTypes} from 'element-plus'
import { useAdminStore } from '@/stores/admin.js';
import { useUserStore } from '@/stores/user.js'
// 获取用户权限Token
const accessStore = useAccessStore();
const adminStore = useAdminStore();
const accessToken = accessStore.accessToken;
const userStore = useUserStore();

// 通过路由参数获取当前页面的参数方便获取该界面的信息
const route = useRoute(); 
// 存储该界面的参数
const page = ref('')
const level = ref('');
// 控制界面组件的可编辑性质
const isContentEditAble = ref(false);
const isEditButtonAble = ref(false);
const isDeleteButtonAble = ref(false);

// 存储界面上显示的access数据
const access = ref({
    "id": '',
    "title": '',
    "content": '',
});
// 清空access 数据
const clear = () => {
    isContentEditAble.value = false;
    isEditButtonAble.value = false;
    isDeleteButtonAble.value = false;
}

import { childLevelService } from '@/api/level.js'

const initAccess = async() => {
    let result = await childLevelService(page.value, userStore.userToken.roleid);
    level.value = result.data;
    if (level.value > 0 ) {
        clear();
        if(level.value === 3) {
            isDeleteButtonAble.value = true;
            isEditButtonAble.value = true;
            isContentEditAble.value = true;
        } else if (level.value === 2) {
            isEditButtonAble.value = true;
            isContentEditAble.value = true;
        }
        
    } else {
        getAccess
    }
}

const getAccess = () => {
    clear();
    if (adminStore.adminToken.admin) {
        isDeleteButtonAble.value = true;
        isEditButtonAble.value = true;
        isContentEditAble.value = true;
        return;
    }
    const accessToken = accessStore.accessToken;
    for (let i = 0 ; i < accessToken.length ; i ++ ) {
        if (page.value === String(accessToken[i].ID)) {
            if (accessToken[i].level === 3) {
                isDeleteButtonAble.value = true;
                isEditButtonAble.value = true;
                isContentEditAble.value = true;
            } else if (accessToken[i].level === 2) {
                isEditButtonAble.value = true;
                isContentEditAble.value = true;
            }
            return 
        }
    }
    console.log("Access.vue", "当前界面要显示的内容为", access.value)
}

const initAccessContent = async(id) => {
    let result = await accessListService(0, id);
    access.value.id = result.data.ID;
    access.value.title = result.data.title;
    access.value.content = result.data.content;
}

// console.log('@@route = ', route);
    // level.value = route.query.level;
    // console.log('子角色level', level.value);

console.log('@', route);
console.log('@', route.fullPath, route.params, route.path, route.query);

onMounted(() => {
    page.value = route.params.page;
    initAccessContent(route.params.page);
    initAccess();
})

watch(() => route.params.page, (newValue, oldValue) => {
    // console.log("newValue = ", newValue)
    // console.log("oldValue = ", oldValue)
    // console.log('@@route = ', route);
    // console.log('@', route.query);
    if (typeof newValue === 'undefined') {
        return;
    }
    page.value = newValue
    initAccessContent(route.params.page);
    initAccess();
    getAccess();
});

const updateAccess = async() => {
    let result = await updateAccessService(access.value);
    ElMessage.success(result.msg ? result.msg : '更新成功'); 
}

const deleteAccess = () => {
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
            let result = await deleteAccessService(access.value.id);
            console.log('删除权限的返回信息', result.data);
            ElMessage({
                type: 'success',
                message: '删除成功',
            })
            //刷新列表
            accessList(0, 0);
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '用户取消了删除',
            })
        })
}



</script>

<template>
    <el-card class="page-container">
        <template #header>
            <div class="header">
                <span>权限查看</span>
            </div>
        </template>

        <el-row class="center-row">
            <el-col :span="20">
                <el-form :model="access"  label-width="100px" size="large">
                    <el-form-item label="标题">
                        <el-input v-model="access.title" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="内容" prop="nickname">
                        <el-input 
                        v-model="access.content" 
                        :disabled="!isContentEditAble"
                        type="textarea" 
                        ></el-input>
                    </el-form-item>
                    <el-form-item >
                        <el-button 
                        type="primary" 
                        :icon="Edit" 
                        :disabled="!isEditButtonAble"
                        round
                        @click="updateAccess"
                        >修改</el-button>
                        <el-button 
                        type="danger" 
                        :icon="Delete" 
                        :disabled="!isDeleteButtonAble"
                        round
                        @click="deleteAccess"
                        >删除</el-button>
                    </el-form-item>
                </el-form>
            </el-col>
        </el-row>
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

.center-row {
    display: flex;
    justify-content: center;
    align-items: center;
     // You can adjust this based on your layout needs
}

</style>