<template>
  <div class="p-4">
    <h1 class="text-3xl font-bold mb-4">Accueil</h1>

    <!--<UserForm v-if="!userId" />

    <template v-else> -->
      <h2 class="text-2xl font-bold mb-4">Utilisateurs sans groupe</h2>
      <UsersWithoutGroup :users="usersWithoutGroup" />
      <GroupList :groups="groups" />
      <JoinGroup />

      <!--<h2 class="text-2xl font-bold mb-4">Membres du groupe 1</h2>
      <GroupUsers :group="groups[0]" :users="group1Users" />
    </template> -->
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import UsersWithoutGroup from './components/UsersWithoutGroup.vue'
import GroupList from './components/GroupList.vue'
import JoinGroup from './components/JoinGroup.vue'
// import GroupUsers from './components/GroupUsers.vue'
// import UserForm from './components/UserForm.vue'

const userId = ref<string | null>(null)

const users = ref([
  { id: 1, name: 'Utilisateur 1', group_id: null },
  { id: 2, name: 'Utilisateur 2', group_id: null },
  { id: 3, name: 'Utilisateur 3', group_id: 1 },
  { id: 4, name: 'Utilisateur 4', group_id: 1 },
])

const groups = ref([
  { id: 1, name: 'Group 1' },
  { id: 2, name: 'Group 2' },
])

const usersWithoutGroup = computed(() => {
  return users.value.filter(user => user.group_id === null)
})

const group1Users = computed(() => {
  return users.value.filter(user => user.group_id === 1)
})

onMounted(() => {
  userId.value = localStorage.getItem('userId')
})
</script>
