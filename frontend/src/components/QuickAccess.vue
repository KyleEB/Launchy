<template>
  <div v-if="!searchQuery" class="mb-8">
    <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Quick Access</h2>
    
    <!-- Favorites Section -->
    <div v-if="favorites.length > 0" class="mb-6">
      <div class="flex items-center space-x-2 mb-3">
        <StarIcon class="w-5 h-5 text-yellow-500" />
        <h3 class="text-md font-medium text-gray-900 dark:text-white">Favorites</h3>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
        <AppCard
          v-for="app in favorites"
          :key="app.id"
          :app="app"
          @launch="handleLaunch"
          @toggle-favorite="handleToggleFavorite"
        />
      </div>
    </div>
    
    <!-- Recently Used Section -->
    <div v-if="recentlyUsed.length > 0" class="mb-6">
      <div class="flex items-center space-x-2 mb-3">
        <ClockIcon class="w-5 h-5 text-blue-500" />
        <h3 class="text-md font-medium text-gray-900 dark:text-white">Recently Used</h3>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
        <AppCard
          v-for="app in recentlyUsed"
          :key="app.id"
          :app="app"
          @launch="handleLaunch"
          @toggle-favorite="handleToggleFavorite"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { StarIcon, ClockIcon } from 'lucide-vue-next'
import AppCard from './AppCard.vue'

export default {
  name: 'QuickAccess',
  components: {
    StarIcon,
    ClockIcon,
    AppCard
  },
  props: {
    searchQuery: {
      type: String,
      default: ''
    },
    favorites: {
      type: Array,
      default: () => []
    },
    recentlyUsed: {
      type: Array,
      default: () => []
    }
  },
  emits: ['launch', 'toggle-favorite'],
  setup(props, { emit }) {
    const handleLaunch = (appId) => {
      emit('launch', appId)
    }

    const handleToggleFavorite = (appId) => {
      emit('toggle-favorite', appId)
    }

    return {
      handleLaunch,
      handleToggleFavorite
    }
  }
}
</script>
