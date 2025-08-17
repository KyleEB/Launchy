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
      <SearchBar
        @search="handleSearchResults"
        @enter="handleSearchEnter"
        ref="searchBarRef"
      />

      <!-- Quick Actions -->
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

      <!-- Search Loading -->
      <div v-if="searchQuery && isSearching" class="text-center py-12">
        <div class="w-12 h-12 text-primary-500 mx-auto mb-4 animate-spin">
          <svg class="w-full h-full" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">Searching...</h3>
        <p class="text-gray-500 dark:text-gray-400">Looking for applications matching "{{ searchQuery }}"</p>
      </div>

      <!-- Search Results -->
      <div v-else-if="searchQuery && searchResults.length > 0" class="mb-8">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">
          Search Results ({{ searchResults.length }})
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
          <AppCard
            v-for="(app, index) in searchResults"
            :key="app.id"
            :app="app"
            :isSelected="index === 0"
            @launch="launchApp"
            @toggle-favorite="toggleFavorite"
          />
        </div>
      </div>

      <!-- No Results -->
      <div v-else-if="searchQuery && searchResults.length === 0 && !isSearching" class="text-center py-12">
        <SearchIcon class="w-12 h-12 text-gray-400 mx-auto mb-4" />
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">No applications found</h3>
        <p class="text-gray-500 dark:text-gray-400">Try a different search term or browse all applications.</p>
      </div>

      <!-- Categories -->
      <div v-if="!searchQuery" class="mb-8">
        <div class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700">
          <!-- Categories Header -->
          <button
            @click="toggleCategories"
            class="w-full p-4 flex items-center justify-between hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-200"
          >
            <h2 class="text-lg font-semibold text-gray-900 dark:text-white">Categories</h2>
            <ChevronDownIcon 
              v-if="isCategoriesExpanded" 
              class="w-5 h-5 text-gray-500 dark:text-gray-400 transition-transform duration-200" 
            />
            <ChevronRightIcon 
              v-else 
              class="w-5 h-5 text-gray-500 dark:text-gray-400 transition-transform duration-200" 
            />
          </button>
          
          <!-- Categories Content (Collapsible) -->
          <div 
            v-if="isCategoriesExpanded"
            class="px-4 pb-4 border-t border-gray-100 dark:border-gray-700"
          >
            <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-3 mt-4">
              <button
                v-for="category in categories"
                :key="category"
                @click="selectCategory(category)"
                class="p-3 bg-gray-50 dark:bg-gray-700 rounded-lg border border-gray-200 dark:border-gray-600 hover:border-primary-300 dark:hover:border-primary-600 transition-colors duration-200 text-left"
              >
                <h3 class="font-medium text-gray-900 dark:text-white text-sm">{{ category }}</h3>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- All Applications -->
      <div v-if="!searchQuery && !selectedCategory">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">All Applications</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
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
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-3">
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
  StarIcon,
  ClockIcon,
  GridIcon,
  RefreshCwIcon,
  SunIcon,
  MoonIcon,
  ArrowLeftIcon,
  RocketIcon,
  ChevronDownIcon,
  ChevronRightIcon
} from 'lucide-vue-next'
import AppCard from './components/AppCard.vue'
import SearchBar from './components/SearchBar.vue'

export default {
  name: 'App',
  components: {
    StarIcon,
    ClockIcon,
    GridIcon,
    RefreshCwIcon,
    SunIcon,
    MoonIcon,
    ArrowLeftIcon,
    RocketIcon,
    ChevronDownIcon,
    ChevronRightIcon,
    AppCard,
    SearchBar
  },
  setup() {
    const isDark = useDark()
    const toggleDarkMode = useToggle(isDark)
    
    const searchBarRef = ref(null)
    const searchResults = ref([])
    const allApps = ref([])
    const favorites = ref([])
    const recentlyUsed = ref([])
    const categories = ref([])
    const selectedCategory = ref('')
    const categoryApps = ref([])
    const isRefreshing = ref(false)
    const isCategoriesExpanded = ref(false)

    // Focus search input on mount
    onMounted(async () => {
      await nextTick()
      if (searchBarRef.value) {
        searchBarRef.value.focus()
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

    const handleSearchResults = (results) => {
      searchResults.value = results
    }

    const handleSearchEnter = () => {
      if (searchResults.value.length > 0) {
        launchApp(searchResults.value[0].id)
      }
    }

    // Computed property to get search query from SearchBar component
    const searchQuery = computed(() => {
      return searchBarRef.value?.searchQuery || ''
    })

    // Computed property to get search state from SearchBar component
    const isSearching = computed(() => {
      return searchBarRef.value?.isSearching || false
    })

    const launchApp = async (appId) => {
      try {
        console.log('Attempting to launch app with ID:', appId)
        await window.go.main.App.LaunchApp(appId)
        console.log('Successfully launched app:', appId)
        // Update recently used
        await loadInitialData()
      } catch (error) {
        console.error('Failed to launch app:', appId, error)
        // Show user-friendly error message
        alert(`Failed to launch application: ${error.message || error}`)
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

    const toggleCategories = () => {
      isCategoriesExpanded.value = !isCategoriesExpanded.value
    }

    return {
      searchBarRef,
      searchQuery,
      searchResults,
      allApps,
      favorites,
      recentlyUsed,
      categories,
      selectedCategory,
      categoryApps,
      isRefreshing,
      isSearching,
      isDark,
      toggleDarkMode,
      handleSearchResults,
      handleSearchEnter,
      launchApp,
      toggleFavorite,
      refreshApps,
      selectCategory,
      clearCategory,
      toggleCategories,
      isCategoriesExpanded
    }
  }
}
</script>
