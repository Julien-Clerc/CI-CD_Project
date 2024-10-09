<template>
    <div class="p-4">
      <h1 class="text-2xl font-bold mb-4">Se retirer du groupe</h1>
      <button @click="leaveGroup" class="bg-red-500 text-white p-2">Se retirer</button>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  
  export default defineComponent({
    name: 'LeaveGroup',
    methods: {
      leaveGroup() {
        // Logique pour se retirer du groupe
        console.log('Se retirer du groupe')
      },
    },
  })
  </script>
  