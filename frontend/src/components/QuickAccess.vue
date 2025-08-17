<template>
  <div v-if="!searchQuery" class="mb-8">
    <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Quick Access</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <!-- Favorites -->
      <div class="app-card">
        <div class="flex items-center space-x-3">
          <StarIcon class="w-6 h-6 text-yellow-500" />
          <div>
            <h3 class="font-medium text-gray-900 dark:text-white">Favorites</h3>
            <p class="text-sm text-gray-500 dark:text-gray-400">{{ favorites.length }} apps</p>
          </div>
        </div>
      </div>
      
      <!-- Recently Used -->
      <div class="app-card">
        <div class="flex items-center space-x-3">
          <ClockIcon class="w-6 h-6 text-blue-500" />
          <div>
            <h3 class="font-medium text-gray-900 dark:text-white">Recently Used</h3>
            <p class="text-sm text-gray-500 dark:text-gray-400">{{ recentlyUsed.length }} apps</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { StarIcon, ClockIcon } from 'lucide-vue-next'

export default {
  name: 'QuickAccess',
  components: {
    StarIcon,
    ClockIcon
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
  }
}
</script>
