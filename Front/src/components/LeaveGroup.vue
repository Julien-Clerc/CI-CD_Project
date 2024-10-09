<template>
    <div class="p-4">
      <h1 class="text-2xl font-bold mb-4">Se retirer du groupe</h1>
      <button @click="leaveGroup" class="bg-red-500 text-white p-2">Se retirer</button>
    </div>
  </template>

  <script lang="ts">
  import { defineComponent } from 'vue'
  import { storeToRefs } from 'pinia';
  import { useUser } from '../stores/user.store';

  export default defineComponent({
    name: 'LeaveGroup',
    methods: {
      leaveGroup() {
        let { currentUser } = storeToRefs(useUser())

        if (currentUser?.value) {
          currentUser.value.group_id = undefined
          localStorage.setItem('currentUser', JSON.stringify(currentUser.value))
        } else {
          console.log('current user is null')
        }

      },
    },
  })
  </script>
