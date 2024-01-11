<script setup>

</script>

<template>
    <el-card class="page-container">
        <template #header><div class="header"><span>RBAC demo 主页</span></div></template>
        <p>
            基于角色的访问控制（Role-based Access Control，RBAC）是围绕角色和权限定义的一种政策中立的访问控制机制。RBAC 的组成部分，如角色-权限、用户-角色和角色-角色的关系，使用户分配变得简单。NIST（美国国家标准与技术研究院）的一项研究表明，RBAC 满足商业和政府组织的许多需求。RBAC 可以用于在拥有数百用户和数千个权限的大型组织中简化安全管理。尽管 RBAC 与强制访问控制（MAC）和自主访问控制（DAC）框架不同，但它可以在不引起复杂性的情况下执行这些策略。

        </p>
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