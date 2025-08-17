<template>
  <div v-if="!searchQuery" class="mb-8">
    <div class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700">
      <!-- Categories Header -->
      <button
        @click="handleToggleCategories"
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
            @click="handleSelectCategory(category)"
            class="p-3 bg-gray-50 dark:bg-gray-700 rounded-lg border border-gray-200 dark:border-gray-600 hover:border-primary-300 dark:hover:border-primary-600 transition-colors duration-200 text-left"
          >
            <h3 class="font-medium text-gray-900 dark:text-white text-sm">{{ category }}</h3>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ChevronDownIcon, ChevronRightIcon } from 'lucide-vue-next'

export default {
  name: 'Categories',
  components: {
    ChevronDownIcon,
    ChevronRightIcon
  },
  props: {
    searchQuery: {
      type: String,
      default: ''
    },
    categories: {
      type: Array,
      default: () => []
    },
    isCategoriesExpanded: {
      type: Boolean,
      default: false
    }
  },
  emits: ['toggle-categories', 'select-category'],
  setup(props, { emit }) {
    const handleToggleCategories = () => {
      emit('toggle-categories')
    }

    const handleSelectCategory = (category) => {
      emit('select-category', category)
    }

    return {
      handleToggleCategories,
      handleSelectCategory
    }
  }
}
</script>
