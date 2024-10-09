<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Liste des groupes</h1>
    <ul class="list-disc pl-5">
      <li v-for="group in groups" :key="group.id" class="mb-2">
        {{ group.name }}
        <button @click="joinGroup(String(group.id))" class="bg-blue-500 text-white p-2 ml-4">Entrer dans le groupe</button>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
  import { defineProps } from 'vue'
  import { storeToRefs } from 'pinia';
  import { useUser } from '../stores/user.store';

  defineProps<{
    groups: {
      id: string
      name: string
    }[]
  }>()

  let { currentUser } = storeToRefs(useUser())

  const joinGroup = (groupId: string) => {
    if (currentUser?.value) {
      currentUser.value.group_id = groupId
      localStorage.setItem('currentUser', JSON.stringify(currentUser.value))
    } else {
      console.log('current user is null')
    }
  }
</script>
