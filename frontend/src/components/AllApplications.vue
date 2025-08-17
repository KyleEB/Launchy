<template>
  <div v-if="!searchQuery && !selectedCategory">
    <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">All Applications</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
      <AppCard
        v-for="app in allApps"
        :key="app.id"
        :app="app"
        @launch="handleLaunch"
        @toggle-favorite="handleToggleFavorite"
      />
    </div>
  </div>
</template>

<script>
import AppCard from './AppCard.vue'

export default {
  name: 'AllApplications',
  components: {
    AppCard
  },
  props: {
    searchQuery: {
      type: String,
      default: ''
    },
    selectedCategory: {
      type: String,
      default: ''
    },
    allApps: {
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
