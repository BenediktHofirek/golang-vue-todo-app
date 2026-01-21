<script setup lang="ts">
import AddTask from '@material-symbols/svg-600/outlined/add_task.svg'
import ArrowRight from '@material-symbols/svg-600/outlined/arrow_right.svg'
import TodoItem from './TodoItem.vue'
import {
  useMutation,
  useMutationState,
  useQuery,
} from '@tanstack/vue-query'
import { computed, ref } from 'vue'
import type { Todo } from '@/models'
import { NEW_TODO_TEMP_ID } from '@/constants'
import { isTodoBlank } from '@/utils'
import { api } from '@/apiClient'

const {
  isPending,
  isError,
  isSuccess,
  data: todoList,
} = useQuery({
  queryKey: ['todos'],
  queryFn: api.getTodoList,
})

const editedTodoId = ref<number | null>(null)

const updatedTodoList = useMutationState({
  filters: { mutationKey: ['updateTodo'], status: 'pending' },
  select: (mutation) => mutation.state.variables as Todo,
})

const deletedTodoList = useMutationState({
  filters: { mutationKey: ['deleteTodo'], status: 'pending' },
  select: (mutation) => mutation.state.variables as number,
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

const { mutate: saveNewTodo } = useMutation({
  mutationFn: api.updateTodo,
  mutationKey: ['updateTodo'],
})

const { mutate: addTodo } = useMutation({
  mutationFn: api.createBlankTodo,
  onMutate: async (_, context) => {
    await context.client.cancelQueries({ queryKey: ['todos'] })
    const existingBlankTodo = (
      context.client.getQueryData(['todos']) as Todo[]
    ).find((todo) => isTodoBlank(todo))

    if (existingBlankTodo) return
    const blankTodo = {
      id: NEW_TODO_TEMP_ID,
      title: null,
      description: null,
      completed: false,
      dueDate: null,
      starred: false,
      scheduledDate: null,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    }
    await context.client.setQueryData(['todos'], (old: Todo[]) =>
      old.concat(blankTodo),
    )
    editedTodoId.value = NEW_TODO_TEMP_ID
  },
  onSuccess: async (newTodoId, _variables, _onMutateResult, context) => {
    await context.client.setQueryData(['todos'], (old: Todo[]) =>
      old.map((todo) => {
        if (todo.id !== NEW_TODO_TEMP_ID) return todo
        return {
          ...todo,
          id: newTodoId,
        }
      }),
    )

    const newTodo = (context.client.getQueryData(['todos']) as Todo[]).find(
      (todo) => {
        return todo.id === newTodoId
      },
    )
    if (!newTodo) return
    saveNewTodo(newTodo)
  },
})

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
      <TodoItem
        v-for="todo in openTodoList"
        :key="todo.id"
        :todo
        :isEdited="editedTodoId === todo.id"
      />
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
      <TodoItem
        v-for="todo in closedTodoList"
        :key="todo.id"
        :todo
        :isEdited="editedTodoId === todo.id"
      />
    </ul>
  </div>
</template>
