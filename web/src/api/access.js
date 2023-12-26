import request from '@/utils/request.js'

export const getAccessByRoleIdService = (id) => {
    return request.get("/roleaccess?id="+id)
}

export const accessesListService = () => {
    return request.get("/access")
}

export const addAccessService = (data) => {
    // const requestData = {
    //     "title": data.title,
    //     "content": data.content,
    // }
    return request.post("/access", data)
}

export const updateAccessService = (data) => {
    return request.put("/access", data)
}

export const deleteAccessService = (id) => {
    return request.delete("/access?id="+id)
}