<script setup lang="ts">
import { onMounted, shallowRef } from 'vue';

const userId = 'a1b2c3d4-e5f6-7890-1234-567890abcdef'

let todoList = shallowRef();

onMounted(() => {
  fetch(`/api/v1/todos/${userId}`).then(async (response) => {
    const todos = await response.json();
    todoList.value = todos
  })
})

</script>

<template>
  <div class="bg-white w-full max-w-2xl rounded-xl border border-gray-100 hover:shadow-md p-4">
    <ul v-if="todoList">
      <li v-for="todo in todoList" :key="todo.id">
        {{ todo.title }}
      </li>
    </ul>
    <div v-else>Fetching todo list...</div>
  </div>
</template>
