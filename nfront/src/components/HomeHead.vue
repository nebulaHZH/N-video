<template>
 <el-menu 
    mode="horizontal"
    :ellipsis="false"
    class="head-content"
    >
    <el-menu-item index="0">
        <el-image :src="url"/>
    </el-menu-item>
    <el-menu-item class="search">
        <el-input :suffix-icon="Search" size="large" />
    </el-menu-item>
    <div class="flex-grow" />
    <el-menu-item>首页</el-menu-item>
    <el-menu-item>电影</el-menu-item>
    <el-menu-item>电视剧</el-menu-item>
    <el-menu-item>综艺</el-menu-item>
    <el-menu-item>动漫</el-menu-item>
    <el-menu-item>纪录片</el-menu-item>
    <el-menu-item>排行榜</el-menu-item>
    <el-menu-item>历史</el-menu-item>
 </el-menu>
</template>

<script setup>
 import { Search } from '@element-plus/icons-vue'
 import { ref } from 'vue'
 const url = ref('../assets/logo.png')
</script>

<style scoped lang="less">
.el-menu--horizontal .el-menu-item:not(.is-disabled):hover{
    background-color: #ffffff !important;
}
.head-content{
    width: 100%;
}
.search{
    width: 50vw;
}
.flex-grow{
  flex-grow: 1;
}
</style>