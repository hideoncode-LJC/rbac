import {defineStore} from 'pinia'
import {ref} from 'vue'

export const useAdminStore = defineStore('adminToken',()=>{

    const adminToken = ref({}); //定义状态相关的内容

    const setToken = (newAdmin)=>{
        adminToken.value = {};
        adminToken.value = newAdmin;
    }

    const removeToken = ()=>{
        adminToken.value = {};
    }

    return {adminToken,setToken,removeToken};
},{persist:true})
