<script setup lang="ts">
import { onMounted, ref } from 'vue';

const userId = 'a1b2c3d4-e5f6-7890-1234-567890abcdef'

let todoList = ref();

onMounted(() => {
  fetch(`http://localhost:8080/api/v1/todos/${userId}`).then((todos) => {
    todoList.value = todos
  })
})

</script>

<template>
  <div class="w-20 h-20 bg-green-400">
    <ul v-if="todoList">
      <li v-for="todo in todoList" :key="todo.id">
        {{ todo.title }}
      </li>
    </ul>
    <div v-else>Fetching todo list...</div>
  </div>
</template>
