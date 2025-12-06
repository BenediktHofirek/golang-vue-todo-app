<script setup lang="ts">
import { VueQueryDevtools } from '@tanstack/vue-query-devtools'
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header class="mb-12 flex">
    <nav>
      <RouterLink to="/">Home</RouterLink>
    </nav>
  </header>

  <RouterView />
  <VueQueryDevtools />
</template>

<style scoped></style>
