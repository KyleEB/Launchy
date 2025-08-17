<template>
  <div class="mb-8">
    <div class="relative cursor-text" @click="handleSearchBarClick">
      <SearchIcon v-if="!isSearching" class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400 pointer-events-none z-10" />
      <LoaderIcon v-else class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-primary-500 pointer-events-none z-10 animate-spin" />
      <input
        v-model="searchQuery"
        @input="handleSearch"
        @keydown.enter="handleEnter"
        @keydown.escape="clearSearch"
        @keyup="handleSearch"
        @click.stop="handleSearchClick"
        type="text"
        placeholder="Search for applications..."
        class="search-input pl-12 pr-4 relative z-20"
        ref="searchInput"
      />
      <button
        v-if="searchQuery && !isSearching"
        @click.stop="clearSearch"
        class="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 z-30"
      >
        <XIcon class="w-5 h-5" />
      </button>
    </div>
  </div>
</template>

<script>
import { ref, watch, nextTick } from 'vue'
import { SearchIcon, XIcon, LoaderIcon } from 'lucide-vue-next'

export default {
  name: 'SearchBar',
  components: {
    SearchIcon,
    XIcon,
    LoaderIcon
  },
  emits: ['search', 'clear', 'enter'],
  setup(props, { emit }) {
    const searchInput = ref(null)
    const searchQuery = ref('')
    const isSearching = ref(false)
    const searchTimeout = ref(null)

    // Watch for search query changes with debouncing
    watch(searchQuery, async (newQuery) => {
      if (!newQuery.trim()) {
        emit('search', [])
        isSearching.value = false
        // Clear any pending search timeout
        if (searchTimeout.value) {
          clearTimeout(searchTimeout.value)
          searchTimeout.value = null
        }
        return
      }

      // Set loading state immediately
      isSearching.value = true

      // Clear any existing timeout
      if (searchTimeout.value) {
        clearTimeout(searchTimeout.value)
      }

      // Debounce the search
      searchTimeout.value = setTimeout(async () => {
        // Only proceed if the query hasn't changed during the debounce period
        if (searchQuery.value === newQuery) {
          try {
            const results = await window.go.main.App.SearchApps(newQuery)
            // Only update results if the query is still the same
            if (searchQuery.value === newQuery) {
              emit('search', results)
            }
          } catch (error) {
            console.error('Search failed:', error)
            // Only update results if the query is still the same
            if (searchQuery.value === newQuery) {
              emit('search', [])
            }
          } finally {
            // Only clear loading state if this was the most recent search
            if (searchQuery.value === newQuery) {
              isSearching.value = false
            }
          }
        }
      }, 300) // 300ms debounce
    })

    const handleSearch = () => {
      // The search is now handled by the watcher with debouncing
      // This function is kept for backward compatibility
    }

    const handleSearchClick = () => {
      // Ensure the search input is focused when clicked
      if (searchInput.value) {
        searchInput.value.focus()
      }
    }

    const handleSearchBarClick = () => {
      // Focus the search input when clicking anywhere in the search bar area
      if (searchInput.value) {
        searchInput.value.focus()
      }
    }

    const handleEnter = () => {
      emit('enter', searchQuery.value)
    }

    const clearSearch = () => {
      searchQuery.value = ''
      emit('search', [])
      isSearching.value = false
      // Clear any pending search timeout
      if (searchTimeout.value) {
        clearTimeout(searchTimeout.value)
        searchTimeout.value = null
      }
      searchInput.value?.focus()
    }

    const focus = () => {
      searchInput.value?.focus()
    }

    const setQuery = (query) => {
      searchQuery.value = query
    }

    // Expose isSearching for parent component
    const getIsSearching = () => {
      return isSearching.value
    }

    return {
      searchInput,
      searchQuery,
      isSearching,
      searchTimeout,
      handleSearch,
      handleSearchClick,
      handleSearchBarClick,
      handleEnter,
      clearSearch,
      focus,
      setQuery,
      getIsSearching
    }
  }
}
</script>
