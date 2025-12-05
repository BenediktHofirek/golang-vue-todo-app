<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header class="flex mb-12">
      <nav>
        <RouterLink to="/">Home</RouterLink>
      </nav>
  </header>

  <RouterView />
</template>

<style scoped>

</style>
