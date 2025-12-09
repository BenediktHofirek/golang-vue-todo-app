<script setup lang="ts">
import AddTask from '@material-symbols/svg-600/outlined/add_task.svg'
import ArrowRight from '@material-symbols/svg-600/outlined/arrow_right.svg'
import TodoItem from './TodoItem.vue'
import { useMutationState, useQuery } from '@tanstack/vue-query'
import { apiClient } from '@/apiClient'
import { computed, ref } from 'vue'
import type { Todo } from '@/models'

const {
  isPending,
  isError,
  isSuccess,
  data: todoList,
} = useQuery({
  queryKey: ['todos'],
  queryFn: ({ signal }) => {
    return apiClient.get('/todos', {
      signal,
    }) as Promise<Todo[]>
  },
})

const updatedTodoList = useMutationState<Todo>({
  filters: { mutationKey: ['updateTodo'], status: 'pending' },
  select: (mutation) => mutation.state.variables as any,
})

const deletedTodoList = useMutationState<number>({
  filters: { mutationKey: ['deleteTodo'], status: 'pending' },
  select: (mutation) => mutation.state.variables as any,
})

const todoListWithUpdates = computed(() => {
  if (!todoList.value) return []

  const updatedTodoMap = updatedTodoList.value.reduce(
    (acc, todo) => ({
      ...acc,
      [todo.id]: todo,
    }),
    {} as Record<number, Todo>,
  )

  return todoList.value
    .filter((todo) => {
      return !deletedTodoList.value.includes(todo.id)
    })
    .map((todo) => ({
      ...todo,
      ...(updatedTodoMap[todo.id] || {}),
    }))
})

function addTodo() {
  console.log('addingTodo')
}

const isClosedTodosVisible = ref(false)

const openTodoList = computed(() => {
  if (!todoList.value) return []

  return todoListWithUpdates.value.filter((todo) => !todo.completed)
})

const closedTodoList = computed(() => {
  if (!todoList.value) return []

  return todoListWithUpdates.value.filter((todo) => todo.completed)
})
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
    <ul v-else-if="isSuccess">
      <TodoItem v-for="todo in openTodoList" :key="todo.id" :todo />
    </ul>

    <div
      @click="isClosedTodosVisible = !isClosedTodosVisible"
      v-if="closedTodoList.length"
      class="flex w-full cursor-pointer items-center justify-start gap-2 p-2"
    >
      <ArrowRight
        class="
          size-10 rounded-full p-2 transition-colors
          hover:bg-gray-100
          active:bg-gray-200
        "
        :class="isClosedTodosVisible ? ['rotate-90'] : []"
      />
      <div class="h-min pt-0.5 text-base font-medium select-none">
        Completed: ({{ closedTodoList.length }})
      </div>
    </div>

    <ul v-if="closedTodoList.length && isClosedTodosVisible">
      <TodoItem v-for="todo in closedTodoList" :key="todo.id" :todo />
    </ul>
  </div>
</template>
