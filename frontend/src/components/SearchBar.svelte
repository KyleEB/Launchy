<script lang="ts">
  import { onMount } from 'svelte';
  
  const { onSearch, onEnter, onArrow } = $props<{
    onSearch: (query: string) => void;
    onEnter: () => void;
    onArrow: (direction: string) => void;
  }>();
  
  let searchQuery = $state('');
  let searchTimeout = $state<number>(0);
  let searchInput = $state<HTMLInputElement | null>(null);
  
  function handleInput() {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
      onSearch(searchQuery);
    }, 150);
  }
  
  function clearSearch() {
    searchQuery = '';
    onSearch('');
  }

  // Focus the search input
  function focusSearch() {
    if (searchInput) {
      searchInput.focus();
    }
  }

  // Auto-focus on mount
  onMount(() => {
    focusSearch();
  });

  // Method to focus and type a character
  function focusAndType(char: string) {
    focusSearch();
    searchQuery += char;
    onSearch(searchQuery);
  }

  // Expose the method for parent components
  export { focusAndType };
</script>

<div class="relative flex items-center w-full">
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none">
    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="text-gray-400">
      <circle cx="11" cy="11" r="8"/>
      <path d="m21 21-4.35-4.35"/>
    </svg>
  </div>
  
  <input
    type="text"
    placeholder="Search applications, commands, or executables..."
    bind:value={searchQuery}
    on:input={handleInput}
    on:keydown={(e) => {
      if (e.key === 'Enter') {
        e.preventDefault();
        onEnter();
      }
      if (e.key === 'ArrowUp' || e.key === 'ArrowDown') {
        e.preventDefault();
        onArrow(e.key);
      }
    }}
    aria-label="Search applications"
    class="w-full pl-10 pr-8 py-2 text-sm rounded-md border border-gray-600 bg-gray-800 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
    bind:this={searchInput}
  />
  
  {#if searchQuery}
    <button
      type="button"
      on:click={clearSearch}
      aria-label="Clear search"
      class="absolute right-2 top-1/2 -translate-y-1/2 p-1 rounded text-gray-400 hover:bg-gray-700 transition-colors"
    >
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <line x1="18" y1="6" x2="6" y2="18"/>
        <line x1="6" y1="6" x2="18" y2="18"/>
      </svg>
    </button>
  {/if}
</div>


