<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800">
    <!-- Header -->
    <header class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-4">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-primary-600 rounded-lg flex items-center justify-center">
              <RocketIcon class="w-5 h-5 text-white" />
            </div>
            <h1 class="text-xl font-bold text-gray-900 dark:text-white">Launchy</h1>
            <span class="text-sm text-gray-500 dark:text-gray-400">CachyOS App Launcher</span>
          </div>
          
          <div class="flex items-center space-x-2">
            <button 
              @click="refreshApps"
              :disabled="isRefreshing"
              class="btn-secondary flex items-center space-x-2"
            >
              <RefreshCwIcon class="w-4 h-4" :class="{ 'animate-spin': isRefreshing }" />
              <span>Refresh</span>
            </button>
            
            <button 
              @click="toggleDarkMode"
              class="btn-secondary p-2"
            >
              <SunIcon v-if="isDark" class="w-4 h-4" />
              <MoonIcon v-else class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Search Bar -->
      <div class="mb-8">
        <div class="relative">
          <SearchIcon class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
          <input
            v-model="searchQuery"
            @input="handleSearch"
            @keydown.enter="handleEnter"
            @keydown.escape="clearSearch"
            type="text"
            placeholder="Search for applications..."
            class="search-input pl-12 pr-4"
            ref="searchInput"
          />
          <button
            v-if="searchQuery"
            @click="clearSearch"
            class="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
          >
            <XIcon class="w-5 h-5" />
          </button>
        </div>
      </div>

      <!-- Quick Actions -->
      <div v-if="!searchQuery" class="mb-8">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Quick Access</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
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
          
          <!-- All Apps -->
          <div class="app-card">
            <div class="flex items-center space-x-3">
              <GridIcon class="w-6 h-6 text-green-500" />
              <div>
                <h3 class="font-medium text-gray-900 dark:text-white">All Applications</h3>
                <p class="text-sm text-gray-500 dark:text-gray-400">{{ allApps.length }} apps</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Search Results -->
      <div v-if="searchQuery && searchResults.length > 0" class="mb-8">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">
          Search Results ({{ searchResults.length }})
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          <AppCard
            v-for="app in searchResults"
            :key="app.id"
            :app="app"
            @launch="launchApp"
            @toggle-favorite="toggleFavorite"
          />
        </div>
      </div>

      <!-- No Results -->
      <div v-else-if="searchQuery && searchResults.length === 0" class="text-center py-12">
        <SearchIcon class="w-12 h-12 text-gray-400 mx-auto mb-4" />
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">No applications found</h3>
        <p class="text-gray-500 dark:text-gray-400">Try a different search term or browse all applications.</p>
      </div>

      <!-- Categories -->
      <div v-if="!searchQuery" class="mb-8">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Categories</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-3">
          <button
            v-for="category in categories"
            :key="category"
            @click="selectCategory(category)"
            class="p-3 bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-primary-300 dark:hover:border-primary-600 transition-colors duration-200 text-left"
          >
            <h3 class="font-medium text-gray-900 dark:text-white text-sm">{{ category }}</h3>
          </button>
        </div>
      </div>

      <!-- All Applications -->
      <div v-if="!searchQuery && !selectedCategory">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">All Applications</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          <AppCard
            v-for="app in allApps"
            :key="app.id"
            :app="app"
            @launch="launchApp"
            @toggle-favorite="toggleFavorite"
          />
        </div>
      </div>

      <!-- Category Apps -->
      <div v-if="selectedCategory">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
            {{ selectedCategory }} ({{ categoryApps.length }})
          </h2>
          <button @click="clearCategory" class="btn-secondary">
            <ArrowLeftIcon class="w-4 h-4 mr-2" />
            Back to All
          </button>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          <AppCard
            v-for="app in categoryApps"
            :key="app.id"
            :app="app"
            @launch="launchApp"
            @toggle-favorite="toggleFavorite"
          />
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { ref, onMounted, computed, nextTick } from 'vue'
import { useDark, useToggle } from '@vueuse/core'
import {
  SearchIcon,
  XIcon,
  StarIcon,
  ClockIcon,
  GridIcon,
  RefreshCwIcon,
  SunIcon,
  MoonIcon,
  ArrowLeftIcon,
  RocketIcon
} from 'lucide-vue-next'
import AppCard from './components/AppCard.vue'

export default {
  name: 'App',
  components: {
    SearchIcon,
    XIcon,
    StarIcon,
    ClockIcon,
    GridIcon,
    RefreshCwIcon,
    SunIcon,
    MoonIcon,
    ArrowLeftIcon,
    RocketIcon,
    AppCard
  },
  setup() {
    const isDark = useDark()
    const toggleDarkMode = useToggle(isDark)
    
    const searchInput = ref(null)
    const searchQuery = ref('')
    const searchResults = ref([])
    const allApps = ref([])
    const favorites = ref([])
    const recentlyUsed = ref([])
    const categories = ref([])
    const selectedCategory = ref('')
    const categoryApps = ref([])
    const isRefreshing = ref(false)

    // Focus search input on mount
    onMounted(async () => {
      await nextTick()
      if (searchInput.value) {
        searchInput.value.focus()
      }
      await loadInitialData()
    })

    const loadInitialData = async () => {
      try {
        const [apps, favs, recent, cats] = await Promise.all([
          window.go.main.App.SearchApps(''),
          window.go.main.App.GetFavorites(),
          window.go.main.App.GetRecentlyUsed(10),
          window.go.main.App.GetCategories()
        ])
        
        allApps.value = apps
        favorites.value = favs
        recentlyUsed.value = recent
        categories.value = cats
      } catch (error) {
        console.error('Failed to load initial data:', error)
      }
    }

    const handleSearch = async () => {
      if (!searchQuery.value.trim()) {
        searchResults.value = []
        return
      }

      try {
        const results = await window.go.main.App.SearchApps(searchQuery.value)
        searchResults.value = results
      } catch (error) {
        console.error('Search failed:', error)
        searchResults.value = []
      }
    }

    const handleEnter = () => {
      if (searchResults.value.length > 0) {
        launchApp(searchResults.value[0].id)
      }
    }

    const clearSearch = () => {
      searchQuery.value = ''
      searchResults.value = []
      searchInput.value?.focus()
    }

    const launchApp = async (appId) => {
      try {
        await window.go.main.App.LaunchApp(appId)
        // Update recently used
        await loadInitialData()
      } catch (error) {
        console.error('Failed to launch app:', error)
      }
    }

    const toggleFavorite = async (appId) => {
      try {
        await window.go.main.App.ToggleFavorite(appId)
        // Refresh favorites
        await loadInitialData()
      } catch (error) {
        console.error('Failed to toggle favorite:', error)
      }
    }

    const refreshApps = async () => {
      isRefreshing.value = true
      try {
        await window.go.main.App.RefreshApps()
        await loadInitialData()
      } catch (error) {
        console.error('Failed to refresh apps:', error)
      } finally {
        isRefreshing.value = false
      }
    }

    const selectCategory = async (category) => {
      selectedCategory.value = category
      try {
        const apps = await window.go.main.App.GetAppsByCategory(category)
        categoryApps.value = apps
      } catch (error) {
        console.error('Failed to load category apps:', error)
        categoryApps.value = []
      }
    }

    const clearCategory = () => {
      selectedCategory.value = ''
      categoryApps.value = []
    }

    return {
      searchInput,
      searchQuery,
      searchResults,
      allApps,
      favorites,
      recentlyUsed,
      categories,
      selectedCategory,
      categoryApps,
      isRefreshing,
      isDark,
      toggleDarkMode,
      handleSearch,
      handleEnter,
      clearSearch,
      launchApp,
      toggleFavorite,
      refreshApps,
      selectCategory,
      clearCategory
    }
  }
}
</script>
