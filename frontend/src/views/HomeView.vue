<script setup lang="ts">
import { ref } from 'vue';

  let text = "Hello world"
  const id = "somethin"

  let count = ref(0)
</script>

<template>
  <main :id>
    The count is {{ count }}
    <button @click.prevent="count++">Add one</button>
  </main>
</template>
