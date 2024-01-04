<template>
    <div class="wrapper">
      <div class="navigation">
        <Navigation />
      </div>
      <br>
      <div class="search">
        <SearchUserByName />
      </div>
    </div>
</template>
<script setup>
import Navigation from "@/view/home/Navigation.vue";
import SearchUserByName from "@/view/home/SearchUserByName.vue";
</script>
<style scoped>

</style>