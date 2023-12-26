<script setup>
import { 
    Edit,
    Delete,
} from '@element-plus/icons-vue';
import { useRoute } from 'vue-router';
import {useTokenStore} from '@/stores/token.js'
import { onMounted, ref, watch} from 'vue';
const tokenStore = useTokenStore();
const route = useRoute();
// 通过路由参数获取当前页面
const page = ref('')
const isContentEditAble = ref(false);
const isEditButtonAble = ref(false);
const isDeleteButtonAble = ref(false);

const access = ref({
    "id": '',
    "title": '',
    "content": '',
});

const getAccess = () => {
    const token = tokenStore.token
    for (let i = 0 ; i < token.access.length ; i ++ ) {
        if (page.value === String(token.access[i].Access.ID)) {
            access.value.id = token.access[i].Access.ID;
            access.value.title = token.access[i].Access.title;
            console.log("title = ", token.access[i].Access.title)
            access.value.content = token.access[i].Access.content;
            console.log("content = ", token.access[i].Access.content)
            return 
        }
    }
    console.log("access = ", access.value)
}


onMounted(() => {
    page.value = route.params.page;
    getAccess();
})

watch(() => route.params.page, (newValue, oldValue) => {
    console.log("newValue = ", newValue)
    console.log("oldValue = ", oldValue)
    page.value = newValue
    getAccess();
});



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
                        >修改</el-button>
                        <el-button 
                        type="danger" 
                        :icon="Delete" 
                        :disabled="!isDeleteButtonAble"
                        round
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