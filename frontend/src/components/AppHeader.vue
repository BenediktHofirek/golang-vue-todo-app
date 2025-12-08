<script setup lang="ts">
import { firebaseAuth } from '@/firebase'
import Menu from '@material-symbols/svg-600/outlined/menu.svg'
import { useAuth } from '@vueuse/firebase'
import { ref } from 'vue'

function toggleMenu() {
  console.log('togglingMenu')
}

function toggleProfileMenu() {
  console.log('togglingProfileMenu')
}

const isImageLoadingError = ref(false)

const { user } = useAuth(firebaseAuth)
</script>

<template>
  <header class="flex items-center justify-between p-2">
    <div class="flex items-center gap-4">
      <button
        class="
          cursor-pointer rounded-full p-3.5
          hover:bg-gray-200
        "
        @click="toggleMenu"
      >
        <Menu class="size-6 fill-gray-700" />
      </button>
      <h1 class="text-2xl font-semibold text-gray-700">Todo app</h1>
    </div>

    <div
      @click="toggleProfileMenu"
      class="size-10 cursor-pointer rounded-full overflow-clip object-cover"
    >
      <img
        v-if="user?.photoURL && !isImageLoadingError"
        :src="user?.photoURL"
        :alt="user?.displayName || ''"
        @error="isImageLoadingError = true"
        class="v-full h-full bg-gray-200 object-cover"
      />
      <div
        v-else
        class="
          v-full flex h-full items-center justify-center bg-blue-100 text-2xl font-bold
          text-gray-800
        "
      >
        {{ user?.displayName?.slice(0, 1).toUpperCase() || '' }}
      </div>
    </div>
  </header>
</template>
