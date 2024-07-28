<template>
 <el-row class="waterfall-content" :gutter="20">
    <el-col v-for="(item,index) in list" :key="index" :span="4">
        <el-card style="margin-top: 10px;width: fit-content;" class="waterfall-item" shadow="hover" @click="toVideoDetail(item.vid)">
            <el-image :src="item.url" fit="cover"/>
            <div class="waterfall-item-title">{{item.title}}</div>
        </el-card>
        
    </el-col>
</el-row>
</template>

<script setup>
import { ElCard } from 'element-plus'
import { ref } from 'vue'
const list = ref([
    {
        vid:1,
        title:'决斗场',
        url:'https://images.weserv.nl/?url='+'https://647bc185.szrtcpa.com/2024/01/02/3775e1389c52c.jpg'
    },
    {
        vid:1,
        title:'决斗场',
        url:'https://images.weserv.nl/?url='+'https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2909997626.jpg'
    },
    {
        vid:1,
        title:'决斗场',
        url:'https://images.weserv.nl/?url='+'https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2909997626.jpg'
    },
    {
        vid:1,
        title:'决斗场',
        url:'https://images.weserv.nl/?url='+'https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2909997626.jpg'
    },
    {
        vid:1,
        title:'决斗场',
        url:'https://images.weserv.nl/?url='+'https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2909997626.jpg'
    },
    {
        vid:1,
        title:'决斗场',
        url:'https://images.weserv.nl/?url='+'https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2909997626.jpg'
    },
    {
        vid:1,
        title:'决斗场',
        url:'https://images.weserv.nl/?url='+'https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2909997626.jpg'
    }
])
</script>

<style scoped>
.waterfall-item:hover{
    cursor: pointer;
    transform: scale(1.1);
    transition: all 0.3s ease-in-out;
}
.waterfall-item-title{
    text-align: left;
}
</style>