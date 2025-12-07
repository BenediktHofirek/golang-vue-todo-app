<script setup lang="ts">
import { firebaseAuth, signInWithGooglePopup } from '@/firebase'
import { useAuth } from '@vueuse/firebase/useAuth'

const { isAuthenticated, user } = useAuth(firebaseAuth)
</script>

<template>
  <main class="flex h-full w-full">
    Welcome home!
    <pre v-if="isAuthenticated">{{ user }}</pre>
    <div v-else>
      <button @click="signInWithGooglePopup">Sign In with Google</button>
    </div>
  </main>
</template>
