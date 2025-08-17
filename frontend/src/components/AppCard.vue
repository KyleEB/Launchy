<template>
  <div class="app-card group" @click="handleLaunch">
    <div class="flex items-start space-x-3">
      <!-- App Icon -->
      <div class="flex-shrink-0">
        <div class="w-12 h-12 bg-gradient-to-br from-primary-500 to-primary-600 rounded-lg flex items-center justify-center text-white font-semibold text-lg">
          {{ appIcon }}
        </div>
      </div>
      
      <!-- App Info -->
      <div class="flex-1 min-w-0">
        <div class="flex items-start justify-between">
          <div class="flex-1 min-w-0">
            <h3 class="text-sm font-medium text-gray-900 dark:text-white truncate">
              {{ app.displayName || app.name }}
            </h3>
            <p v-if="app.description" class="text-xs text-gray-500 dark:text-gray-400 mt-1 line-clamp-2">
              {{ app.description }}
            </p>
            <div class="flex items-center space-x-2 mt-2">
              <span v-if="app.useCount > 0" class="text-xs text-gray-400 dark:text-gray-500">
                Used {{ app.useCount }} times
              </span>
              <span v-if="app.categories && app.categories.length > 0" class="text-xs text-gray-400 dark:text-gray-500">
                {{ app.categories[0] }}
              </span>
            </div>
          </div>
          
          <!-- Favorite Button -->
          <button
            @click.stop="handleToggleFavorite"
            class="flex-shrink-0 p-1 text-gray-400 hover:text-yellow-500 transition-colors duration-200 opacity-0 group-hover:opacity-100"
            :class="{ 'text-yellow-500 opacity-100': app.isFavorite }"
          >
            <StarIcon class="w-4 h-4" :fill="app.isFavorite ? 'currentColor' : 'none'" />
          </button>
        </div>
      </div>
    </div>
    
    <!-- Hover Actions -->
    <div class="absolute inset-0 bg-primary-600 bg-opacity-0 group-hover:bg-opacity-5 transition-all duration-200 rounded-lg flex items-center justify-center opacity-0 group-hover:opacity-100">
      <div class="flex items-center space-x-2">
        <button
          @click.stop="handleLaunch"
          class="btn-primary text-sm px-3 py-1"
        >
          Launch
        </button>
        <button
          @click.stop="handleToggleFavorite"
          class="btn-secondary text-sm px-3 py-1"
        >
          {{ app.isFavorite ? 'Unfavorite' : 'Favorite' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { StarIcon } from 'lucide-vue-next'

export default {
  name: 'AppCard',
  components: {
    StarIcon
  },
  props: {
    app: {
      type: Object,
      required: true
    }
  },
  emits: ['launch', 'toggle-favorite'],
  setup(props, { emit }) {
    const appIcon = computed(() => {
      const name = props.app.displayName || props.app.name
      if (!name) return '?'
      
      // Get first letter of each word, up to 2 characters
      const words = name.split(' ')
      if (words.length >= 2) {
        return (words[0][0] + words[1][0]).toUpperCase()
      }
      return name[0].toUpperCase()
    })

    const handleLaunch = () => {
      emit('launch', props.app.id)
    }

    const handleToggleFavorite = () => {
      emit('toggle-favorite', props.app.id)
    }

    return {
      appIcon,
      handleLaunch,
      handleToggleFavorite
    }
  }
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
