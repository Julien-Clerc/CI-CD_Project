<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Liste des groupes</h1>
    <ul class="list-disc pl-5">
      <li v-for="group in groups" :key="group.id" class="mb-2">
        {{ group.name }}
        <button @click="joinGroup(group.id)" class="bg-blue-500 text-white p-2 ml-4">Entrer dans le groupe</button>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

const props = defineProps<{
  groups: {
    id: number
    name: string
  }[]
}>()

const emit = defineEmits<{
  (e: 'joinGroup', groupId: number): void
}>()

const joinGroup = (groupId: number) => {
  emit('joinGroup', groupId)
}
</script>
