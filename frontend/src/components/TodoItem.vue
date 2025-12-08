<script setup lang="ts">
import type { Todo } from '@/models'
import RadioButtonUnchecked from '@material-symbols/svg-600/outlined/radio_button_unchecked.svg'
import Check from '@material-symbols/svg-600/outlined/check.svg'
import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { apiClient } from '@/apiClient'

defineProps<{
  todo: Todo
}>()

// const createdDate = useDateFormat(props.todo.createdAt, 'ddd D MMM')

const queryClient = useQueryClient()
const { mutate: updateTodo } = useMutation({
  mutationFn: (updatedTodo: Partial<Todo>) => {
    return apiClient.patch(`/todos/${updatedTodo.id}`, updatedTodo)
  },
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['todos'] }),
  mutationKey: ['updateTodo'],
})

async function toggleCompleted(todo: Todo) {
  updateTodo({
    id: todo.id,
    completed: !todo.completed,
  });
}
</script>

<template>
  <div class="w-full py-2">
    <div
      class="
        flex cursor-default justify-start px-4 transition-colors
        hover:bg-blue-50
      "
    >
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
        <div v-else class="group flex size-6 items-center justify-center">
          <RadioButtonUnchecked class="
            size-5
            group-hover:hidden
          " />
          <Check
            class="
              hidden size-6 bg-gray-200 fill-blue-600
              group-hover:block
            "
          />
        </div>
      </div>
      <div class="flex flex-col">
        <div class="text-base" :class="todo.completed ? 'line-through' : ''">
          {{ todo.title }}
        </div>
        <p class="text-sm">{{ todo.description }}</p>
      </div>
    </div>
  </div>
</template>
