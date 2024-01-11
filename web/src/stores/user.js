import {defineStore} from 'pinia'

import {ref} from 'vue'

export const useUserStore = defineStore('userToken',()=>{

    const userToken = ref({}); //定义状态相关的内容

    const setToken = (newUser)=>{
        userToken.value = {};
        userToken.value = newUser;
    }
    const removeToken = ()=>{
        userToken.value = {};
    }
    return {userToken,setToken,removeToken};
},{persist:true})

