<template>
  <div class="container">
    <el-container>
      <el-header>
        <HomeHead />
      </el-header>
      <el-main class="scrollable-content">
        <router-view/>
      </el-main>
    </el-container>
    <el-button class="footer" size="large" type="primary" circle="true" @click="toMyPage">我的
    </el-button>
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router'
import { ElContainer, ElHeader, ElMain,ElButton } from 'element-plus'
import HomeHead from '../components/HomeHead.vue'
import {useRouter} from 'vue-router'
const r= useRouter()
const toMyPage = () =>{
  r.push('/my')
}
</script>

<style scoped>
.scrollable-content{
  height: 95vh;
  width:90%;
  margin-left:5%;
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