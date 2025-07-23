import { defineStore } from 'pinia'
import request from '@/utils/request'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    user: JSON.parse(localStorage.getItem('user') || 'null')
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    isAdmin: (state) => state.user?.is_admin || false
  },

  actions: {
    async login(username, password) {
      try {
        const response = await request.post('/auth/login', {
          username,
          password
        })
        
        this.token = response.token
        this.user = response.user
        
        localStorage.setItem('token', response.token)
        localStorage.setItem('user', JSON.stringify(response.user))
        
        return response
      } catch (error) {
        throw error
      }
    },

    async register(username, password, email) {
      try {
        const response = await request.post('/auth/register', {
          username,
          password,
          email
        })
        
        this.token = response.token
        this.user = response.user
        
        localStorage.setItem('token', response.token)
        localStorage.setItem('user', JSON.stringify(response.user))
        
        return response
      } catch (error) {
        throw error
      }
    },

    async getMe() {
      try {
        const user = await request.get('/auth/me')
        this.user = user
        localStorage.setItem('user', JSON.stringify(user))
        return user
      } catch (error) {
        throw error
      }
    },

    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    }
  }
})