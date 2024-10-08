<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4 text-orange-500">Liste des utilisateurs</h1>
    <ul class="list-disc pl-5">
      <li v-for="user in users" :key="user.id" class="mb-2">{{ user.name }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'UserList',
  data() {
    return {
      users: [
        { id: 1, name: 'Utilisateur 1' },
        { id: 2, name: 'Utilisateur 2' },
        // Ajoutez d'autres utilisateurs ici
      ],
    }
  },
})
</script>