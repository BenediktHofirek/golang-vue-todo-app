<script setup lang="ts">
import AddTask from '@material-symbols/svg-600/outlined/add_task.svg'
import TodoItem from './TodoItem.vue'
import { useQuery } from '@tanstack/vue-query'
import { apiClient } from '@/apiClient'

const userId = 'a1b2c3d4-e5f6-7890-1234-567890abcdef'

const {
  isPending,
  isError,
  data: todoList,
} = useQuery({
  queryKey: ['todos'],
  queryFn: ({ signal }) => {
    return apiClient.get(`/todos/${userId}`, {
      signal,
    })
  },
})

function addTodo() {
  console.log('addingTodo')
}
</script>

<template>
  <div
    class="
      w-full max-w-2xl rounded-xl border border-gray-100 bg-white
      hover:shadow-md
    "
  >
    <h4 class="p-4 text-xl font-medium">Task List</h4>

    <button
      @click="addTodo()"
      class="
        m-auto mb-4 flex w-[calc(100%-1rem)] items-center justify-start gap-4 rounded-full p-2
        font-semibold text-blue-600 transition-colors
        hover:cursor-pointer hover:bg-blue-50
        active:bg-blue-100
      "
    >
      <AddTask class="size-6 fill-blue-600" />
      <span>Add a task</span>
    </button>

    <div v-if="isPending">Todo list is pending...</div>
    <div v-else-if="isError">
      Oh no, an error has ocurred during fetch of todo list!
    </div>
    <ul v-else-if="todoList">
      <TodoItem v-for="todo in todoList" :key="todo.id" :todo />
    </ul>
  </div>
</template>
