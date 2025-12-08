<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue';
import TodoList from '@/components/TodoList.vue'
</script>

<template>
  <AppHeader />
  <main class="flex h-full w-full">
    <TodoList class="m-auto mt-12" />
  </main>
</template>
