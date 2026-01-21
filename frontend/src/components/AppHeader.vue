<script setup lang="ts">
import { firebaseAuth, signOutUser } from '@/firebase'
import Menu from '@material-symbols/svg-600/outlined/menu.svg'
import Logout from '@material-symbols/svg-600/outlined/logout.svg'
import { useAuth } from '@vueuse/firebase'
import { computed, ref, useTemplateRef } from 'vue'
import { useRouter } from 'vue-router'
import { onClickOutside } from '@vueuse/core'

function toggleMenu() {
  console.log('togglingMenu')
}

const isProfileMenuDisplayed = ref(false)
const isImageLoadingError = ref(false)

const profileMenuDialogRef = useTemplateRef('profileMenuDialogRef')

onClickOutside(profileMenuDialogRef, () => {
  isProfileMenuDisplayed.value = false
})

const profileMenuClasses = computed(() => {
  if (isProfileMenuDisplayed.value) {
    return ['z-50', 'opacity-100']
  } else {
    return ['z-[-1]', 'opacity-0']
  }
})

const { user } = useAuth(firebaseAuth)
const router = useRouter()

async function handleSignOut() {
  await signOutUser()
  await router.replace({ name: 'home' })
}
</script>

<template>
  <header class="flex items-center justify-between p-2">
    <div class="flex items-center gap-2">
      <button
        class="
          cursor-pointer rounded-full p-3.5
          hover:bg-gray-200
        "
        @click="toggleMenu"
      >
        <Menu class="size-6" />
      </button>
      <h1 class="text-2xl font-semibold">Todo app</h1>
    </div>

    <div class="relative mr-3.5 block h-min w-min" ref="profileMenuDialogRef">
      <div
        @click="isProfileMenuDisplayed = !isProfileMenuDisplayed"
        class="
          size-10 cursor-pointer overflow-clip rounded-full object-cover
          hover:outline-4 hover:outline-gray-200
        "
        :class="isProfileMenuDisplayed ? ['outline-gray-200', 'outline-4'] : []"
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
          class="v-full flex h-full items-center justify-center bg-blue-100 text-2xl font-bold"
        >
          {{ user?.displayName?.slice(0, 1).toUpperCase() || '' }}
        </div>
      </div>

      <div
        class="
          absolute top-[calc(100%+0.5rem)] right-0 overflow-hidden rounded-2xl bg-blue-100 shadow-md
          shadow-gray-300 transition-opacity
        "
        :class="profileMenuClasses"
        @click="isProfileMenuDisplayed = !isProfileMenuDisplayed"
      >
        <div
          class="
            flex h-max w-max cursor-pointer items-center gap-2 bg-blue-100 p-4 transition-colors
            hover:bg-blue-200
          "
          @click="handleSignOut"
        >
          <Logout class="size-6" />
          <span class="text-xl font-medium">Sign out</span>
        </div>
      </div>
    </div>
  </header>
</template>
