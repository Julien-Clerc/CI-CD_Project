<template>
  <div class="p-4 flex flex-row gap-6">
    <div class="flex flex-col">
      <h1 class="text-2xl font-bold mb-4">Rejoindre un groupe aléatoire</h1>
      <button @click="autoJoinGroup" class="bg-blue-500 text-white p-2">Rejoindre</button>
    </div>
    <div class="flex flex-col">
      <h1 class="text-2xl font-bold mb-4">Créer un groupe</h1>
      <button @click="openCreateGroupModal" class="bg-red-500 text-black p-2">Créer</button>
    </div>
  </div>

  <GroupForm v-if="showCreateGroupModal" @createGroup="handleCreateGroup" @closeModal="closeCreateGroupModal" />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import GroupForm from './GroupForm.vue'

const showCreateGroupModal = ref(false)

const autoJoinGroup = () => {
  // Logique pour rejoindre un groupe aléatoire
  console.log('Rejoindre un groupe aléatoire')
}

const openCreateGroupModal = () => {
  showCreateGroupModal.value = true
}

const closeCreateGroupModal = () => {
  showCreateGroupModal.value = false
}

const handleCreateGroup = (groupName: string) => {
  // Logique pour créer un groupe
  console.log('Groupe créé:', groupName)
  closeCreateGroupModal()
}
</script>
