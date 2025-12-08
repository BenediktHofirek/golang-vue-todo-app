<script setup lang="ts">
import { signInWithGooglePopup } from '@/firebase'
import { useRouter } from 'vue-router'

const router = useRouter()

async function signInWithGoogle() {
  await signInWithGooglePopup()
  await router.replace({ name: 'app' })
}
</script>

<template>
  <main class="flex h-full w-full">
    Welcome home!
    <button @click="signInWithGoogle">Sign In with Google</button>
  </main>
</template>
