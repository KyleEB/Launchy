<template>
  <div v-if="searchQuery && searchResults.length > 0" class="mb-8">
    <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">
      Search Results ({{ searchResults.length }})
    </h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
      <AppCard
        v-for="(app, index) in searchResults"
        :key="app.id"
        :app="app"
        :isSelected="index === 0"
        @launch="handleLaunch"
        @toggle-favorite="handleToggleFavorite"
      />
    </div>
  </div>
</template>

<script>
import AppCard from './AppCard.vue'

export default {
  name: 'SearchResults',
  components: {
    AppCard
  },
  props: {
    searchQuery: {
      type: String,
      default: ''
    },
    searchResults: {
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
