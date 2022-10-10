<template>
  <div>
    <navbar />
    <div class="show-page404">
      <b-alert show variant="danger" >您所访问的内容不存在</b-alert>
    </div>
    <footbar />
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import Footbar from '@/components/Footbar.vue'

export default {
  components: { Navbar, Footbar }
}
</script>

<style scoped lang="scss">
.show-page404{
  margin: 280px auto;
  width: 400px;
}
</style>
