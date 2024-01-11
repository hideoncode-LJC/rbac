import request from '@/utils/request.js'

export const getAccessByRoleIdService = (id) => {
    return request.get("/roleaccess?id="+id)
}

export const accessListService = (roleid, accessid) => {
    return request.get("/access?roleid=" + roleid + "&" + "accessid=" + accessid);
}

export const accessByIDService = (id) => {
    return request.get("/accessbyid?accessid=" + id)
}

export const addAccessService = (data) => {
    return request.post("/access", data)
}

export const updateAccessService = (data) => {
    return request.put("/access?id="+ data.id, data)
}

export const deleteAccessService = (id) => {
    return request.delete("/access?id="+id)
}