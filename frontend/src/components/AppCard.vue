<template>
  <div 
    class="app-card group relative" 
    :class="{ 
      'opacity-60': !isLaunchable,
      'border-2 border-yellow-400 dark:border-yellow-500': isSelected
    }"
    @click="handleLaunch"
  >
    <!-- Star Icon (Positioned absolutely) -->
    <button
      @click.stop="handleToggleFavorite"
      class="absolute top-2 right-2 z-50 text-gray-400 hover:text-yellow-500 transition-colors duration-200 p-1 rounded bg-white dark:bg-gray-800 shadow-sm"
      :class="{ 'text-yellow-500': app.isFavorite }"
    >
      <StarIcon class="w-4 h-4" :fill="app.isFavorite ? 'currentColor' : 'none'" />
    </button>

    <!-- Main Card Content -->
    <div class="w-full h-full">
      <div class="flex items-center justify-between">
        <!-- App Info -->
        <div class="flex-1 min-w-0">
          <h3 class="text-sm font-medium text-gray-900 dark:text-white truncate">
            {{ app.displayName || app.name }}
          </h3>
          <p v-if="app.description" class="text-xs text-gray-500 dark:text-gray-400 mt-1 line-clamp-1">
            {{ app.description }}
          </p>
          <div class="flex items-center space-x-2 mt-1">
            <span v-if="app.useCount > 0" class="text-xs text-gray-400 dark:text-gray-500">
              {{ app.useCount }} uses
            </span>
            <span v-if="app.categories && app.categories.length > 0" class="text-xs text-gray-400 dark:text-gray-500">
              {{ app.categories[0] }}
            </span>
          </div>
          <!-- Small warning icon for non-launchable apps -->
          <AlertTriangleIcon v-if="!isLaunchable" class="absolute top-2 right-2 w-3 h-3 text-yellow-500 z-10" />
        </div>
      </div>
    </div>
    
    <!-- Hover Actions -->
    <div class="absolute inset-0 bg-primary-600 bg-opacity-0 group-hover:bg-opacity-5 transition-all duration-200 rounded-lg flex items-center justify-center opacity-0 group-hover:opacity-100">
      <div class="flex items-center space-x-2">
        <div
          v-if="!isLaunchable"
          class="flex items-center space-x-1 text-sm text-yellow-600 dark:text-yellow-400 px-3 py-1"
        >
          <AlertTriangleIcon class="w-4 h-4" />
          <span>No executable</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { StarIcon, AlertTriangleIcon } from 'lucide-vue-next'

export default {
  name: 'AppCard',
  components: {
    StarIcon,
    AlertTriangleIcon
  },
  props: {
    app: {
      type: Object,
      required: true
    },
    isSelected: {
      type: Boolean,
      default: false
    }
  },
  emits: ['launch', 'toggle-favorite'],
  setup(props, { emit }) {
    const isLaunchable = computed(() => {
      return props.app.execPath && props.app.execPath.trim() !== ''
    })

    const handleLaunch = () => {
      if (!isLaunchable.value) {
        return
      }
      emit('launch', props.app.id)
    }

    const handleToggleFavorite = () => {
      emit('toggle-favorite', props.app.id)
    }

    return {
      isLaunchable,
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
