import request from '@/utils/request.js'

export const childLevelService = (accessid, roleid) => {
    return request.get('/level?accessid=' + accessid + '&roleid=' + roleid);
}