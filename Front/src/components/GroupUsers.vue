<template>
  <div class="p-4">
    <ul class="list-disc pl-5">
      <li v-for="user in users" :key="user.id" class="mb-2">{{ user.name }}</li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

defineProps<{
  users: {
    id: string
    name: string
    group_id?: string
  }[]
  group: {
    id: string
    name: string
  }
}>()
</script>
