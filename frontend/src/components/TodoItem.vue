<script setup lang="ts">
import type { Todo } from '@/models'
import RadioButtonUnchecked from '@material-symbols/svg-600/rounded/radio_button_unchecked.svg'
import Delete from '@material-symbols/svg-600/rounded/delete.svg'
import MoreVert from '@material-symbols/svg-700/rounded/more_vert.svg'
import Star from '@material-symbols/svg-600/rounded/star.svg'
import StarFilled from '@material-symbols/svg-600/rounded/star-fill.svg'
import Check from '@material-symbols/svg-600/rounded/check.svg'
import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { apiClient } from '@/apiClient'

defineProps<{
  todo: Todo
}>()

const queryClient = useQueryClient()
const { mutate: updateTodo } = useMutation({
  mutationFn: (updatedTodo: Partial<Todo>) => {
    return apiClient.patch(`/todos/${updatedTodo.id}`, updatedTodo)
  },
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['todos'] }),
  mutationKey: ['updateTodo'],
})

const { mutate: deleteTodo } = useMutation({
  mutationFn: (todoId: number) => {
    return apiClient.delete(`/todos/${todoId}`)
  },
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['todos'] }),
  mutationKey: ['deleteTodo'],
})

async function toggleCompleted(todo: Todo) {
  updateTodo({
    id: todo.id,
    completed: !todo.completed,
  })
}

async function toggleStarred(todo: Todo) {
  updateTodo({
    id: todo.id,
    starred: !todo.starred,
  })
}
</script>

<template>
  <div class="
    group w-full py-2
    hover:bg-blue-50
  ">
    <div class="flex cursor-default justify-start px-4 transition-colors">
      <div
        class="mt-1 mr-4 size-6 cursor-pointer overflow-hidden rounded-full"
        @click.once="toggleCompleted(todo)"
      >
        <Check
          v-if="todo.completed"
          class="
            block size-6 fill-blue-600
            hover:bg-gray-200
          "
        />
        <div v-else class="group/check flex size-6 items-center justify-center">
          <RadioButtonUnchecked class="
            size-5
            group-hover/check:hidden
          " />
          <Check
            class="
              hidden size-6 bg-gray-200 fill-blue-600
              group-hover/check:block
            "
          />
        </div>
      </div>
      <div class="flex w-full flex-col">
        <div class="flex w-full items-center justify-between">
          <div class="text-base" :class="todo.completed ? 'line-through' : ''">
            {{ todo.title }}
          </div>

          <Delete
            v-if="todo.completed"
            class="
              invisible size-9 cursor-pointer rounded-full p-2 transition-colors
              group-hover:visible
              hover:bg-gray-100
              active:bg-gray-200
            "
            @click.once="deleteTodo(todo.id)"
          />

          <div v-else class="flex items-center">
            <MoreVert
              class="
                invisible size-9 cursor-pointer rounded-full p-2 transition-colors
                group-hover:visible
                hover:bg-gray-100
                active:bg-gray-200
              "
            />
            <StarFilled
              v-if="todo.starred"
              @click="toggleStarred(todo)"
              class="
                size-9 cursor-pointer rounded-full fill-blue-600 p-2 transition-colors
                hover:bg-gray-100
                active:bg-gray-200
              "
            />
            <Star
              v-else
              @click="toggleStarred(todo)"
              class="
                invisible size-9 cursor-pointer rounded-full p-2 transition-colors
                group-hover:visible
                hover:bg-gray-100
                active:bg-gray-200
              "
            />
          </div>

        </div>
        <p class="text-sm">{{ todo.description }}</p>
      </div>
    </div>
  </div>
</template>
