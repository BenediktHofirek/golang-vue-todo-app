<script setup lang="ts">
import { onMounted, shallowRef } from 'vue';
import AddTask from '@material-symbols/svg-600/outlined/add_task.svg'
import Todo from './Todo.vue';

const userId = 'a1b2c3d4-e5f6-7890-1234-567890abcdef'

let todoList = shallowRef();

onMounted(() => {
  fetch(`/api/v1/todos/${userId}`).then(async (response) => {
    const todos = await response.json();
    todoList.value = todos
  })
})

function addTodo() {
  console.log('addingTodo')
}

</script>

<template>
  <div class="bg-white w-full max-w-2xl rounded-xl border border-gray-100 hover:shadow-md px-2 py-4">
    <h4 class="text-xl font-medium mb-2 pl-2">Task List</h4>

    <button @click="addTodo()" class="w-full font-semibold text-blue-600 flex items-center transition-colors justify-start gap-4
      rounded-full hover:bg-blue-50 p-2 hover:cursor-pointer mb-4">
      <AddTask class="size-6 fill-blue-600" />
      <span>Add a task</span>
    </button>

    <ul v-if="todoList" class="px-2">
      <Todo v-for="todo in todoList" :key="todo.id" :todo />
    </ul>
    <div v-else>Fetching todo list...</div>
  </div>
</template>
