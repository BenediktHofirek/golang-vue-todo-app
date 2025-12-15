<script setup lang="ts">
import type { Todo } from '@/models'
import RadioButtonUnchecked from '@material-symbols/svg-600/rounded/radio_button_unchecked.svg'
import Delete from '@material-symbols/svg-600/rounded/delete.svg'
import MoreVert from '@material-symbols/svg-700/rounded/more_vert.svg'
import Star from '@material-symbols/svg-600/rounded/star.svg'
import StarFilled from '@material-symbols/svg-600/rounded/star-fill.svg'
import Check from '@material-symbols/svg-600/rounded/check.svg'
import { useMutation } from '@tanstack/vue-query'
import { api } from '@/apiClient'
import { computed, nextTick, ref, useTemplateRef } from 'vue'
import { onClickOutside, watchDebounced } from '@vueuse/core'

const props = defineProps<{
  todo: Todo
  isEdited: boolean
}>()

const isEdited = ref(props.isEdited)
const title = ref(props.todo.title)
const description = ref(props.todo.description)

const descriptionRowCount = computed(() => {
  return (
    (description.value || '').split('').filter((c) => c === '\n').length + 1
  )
})

const { mutate: updateTodoSilently } = useMutation({
  mutationFn: api.updateTodo,
  mutationKey: ['updateTodo'],
})

const { mutate: updateTodo } = useMutation({
  mutationFn: api.updateTodo,
  mutationKey: ['updateTodo'],
  onSettled: async (_data, _error, _variables, _onMutateResult, context) => {
    await context.client.invalidateQueries({ queryKey: ['todos'] })
  },
})

const formRef = useTemplateRef('formRef')
const titleInputRef = useTemplateRef('titleInputRef')
const descriptionInputRef = useTemplateRef('descriptionInputRef')

onClickOutside(formRef, () => {
  if (!isEdited.value) return

  if (
    title.value === props.todo.title &&
    description.value === props.todo.description
  ) {
    isEdited.value = false
    return
  }

  updateTodo({
    id: props.todo.id,
    title: title.value?.trim(),
    description: description.value?.trim(),
  })
  isEdited.value = false
})
watchDebounced(
  [title, description],
  () => {
    if (!isEdited.value) return

    updateTodoSilently({
      id: props.todo.id,
      title: title.value?.trim(),
      description: description.value?.trim(),
    })
  },
  { debounce: 300 },
)

const { mutate: deleteTodo } = useMutation({
  mutationFn: api.deleteTodo,
  mutationKey: ['deleteTodo'],
})

function toggleCompleted(todo: Todo) {
  updateTodo({
    id: todo.id,
    completed: !todo.completed,
  })
}

function toggleStarred(todo: Todo) {
  updateTodo({
    id: todo.id,
    starred: !todo.starred,
  })
}

async function handleClickTitle() {
  if (isEdited.value) return

  isEdited.value = true
  await nextTick()
  titleInputRef.value?.select()
}

async function handleClickDescription() {
  if (isEdited.value) return

  isEdited.value = true
  await nextTick()
  descriptionInputRef.value?.select()
}

function handleClickContainer() {
  handleClickTitle()
}
</script>

<template>
  <div class="
    group w-full py-2
    hover:bg-blue-50
  " @click="handleClickContainer">
    <div class="flex cursor-default justify-start px-4 transition-colors">
      <div
        class="mt-1.5 mr-4 size-6 cursor-pointer overflow-hidden rounded-full"
        @click.once.stop="toggleCompleted(todo)"
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
      <div class="flex w-full flex-col" ref="formRef">
        <div class="flex w-full items-center justify-between">
          <input
            v-if="isEdited"
            v-model.trim="title"
            ref="titleInputRef"
            class="
              w-full text-base placeholder-gray-500
              focus:outline-none
            "
            placeholder="Title"
          />
          <div
            v-else
            class="flex-1 text-base text-ellipsis"
            @click.stop="handleClickTitle"
            :class="todo.completed ? 'line-through' : ''"
          >
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
            <template v-if="!isEdited">
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
            </template>
          </div>
        </div>

        <textarea
          v-if="isEdited"
          :rows="descriptionRowCount"
          v-model="description"
          placeholder="Detail"
          ref="descriptionInputRef"
          class="
            w-full resize-none text-sm placeholder-gray-500
            focus:outline-none
          "
        ></textarea>
        <p
          v-else
          @click.stop="handleClickDescription"
          class="text-sm whitespace-pre-wrap"
        >
          {{ todo.description }}
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
textarea {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

textarea::-webkit-scrollbar {
  display: none;
}
</style>
