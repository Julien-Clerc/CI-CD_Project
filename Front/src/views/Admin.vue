<template>
  <div class="p-4">
    <h1 class="text-3xl font-bold mb-4">Configuration des groupes</h1>
    <form @submit.prevent="saveConfiguration">
      <div class="mb-4">
        <label for="numUsers" class="block text-gray-700 font-bold mb-2">Nombre d'utilisateurs</label>
        <input
          type="number"
          id="numUsers"
          v-model="numUsers"
          placeholder="Entrez le nombre d'utilisateurs"
          class="border p-2 w-full"
        />
      </div>
      <div class="mb-4">
        <label for="numGroups" class="block text-gray-700 font-bold mb-2">Nombre de groupes</label>
        <input
          type="number"
          id="numGroups"
          v-model="numGroups"
          placeholder="Entrez le nombre de groupes"
          class="border p-2 w-full"
        />
      </div>
      <div class="mb-4">
        <label for="lastGroupConfig" class="block text-gray-700 font-bold mb-2">Configuration du dernier groupe</label>
        <select id="lastGroupConfig" v-model="lastGroupConfig" class="border p-2 w-full">
          <option value="LAST_MIN">LAST_MIN</option>
          <option value="LAST_MAX">LAST_MAX</option>
        </select>
      </div>
      <button type="submit" class="bg-blue-500 text-white p-2">Enregistrer</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const numUsers = ref(0)
const numGroups = ref(0)
const lastGroupConfig = ref('LAST_MIN')

const saveConfiguration = () => {
  // Logique pour enregistrer la configuration
  console.log('Configuration enregistrée:', {
    numUsers: numUsers.value,
    numGroups: numGroups.value,
    lastGroupConfig: lastGroupConfig.value,
  })
  // Vous pouvez envoyer ces données à votre backend ici
}
</script>
