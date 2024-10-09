import { defineStore } from 'pinia';

interface UserState {
    currentUser?: {
        id: string;
        name: string;
        group_id?: string;
    }
}

export const useUser = defineStore('user', {
    state: (): UserState => ({
        currentUser: JSON.parse(localStorage.getItem('currentUser')!)
    }),
})