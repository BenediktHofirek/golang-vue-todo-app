<script setup lang="ts">
import type { Todo } from '@/models'
import { useDateFormat } from '@vueuse/core'

const props = defineProps<{
  todo: Todo
}>()

const createdDate = useDateFormat(props.todo.createdAt, 'ddd D MMM')
</script>

<template>
  <div>
    <h1>{{ createdDate }}</h1>
    <div v-for="(value, key) in todo" :key>
      {{ value }}
    </div>
  </div>
</template>
