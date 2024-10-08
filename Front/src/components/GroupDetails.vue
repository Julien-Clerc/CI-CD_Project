<template>
    <div class="p-4">
      <h1 class="text-2xl font-bold mb-4">Détails du groupe</h1>
      <h2 class="text-xl font-bold mb-2">{{ group.name }}</h2>
      <ul class="list-disc pl-5">
        <li v-for="member in group.members" :key="member.id" class="mb-2">{{ member.name }}</li>
      </ul>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  
  export default defineComponent({
    name: 'GroupDetails',
    data() {
      return {
        group: {
          name: 'Groupe 1',
          members: [
            { id: 1, name: 'Utilisateur 1' },
            { id: 2, name: 'Utilisateur 2' },
            // Ajoutez d'autres membres ici
          ],
        },
      }
    },
  })
  </script>
  