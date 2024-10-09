<template>
  <div class="fixed inset-0 flex items-center justify-center bg-gray-500 bg-opacity-50">
    <div class="bg-white p-6 rounded-lg shadow-lg">
      <h1 class="text-2xl font-bold mb-4">Créer un nouveau groupe</h1>
      <form @submit.prevent="createGroup">
        <div class="mb-4">
          <label for="groupName" class="block text-gray-700 font-bold mb-2">Nom du groupe</label>
          <input
            type="text"
            id="groupName"
            v-model="groupName"
            placeholder="Entrez le nom du groupe"
            class="border p-2 w-full"
          />
        </div>
        <button type="submit" class="bg-blue-500 text-white p-2">Créer</button>
        <button @click="closeModal" class="bg-red-500 text-white p-2 ml-2">Annuler</button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineEmits } from 'vue'

const groupName = ref('')
const emit = defineEmits<{
  (e: 'createGroup', groupName: string): void
  (e: 'closeModal'): void
}>()

const createGroup = () => {
  emit('createGroup', groupName.value)
}

const closeModal = () => {
  emit('closeModal')
}
</script>
