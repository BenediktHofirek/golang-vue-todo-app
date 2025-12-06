<script setup lang="ts">
import type { Todo } from '@/models'
import RadioButtonUnchecked from '@material-symbols/svg-600/outlined/radio_button_unchecked.svg'
import Check from '@material-symbols/svg-600/outlined/check.svg'
import { useDateFormat } from '@vueuse/core'
import { ref } from 'vue'

const props = defineProps<{
  todo: Todo
}>()

const createdDate = useDateFormat(props.todo.createdAt, 'ddd D MMM')

function toggleChecked() {
  console.log('toggling', props.todo.id)
}
</script>

<template>
  <div class="w-full py-2">
    <div class="ml-4">Today</div>
    <div
      class="
        align-start flex px-4 justify-start transition-colors
        hover:bg-blue-50
        cursor-default
      "
    >
      <div
        @click="toggleChecked()"
        class="group mt-1 mr-4 flex size-6 items-center justify-center rounded-full"
      >
        <RadioButtonUnchecked class="
          size-5
          group-hover:hidden
        " />
        <Check
          class="
            hidden size-6 cursor-pointer rounded-full bg-gray-200 fill-blue-600
            group-hover:block
          "
        />
      </div>
      <div class="flex flex-col">
        <div class="text-base">{{ todo.title }}</div>
        <p class="text-sm">{{ todo.description }}</p>
      </div>
    </div>
  </div>
</template>
