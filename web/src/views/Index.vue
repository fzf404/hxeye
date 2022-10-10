<template>
  <div>
    <navbar />
    <home class="show-page"/>
    <footbar />
  </div>
</template>

<script>
import Home from './Home'
import Navbar from '@/components/Navbar.vue'
import Footbar from '@/components/Footbar.vue'

export default {
  components: { Navbar, Footbar, Home }
}

</script>

<style scoped lang="scss">
.show-page{
  margin-top: 180px;
}
</style>
