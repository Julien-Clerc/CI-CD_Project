<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Créer un nouvel utilisateur</h1>
    <form @submit.prevent="createUser">
      <div class="mb-4">
        <label for="username" class="block text-gray-700 font-bold mb-2">Nom d'utilisateur</label>
        <input
          type="text"
          id="username"
          v-model="username"
          placeholder="Entrez votre nom"
          class="border p-2 w-full"
        />
      </div>
      <button type="submit" class="bg-blue-500 text-white p-2">Valider</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'

const username = ref('')

const createUser = async () => {
  try {
    // const response = await axios.post('/api/users', { name: username.value })
    // const userId = response.data.id
    const userId = String(Math.random() * 10);
    const currentUser = {id: userId, name: username.value}

    localStorage.setItem('currentUser', JSON.stringify(currentUser))
    alert('Utilisateur créé avec succès!')
    // Rediriger l'utilisateur vers une autre page si nécessaire
    window.location.reload()
  } catch (error) {
    console.error('Erreur lors de la création de l\'utilisateur:', error)
    alert('Erreur lors de la création de l\'utilisateur. Veuillez réessayer.')
  }
}
</script>
