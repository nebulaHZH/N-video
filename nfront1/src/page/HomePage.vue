<template>
  <div class="container">
    <el-container>
      <el-header>
        <HomeHead />
      </el-header>
      <el-main class="scrollable-content">
        <VideoList />
      </el-main>
    </el-container>
    <el-button class="footer" size="large" type="primary" circle="true" @click="toMyPage">我的</el-button>
  </div>
</template>

<script setup>
import { ElContainer, ElHeader, ElMain, ElFooter } from 'element-plus'
import VideoList from '../components/VideoList.vue'
import HomeHead from '../components/HomeHead.vue'
const toMyPage = () =>{
  window.location.href = './my'
}
</script>

<style scoped>
.scrollable-content{
  height: 95vh;
/* 隐藏浏览器默认的滚动条 */
	-ms-overflow-style: none;
	/* IE 10+ */
	overflow: -moz-scrollbars-none;
	/* Firefox */
	scrollbar-width: none;
	/*firefox*/
}
.footer{
  position: fixed;
  bottom: 5vh;
  right: 3vw;
}

.container{
  position: fixed;
  width: 99.5%;
  overflow: hidden;
  height: 100%; 
}
el-header{
  height: 20vh;
}
</style>