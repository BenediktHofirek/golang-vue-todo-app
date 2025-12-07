<script setup lang="ts">
import { VueQueryDevtools } from '@tanstack/vue-query-devtools'
import { RouterView } from 'vue-router'
</script>

<template>
  <router-view></router-view>
  <VueQueryDevtools />
</template>

<style scoped></style>
