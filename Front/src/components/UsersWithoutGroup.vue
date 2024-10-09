<template>
    <div class="p-4">
      <ul class="list-disc pl-5">
        <li v-for="user in users" :key="user.id" class="mb-2">{{ user.name }}</li>
      </ul>
    </div>
  </template>
  
  <script setup lang="ts">
  defineProps <{
    users: {
      id: number
      name: string
    }[]
  }>();

  </script>
  