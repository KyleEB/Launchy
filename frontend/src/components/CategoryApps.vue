<template>
  <div v-if="selectedCategory">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
        {{ selectedCategory }} ({{ categoryApps.length }})
      </h2>
      <button @click="handleClearCategory" class="btn-secondary">
        <ArrowLeftIcon class="w-4 h-4 mr-2" />
        Back to All
      </button>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
      <AppCard
        v-for="app in categoryApps"
        :key="app.id"
        :app="app"
        @launch="handleLaunch"
        @toggle-favorite="handleToggleFavorite"
      />
    </div>
  </div>
</template>

<script>
import { ArrowLeftIcon } from 'lucide-vue-next'
import AppCard from './AppCard.vue'

export default {
  name: 'CategoryApps',
  components: {
    ArrowLeftIcon,
    AppCard
  },
  props: {
    selectedCategory: {
      type: String,
      default: ''
    },
    categoryApps: {
      type: Array,
      default: () => []
    }
  },
  emits: ['launch', 'toggle-favorite', 'clear-category'],
  setup(props, { emit }) {
    const handleLaunch = (appId) => {
      emit('launch', appId)
    }

    const handleToggleFavorite = (appId) => {
      emit('toggle-favorite', appId)
    }

    const handleClearCategory = () => {
      emit('clear-category')
    }

    return {
      handleLaunch,
      handleToggleFavorite,
      handleClearCategory
    }
  }
}
</script>
