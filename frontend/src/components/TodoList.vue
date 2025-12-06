<script setup lang="ts">
import { onMounted, shallowRef } from 'vue'
import AddTask from '@material-symbols/svg-600/outlined/add_task.svg'
import TodoItem from './TodoItem.vue'

const userId = 'a1b2c3d4-e5f6-7890-1234-567890abcdef'

const todoList = shallowRef()

onMounted(() => {
  fetch(`/api/v1/todos/${userId}`).then(async (response) => {
    const todos = await response.json()
    todoList.value = todos
  })
})

function addTodo() {
  console.log('addingTodo')
}
</script>

<template>
  <div
    class="
      w-full max-w-2xl rounded-xl border border-gray-100 bg-white px-2 py-4
      hover:shadow-md
    "
  >
    <h4 class="mb-2 pl-2 text-xl font-medium">Task List</h4>

    <button
      @click="addTodo()"
      class="
        mb-4 flex w-full items-center justify-start gap-4 rounded-full p-2 font-semibold
        text-blue-600 transition-colors
        hover:cursor-pointer hover:bg-blue-50
      "
    >
      <AddTask class="size-6 fill-blue-600" />
      <span>Add a task</span>
    </button>

    <ul v-if="todoList" class="px-2">
      <TodoItem v-for="todo in todoList" :key="todo.id" :todo />
    </ul>
    <div v-else>Fetching todo list...</div>
  </div>
</template>
