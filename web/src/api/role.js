import request from '@/utils/request.js'

export const roleListService = () => {
    return request.get("/role");
}

export const addRoleService = (data) => {
    return request.post("/role", data);
}

export const updateRoleService = (id, data) => {
    return request.put("/role?id=" + id, data);
}

export const deleteRoleService = (id) => {
    return request.delete("/role?id=" + id);
}

export const childRoleListSerivice = (id) => {
    return request.get("/child?id=" + id);
}