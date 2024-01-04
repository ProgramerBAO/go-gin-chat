<script setup>
import {useUserStore} from "@/stores/modules/user";

const userStore = useUserStore()
</script>

<template>
  <h1>hello {{ userStore.userInfo.userName }}</h1>
</template>

<style scoped>

</style>